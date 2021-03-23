<?php

namespace App\Services;

use App\Exceptions\ErrorException;
use App\Repositories\MakeupInfoRepository;
use App\Repositories\MakeupCostRepository;
use App\Repositories\MakeupSaleRepository;
use Carbon\Carbon;

class MakeupService
{
    protected $makeupInfoRepository;
    protected $makeupCostRepository;
    protected $makeupSaleRepository;

    public function __construct(
        MakeupInfoRepository $makeupInfoRepository,
        MakeupCostRepository $makeupCostRepository,
        MakeupSaleRepository $makeupSaleRepository
    )
    {
        $this->makeupInfoRepository = $makeupInfoRepository;
        $this->makeupCostRepository = $makeupCostRepository;
        $this->makeupSaleRepository = $makeupSaleRepository;
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
            $data = $this->makeupCostRepository->delete($id);
        } catch (\Throwable $th) {
            throw new ErrorException(500, 'error');
        }
        return $data;
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
        try {
            $data = $this->makeupSaleRepository->create($sale_info);
        } catch (\Throwable $th) {
            throw new ErrorException(500, $th->getMessage());
        }
        return $data;
    }

    public function deleteMakeupSale(int $id)
    {
        try {
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
}
