<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class MakeupPrice extends Model
{
    protected $connection = 'mysql';
    protected $table = 'makeup_price';
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'makeup_id',
        'cost_price',
        'sale_price',
        'inventory_count',
        'sold_count',
    ];

    /**
     * The attributes that should be cast to native types.
     *
     * @var array
     */
    protected $casts = [];

}
