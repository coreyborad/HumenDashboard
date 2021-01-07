<?php

namespace App\Repositories;

use App\Models\MakeupCost;
use Prettus\Repository\Eloquent\BaseRepository;

class MakeupCostRepository extends BaseRepository
{
    public function model(): string
    {
        return MakeupCost::class;
    }

}
