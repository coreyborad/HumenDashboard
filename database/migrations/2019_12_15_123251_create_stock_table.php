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
        Schema::create('user_has_stock', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->bigInteger('user_id'); // User id
            $table->string('stock_number'); // 股票編號
            $table->integer('shares'); // 持有股數
            $table->integer('cost'); // 持有成本
            $table->timestamps();
        });

        Schema::create('stock_history', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->string('stock_number'); // 股票編號
            $table->date('deal_date'); //資料日期
            $table->bigInteger('deal_count'); // 成交股數
            $table->bigInteger('deal_price'); // 成交金額
            $table->bigInteger('deal_record'); // 成交筆數
            $table->float('price_on_open'); // 開盤價
            $table->float('price_on_highest'); // 最高價
            $table->float('price_on_lowest'); // 最低價
            $table->float('price_on_close'); // 收盤價
            $table->float('price_diff'); // 漲跌價差
            $table->timestamps();

            $table->unique(['stock_number', 'deal_date']); // 當天同一筆股票只能有一筆資料
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('user_has_stock');
        Schema::dropIfExists('stock_history');
    }
}
