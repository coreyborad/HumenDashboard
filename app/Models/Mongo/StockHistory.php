<?php

namespace App\Models\Mongo;

use Jenssegers\Mongodb\Eloquent\Model;
use App\Models\User;
use Carbon\Carbon;

class StockHistory extends Model
{
    protected $connection = 'mongodb';
    protected $collection = 'stock_history';

    protected $dates = ['deal_date'];

    protected $fillable = [
        'stock_number',
        'deal_date',
        'deal_count',
        'price_on_open',
        'price_on_highest',
        'price_on_lowest',
        'price_on_close',
    ];

    public function user()
    {
        return $this->hasOne(User::class, 'id', 'user_id');
    }

    public function getCreatedAtAttribute($value)
    {
        return Carbon::createFromTimestamp(strtotime($value))
            ->timezone('Asia/Taipei')
            ->toDateTimeString()
        ;
    }
    public function getUpdatedAtAttribute($value)
    {
        return Carbon::createFromTimestamp(strtotime($value))
            ->timezone('Asia/Taipei')
            ->toDateTimeString()
        ;
    }
    public function getDealDateAttribute($value)
    {
        return Carbon::createFromTimestamp(strtotime($value))
            ->timezone('Asia/Taipei')
            ->toDateTimeString()
        ;
    }
}
