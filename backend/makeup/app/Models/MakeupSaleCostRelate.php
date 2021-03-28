<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class MakeupSaleCostRelate extends Model
{
    protected $connection = 'mysql';
    protected $table = 'sale_cost_relations';
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'cost_id',
        'sale_id',
        'relate_count',
    ];

    /**
     * The attributes that should be cast to native types.
     *
     * @var array
     */


    public function cost()
    {
        return $this->hasOne(MakeupCost::class, 'id', 'cost_id');
    }
}
