<?php

namespace App\Services;

use App\Exceptions\ErrorException;
use App\Repositories\UserHasStockRepository;
use App\Repositories\StockInfoRepository;
use App\Repositories\Mongo\StockHistoryRepository;


class StockService
{
    protected $stockHistoryRepository;
    protected $userHasStockRepository;
    protected $stockInfoRepository;

    public function __construct(
        StockHistoryRepository $stockHistoryRepository,
        UserHasStockRepository $userHasStockRepository,
        StockInfoRepository $stockInfoRepository
    )
    {
        $this->stockHistoryRepository = $stockHistoryRepository;
        $this->userHasStockRepository = $userHasStockRepository;
        $this->stockInfoRepository = $stockInfoRepository;
    }

    public function getUserStockList(int $user_id)
    {
        $user_stock_list = $this->userHasStockRepository->with('stock_info')->findByField('user_id', $user_id);
        $user_stock_list = $user_stock_list->map(function($stock){
            $last_stock = $this->stockHistoryRepository
                ->where('stock_number', $stock->stock_number)
                ->orderBy('deal_date', 'desc')
                ->first();
            $stock->last_stock = $last_stock;
            return $stock;
        });
        return $user_stock_list;
    }

    public function createUserStock(int $user_id, array $data)
    {
        try {
            $data = $this->userHasStockRepository->create([
                'user_id' => $user_id,
                'stock_number' => $data['stock_number'],
                'shares' => $data['shares'],
                'cost' => $data['cost'],
            ]);
        } catch (\Throwable $th) {
            throw new ErrorException(500, 'error');
        }
        return $data;
    }

    public function deleteUserStock(int $id)
    {
        try {
            $data = $this->userHasStockRepository->delete($id);
        } catch (\Throwable $th) {
            throw new ErrorException(500, 'error');
        }
        return $data;
    }

    public function getStockList()
    {
        try {
            $data = $this->stockInfoRepository->get();
        } catch (\Throwable $th) {
            throw new ErrorException(500, 'error');
        }
        return $data;
    }

    public function updateUserStock(int $id, array $info)
    {
        try {
            $data = $this->userHasStockRepository->update($info, $id);
        } catch (\Throwable $th) {
            throw new ErrorException(500, 'error');
        }
        return $data;
    }
}
