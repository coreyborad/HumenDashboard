<?php

namespace App\Repositories;

use App\Models\Stock;
use Prettus\Repository\Eloquent\BaseRepository;

class StockRepository extends BaseRepository
{
    public function model(): string
    {
        return Stock::class;
    }

}
