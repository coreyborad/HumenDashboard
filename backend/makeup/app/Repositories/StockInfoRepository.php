<?php

namespace App\Repositories;

use App\Models\StockInfo;
use Prettus\Repository\Eloquent\BaseRepository;

class StockInfoRepository extends BaseRepository
{
    public function model(): string
    {
        return StockInfo::class;
    }

}
