<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class MakeupSale extends Model
{
    protected $connection = 'mysql';
    protected $table = 'makeup_sale';
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'makeup_id',
        'price',
        'count',
        'sold_date',
    ];

    /**
     * The attributes that should be cast to native types.
     *
     * @var array
     */


    // public function cost()
    // {
    //     return $this->hasOne(MakeupCost::class, 'id', 'makeup_cost_id');
    // }
}
