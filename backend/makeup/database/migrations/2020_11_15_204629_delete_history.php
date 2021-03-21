<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class DeleteHistory extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::dropIfExists('stock_history');
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::create('stock_history', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->string('stock_number'); // 股票編號
            $table->date('deal_date'); //資料日期
            $table->bigInteger('deal_count'); // 成交股數
            $table->float('price_on_open'); // 開盤價
            $table->float('price_on_highest'); // 最高價
            $table->float('price_on_lowest'); // 最低價
            $table->float('price_on_close'); // 收盤價
            $table->timestamps();

            $table->unique(['stock_number', 'deal_date']); // 當天同一筆股票只能有一筆資料
        });
    }
}
