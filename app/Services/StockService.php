<?php

namespace App\Services;

use App\Repositories\StockHistoryRepository;

class StockService
{
    protected $stockHistoryRepository;

    public function __construct(StockHistoryRepository $stockHistoryRepository)
    {
        $this->stockHistoryRepository = $stockHistoryRepository;
    }
}
