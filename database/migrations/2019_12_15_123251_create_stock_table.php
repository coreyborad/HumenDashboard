<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateStockTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('stock', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->string('stock_number'); // 股票編號
            $table->bigInteger('user_id'); // User id
            $table->integer('shares'); // 持有股數
            $table->integer('cost'); // 持有成本
            $table->integer('price'); // 目前價位 
            $table->timestamps();
        });

        Schema::create('stock_history', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->string('stock_number'); // 股票編號
            $table->bigInteger('user_id'); // User id
            $table->integer('total_cost'); // 當天總成本
            $table->integer('total_value'); // 當天總價值
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('stock');
    }
}
