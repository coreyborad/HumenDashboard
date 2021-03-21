<?php

namespace App\Repositories\Mongo;

use App\Models\Mongo\StockHistory;
use Prettus\Repository\Eloquent\BaseRepository;

class StockHistoryRepository extends BaseRepository
{
    public function model(): string
    {
        return StockHistory::class;
    }

}
