<?php

namespace App\Services;

use App\Exceptions\ErrorException;
use App\Repositories\MakeupInfoRepository;
use App\Repositories\MakeupCostRepository;
use App\Repositories\MakeupSaleRepository;
use App\Repositories\MakeupSaleCostRelateRepository;
use Carbon\Carbon;

class MakeupService
{
    protected $makeupInfoRepository;
    protected $makeupCostRepository;
    protected $makeupSaleRepository;
    protected $makeupSaleCostRelateRepository;

    public function __construct(
        MakeupInfoRepository $makeupInfoRepository,
        MakeupCostRepository $makeupCostRepository,
        MakeupSaleRepository $makeupSaleRepository,
        MakeupSaleCostRelateRepository $makeupSaleCostRelateRepository
    )
    {
        $this->makeupInfoRepository = $makeupInfoRepository;
        $this->makeupCostRepository = $makeupCostRepository;
        $this->makeupSaleRepository = $makeupSaleRepository;
        $this->makeupSaleCostRelateRepository = $makeupSaleCostRelateRepository;
    }

    public function getUserMakeupList(int $user_id)
    {
        $makeup_list = $this->makeupInfoRepository->select(['brand', 'name'])->groupBy(['name', 'brand'])->get();
        $makeup_list = $makeup_list->map(function($makeup){
            $makeup->cost_total = 0;
            $makeup->cost_count = 0;
            $makeup->sale_total = 0;
            $makeup->sale_count = 0;
            $makeup->income = 0;

            $color_list = $this->makeupInfoRepository->with(['costs', 'sales'])->findWhere(
                [
                    'brand' => $makeup->brand,
                    'name'  => $makeup->name
                ]
            );
            $color_list->each(function($color) use($makeup){
                $color->costs->each(function($cost) use($makeup){
                    $makeup->cost_total += $cost->price * $cost->count;
                    $makeup->cost_count += $cost->count;
                });

                $color->sales->each(function($sale) use($makeup){
                    $makeup->sale_total += $sale->price * $sale->count;
                    $makeup->sale_count += $sale->count;
                });
            });
            $makeup->color_list = $color_list;
            return $makeup;
        });
        return $makeup_list;
    }

    public function getMakeupByQuery(array $query = [])
    {
        $findList = [];
        foreach ($query as $key => $value) {
            array_push($findList, [$key, 'like', '%'. $value .'%']);
        }
        return $this->makeupInfoRepository->findWhere($findList);
    }

    public function createMakeupInfo(array $makeup_info)
    {
        try {
            $data = $this->makeupInfoRepository->create($makeup_info);
        } catch (\Throwable $th) {
            throw new ErrorException(500, $th->getMessage());
        }
        return $data;
    }

    public function deleteMakeupInfo(int $id)
    {
        try {
            $data = $this->makeupInfoRepository->delete($id);
            // 尋找cost跟sale
            $costs = $this->makeupCostRepository->deleteWhere(
                [
                    'makeup_id' => $id
                ]
            );
            $sales = $this->makeupSaleRepository->deleteWhere(
                [
                    'makeup_id' => $id
                ]
            );

        } catch (\Throwable $th) {
            throw new ErrorException(500, 'error');
        }
        return $data;
    }

    public function updateMakeupInfo(int $id, array $makeup_info)
    {
        try {
            $data = $this->makeupInfoRepository->update($makeup_info, $id);
        } catch (\Throwable $th) {
            throw new ErrorException(500, 'error');
        }
        return $data;
    }

    public function createMakeupCost(array $cost_info)
    {
        try {
            $data = $this->makeupCostRepository->create($cost_info);
        } catch (\Throwable $th) {
            throw new ErrorException(500, $th->getMessage());
        }
        return $data;
    }

    public function deleteMakeupCost(int $id)
    {
        try {
            $relates = $this->makeupSaleCostRelateRepository->findWhere([
                'cost_id' => $id,
            ]);
            // 刪除有建立關係的銷售表
            foreach ($relates as $relate) {
                $this->makeupSaleRepository->delete($relate->sale_id);
            }
            // 刪除關係表
            $this->makeupSaleCostRelateRepository->deleteWhere([
                'cost_id' => $id,
            ]);
            $this->makeupCostRepository->delete($id);
        } catch (\Throwable $th) {
            throw new ErrorException(500, 'error');
        }
        return $relates;
    }

