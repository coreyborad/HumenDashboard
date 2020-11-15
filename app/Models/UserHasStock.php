<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class UserHasStock extends Model
{
    protected $connection = 'mysql';
    protected $table = 'user_has_stock';
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'stock_number',
        'user_id',
        'shares',
        'cost'
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
    public function stock_info()
    {
        return $this->hasOne(StockInfo::class, 'stock_number', 'stock_number');
    }
}
