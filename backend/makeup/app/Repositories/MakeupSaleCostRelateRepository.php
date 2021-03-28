<?php

namespace App\Repositories;

use App\Models\MakeupSaleCostRelate;
use Prettus\Repository\Eloquent\BaseRepository;

class MakeupSaleCostRelateRepository extends BaseRepository
{
    public function model(): string
    {
        return MakeupSaleCostRelate::class;
    }

}
