<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class MakeupInfo extends Model
{
    protected $connection = 'mysql';
    protected $table = 'makeup_info';
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'brand',
        'name',
        'color_name',
    ];

    /**
     * The attributes that should be cast to native types.
     *
     * @var array
     */
    protected $casts = [];

    public function costs()
    {
        return $this->hasMany(MakeupCost::class, 'makeup_id', 'id');
    }

    public function sales()
    {
        return $this->hasMany(MakeupSale::class, 'makeup_id', 'id');
    }
}