    public function updateMakeupCost(int $id, array $cost_info)
    {
        try {
            $data = $this->makeupCostRepository->update($cost_info, $id);
        } catch (\Throwable $th) {
            throw new ErrorException(500, 'error');
        }
        return $data;
    }

    public function createMakeupSale(array $sale_info)
    {
        //  ['makeup_id', 'price', 'count', 'sold_date']
        try {
            // 先取得該產品庫存數量避免有庫存不夠，卻還賣東西的狀況
            $inventory_status = $this->getMakeupInventory($sale_info['makeup_id']);
            if($inventory_status['inventory'] < $sale_info['count']){
                throw new ErrorException(500, 'error');
            }
            // 尋找要對應到的成本id
            $remain = $sale_info['count'];
            $cost_list = $this->makeupCostRepository
                ->with(['hadSold'])
                ->orderBy('order_date', 'asc')
                ->findWhere(
                    [
                        'makeup_id' => $sale_info['makeup_id'],
                    ]
                );
            $cost_list = $cost_list->filter(function($cost){
                $rem = ($cost->count - $cost->hadSold->sum('relate_count'));
                if($rem <= 0){
                    return false;
                }else{
                    return true;
                }
            });
            // 先新增銷售單
            $sale = $this->makeupSaleRepository->create($sale_info);
            // 新增至銷售成本關係表
            foreach ($cost_list as $cost) {
                if($remain > $cost->count){
                    $remain = $remain - $cost->count;
                    $this->makeupSaleCostRelateRepository->create([
                        'cost_id' => $cost->id,
                        'sale_id' => $sale->id,
                        'relate_count' => $cost->count
                    ]);
                // 小於或等於時代表銷售數量已攤平完畢
                } else {
                    $this->makeupSaleCostRelateRepository->create([
                        'cost_id' => $cost->id,
                        'sale_id' => $sale->id,
                        'relate_count' => $remain
                    ]);
                    break;
                }
            }
        } catch (\Throwable $th) {
            throw new ErrorException(500, $th->getMessage());
        }
        return $sale;
    }

    public function deleteMakeupSale(int $id)
    {
        try {
            // 要先刪除成本銷售關聯表
            $this->makeupSaleCostRelateRepository->deleteWhere([
                'sale_id' => $id,
            ]);
            $data = $this->makeupSaleRepository->delete($id);
        } catch (\Throwable $th) {
            throw new ErrorException(500, 'error');
        }
        return $data;
    }

    public function updateMakeupSale(int $id, array $sale_info)
    {
        try {
            $data = $this->makeupSaleRepository->update($sale_info, $id);
        } catch (\Throwable $th) {
            throw new ErrorException(500, 'error');
        }
        return $data;
    }

    public function getMakeupCostByDate(Carbon $start, Carbon $end)
    {
        $data = $this->makeupCostRepository
            ->findWhereBetween('order_date', [$start->toDateString(), $end->endOfMonth()->toDateString()])
            ->groupBy(function ($item) {
                return substr($item['order_date'], 0, 7);
            });
        $diff_months = $start->diffInMonths($end);
        $result = [];
        for ($i=0; $i <= $diff_months; $i++) {
            $key = '';
            // start
            if($i === 0){
                $key = substr($start->toDateString(), 0, 7);
            }
            // end
            else if($i === $diff_months){
                $key = substr($end->toDateString(), 0, 7);
            }else{
                $key = substr($start->copy()->addMonthNoOverflow($i)->toDateString(), 0, 7);
            }
            // 如果沒有資料，設定為0
            if (isset($data[$key])) {
                $result[$i] = [
                    'month' => $key,
                    'price' => 0
                ];
                foreach ($data[$key] as $record) {
                    $result[$i]['price'] += $record->price;
                }
            }else{
                $result[$i] = [
                    'month' => $key,
                    'price' => 0
                ];
            }
        }
        return $result;
    }
    public function getMakeupSaleByDate(Carbon $start, Carbon $end)
    {
        $data = $this->makeupSaleRepository
            ->findWhereBetween('sold_date', [$start->toDateString(), $end->endOfMonth()->toDateString()])
            ->groupBy(function ($item) {
                return substr($item['sold_date'], 0, 7);
            });
        $diff_months = $start->diffInMonths($end);
        $result = [];
        for ($i=0; $i <= $diff_months; $i++) {
            $key = '';
            // start
            if($i === 0){
                $key = substr($start->toDateString(), 0, 7);
            }
            // end
            else if($i === $diff_months){
                $key = substr($end->toDateString(), 0, 7);
            }else{
                $key = substr($start->copy()->addMonthNoOverflow($i)->toDateString(), 0, 7);
            }
            // 如果沒有資料，設定為0
            if (isset($data[$key])) {
                $result[$i] = [
                    'month' => $key,
                    'price' => 0
                ];
                foreach ($data[$key] as $record) {
                    $result[$i]['price'] += $record->price;
                }
            }else{
                $result[$i] = [
                    'month' => $key,
                    'price' => 0
                ];
            }
        }
        return $result;
    }

