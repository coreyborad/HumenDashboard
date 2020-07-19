<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class StockHistory extends Model
{
    protected $connection = 'mysql';
    protected $table = 'stock_history';
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'stock_number',
        'deal_date',
        'deal_count',
        'deal_price',
        'deal_record',
        'price_on_open',
        'price_on_highest',
        'price_on_lowest',
        'price_on_close',
        'price_diff',
    ];

    /**
     * The attributes that should be cast to native types.
     *
     * @var array
     */
    protected $casts = [];

    public function user()
    {
        return $this->hasOne(User::class, 'id', 'user_id');
    }
}
