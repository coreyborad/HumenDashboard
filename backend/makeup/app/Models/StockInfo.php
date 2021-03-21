<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class StockInfo extends Model
{
    protected $connection = 'mysql';
    protected $table = 'stock_info';
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'stock_number',
        'stock_name',
    ];

    /**
     * The attributes that should be cast to native types.
     *
     * @var array
     */
    protected $casts = [];
}