    public function getMakeupCostGroupByItemOnDate(Carbon $date)
    {
        $data = $this->makeupCostRepository
            ->with('makeup')
            ->findWhereBetween('order_date', [$date->copy()->startOfMonth()->toDateString(), $date->copy()->endOfMonth()->toDateString()])
            ->groupBy(function ($item) {
                return $item['makeup']['id'];
            });
        return $data;
    }

    public function getMakeupInventory(Int $makeup_id)
    {
        $cost_list = $this->makeupCostRepository
            ->with(['hadSold'])
            ->findWhere(
                [
                    'makeup_id' => $makeup_id,
                ]
            );
        $total = $cost_list->sum(function($cost){ return $cost->count; });
        $inventory = $cost_list->sum(function($cost){
            return ($cost->count - $cost->hadSold->sum('relate_count'));
        });
        return [
            'inventory' => $inventory,
            'total' => $total
        ];
    }

    public function getMakeupRealSaleReportByDate(Carbon $start, Carbon $end)
    {
        $data = $this->makeupSaleRepository
            ->with('hadSold', 'hadSold.cost')
            ->findWhereBetween('sold_date', [$start->startOfMonth()->toDateString(), $end->endOfMonth()->toDateString()])
            ->map(function($item){
                $item->cost = $item->hadSold->sum(function($h){
                    return $h->cost->price;
                });
                return $item;
            })
            ->groupBy(function ($item) {
                return substr($item['sold_date'], 0, 7);
            });
        $diff_months = $start->diffInMonths($end);
        $result = [];
        for ($i=0; $i <= $diff_months; $i++) {
            $key = '';
            // start
            if($i === 0){
                $key = substr($start->toDateString(), 0, 7);
            }
            // end
            else if($i === $diff_months){
                $key = substr($end->toDateString(), 0, 7);
            }else{
                $key = substr($start->copy()->addMonthNoOverflow($i)->toDateString(), 0, 7);
            }
            // 如果沒有資料，設定為0
            if (isset($data[$key])) {
                $result[$i] = [
                    'month' => $key,
                    'cost' => 0,
                    'sale' => 0
                ];
                foreach ($data[$key] as $record) {
                    $result[$i]['sale'] += $record->price;
                    $result[$i]['cost'] += $record->cost;
                }
            }else{
                $result[$i] = [
                    'month' => $key,
                    'cost' => 0,
                    'sale' => 0
                ];
            }
        }
        return $result;
    }

    public function getMakeupSaleCountReportByMonth(Carbon $date)
    {
        $data = $this->makeupSaleRepository
            ->with('makeup')
            ->findWhereBetween('sold_date', [$date->copy()->startOfMonth()->toDateString(), $date->copy()->endOfMonth()->toDateString()])
            ->groupBy(function ($item) {
                return $item['makeup']['id'];
            });
        $result = [];
        foreach ($data as $sale) {
            array_push($result, [
                'info' => $sale[0]['makeup'],
                'count' => $sale->sum(function($s){
                    return $s->count;
                })
            ]);
        }
        return $result;
    }
}
