<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class MakeupCost extends Model
{
    protected $connection = 'mysql';
    protected $table = 'makeup_cost';
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'makeup_id',
        'makeup_cost_id',
        'price',
        'count',
        'sold_date',
    ];

    /**
     * The attributes that should be cast to native types.
     *
     * @var array
     */
    protected $casts = [];

}
