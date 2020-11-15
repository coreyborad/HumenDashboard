<?php

namespace App\Services;

use App\Repositories\UserHasStockRepository;
use App\Repositories\Mongo\StockHistoryRepository;

class StockService
{
    protected $stockHistoryRepository;
    protected $userHasStockRepository;

    public function __construct(
        StockHistoryRepository $stockHistoryRepository,
        UserHasStockRepository $userHasStockRepository
    )
    {
        $this->stockHistoryRepository = $stockHistoryRepository;
        $this->userHasStockRepository = $userHasStockRepository;
    }

    public function getUserStockList(int $user_id)
    {
        $user_stock_list = $this->userHasStockRepository->with('stock_info')->findByField('user_id', $user_id);
        // $lasttwodays = $this->stockHistoryRepository
        //     ->select('deal_date')
        //     ->groupBy('deal_date')
        //     ->orderBy('deal_date', 'desc')
        //     ->limit(2)
        //     ->get();
        $user_stock_list = $user_stock_list->map(function($stock){
            $last_stock = $this->stockHistoryRepository
                ->orderBy('deal_date', 'desc')
                ->where('stock_number', $stock->stock_number)
                ->first();
            $stock->last_stock = $last_stock;
            // $stock->asd = $lasttwodays;
            return $stock;
        });
        return $user_stock_list;
    }
}
