<?php

namespace App\Repositories;

use App\Models\UserHasStock;
use Prettus\Repository\Eloquent\BaseRepository;

class UserHasStockRepository extends BaseRepository
{
    public function model(): string
    {
        return UserHasStock::class;
    }

}
