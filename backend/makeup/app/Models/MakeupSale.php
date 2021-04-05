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

    public function makeup()
    {
        return $this->hasOne(MakeupInfo::class, 'id', 'makeup_id');
    }

    public function hadSold()
    {
        return $this->hasMany(MakeupSaleCostRelate::class, 'sale_id', 'id');
    }
}
