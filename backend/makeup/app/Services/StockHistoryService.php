<?php

namespace App\Services;

use App\Repositories\StockHistoryRepository;

class StockHistoryService
{
    protected $stockHistorykRepository;

    public function __construct(StockHistoryRepository $stockHistoryRepository)
    {
        $this->stockHistoryRepository = $stockHistoryRepository;
    }
}
