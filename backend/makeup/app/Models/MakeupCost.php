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
        'price',
        'count',
        'order_date',
    ];

    /**
     * The attributes that should be cast to native types.
     *
     * @var array
     */
    protected $casts = [
    ];

}
