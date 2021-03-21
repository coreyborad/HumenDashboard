<?php

namespace App\Repositories;

use App\Models\MakeupInfo;
use Prettus\Repository\Eloquent\BaseRepository;

class MakeupInfoRepository extends BaseRepository
{
    public function model(): string
    {
        return MakeupInfo::class;
    }

    public function groupBy(array $column_name){
        return $this->model->groupBy($column_name);
    }
}
