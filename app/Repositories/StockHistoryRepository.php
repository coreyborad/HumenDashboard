<?php

namespace App\Repositories;

use App\Models\StockHistory;
use Prettus\Repository\Eloquent\BaseRepository;

class StockHistoryRepository extends BaseRepository
{
    public function model(): string
    {
        return StockHistory::class;
    }

}
