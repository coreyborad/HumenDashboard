<?php

namespace App\Repositories;

use App\Models\MakeupPrice;
use Prettus\Repository\Eloquent\BaseRepository;

class MakeupPriceRepository extends BaseRepository
{
    public function model(): string
    {
        return MakeupPrice::class;
    }

}
