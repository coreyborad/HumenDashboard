<?php

namespace App\Repositories;

use App\Models\MakeupSale;
use Prettus\Repository\Eloquent\BaseRepository;

class MakeupSaleRepository extends BaseRepository
{
    public function model(): string
    {
        return MakeupSale::class;
    }

}
