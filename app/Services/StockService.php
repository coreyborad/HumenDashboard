<?php

namespace App\Services;

use App\Repositories\StockRepository;

class StockService
{
    protected $stockRepository;

    public function __construct(StockRepository $stockRepository)
    {
        $this->stockRepository = $stockRepository;
    }
}
