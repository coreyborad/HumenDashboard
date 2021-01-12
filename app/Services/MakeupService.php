<?php

namespace App\Services;

use App\Exceptions\ErrorException;
use App\Repositories\MakeupInfoRepository;
use App\Repositories\MakeupCostRepository;
use App\Repositories\MakeupSaleRepository;


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

}
