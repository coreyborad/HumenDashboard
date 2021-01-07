<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateMakeupTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('makeup_info', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->string('brand');
            $table->string('name');
            $table->string('color_name');
            $table->timestamps();

            $table->unique(['brand', 'name', 'color_name']);
        });
        Schema::create('makeup_cost', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->bigInteger('makeup_id');
            $table->float('price');
            $table->integer('count');
            $table->timestamp('order_date');
        });

        Schema::create('makeup_sale', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->bigInteger('makeup_id');
            $table->bigInteger('makeup_cost_id');
            $table->float('price');
            $table->integer('count');
            $table->timestamp('sold_date');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('makeup_info');
        Schema::dropIfExists('makeup_cost');
        Schema::dropIfExists('makeup_sale');
    }
}
