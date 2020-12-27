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

        Schema::create('makeup_price', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->bigInteger('makeup_id');
            $table->float('cost_price');
            $table->float('sale_price');
            $table->integer('inventory_count');
            $table->integer('sold_count');
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
        Schema::dropIfExists('makeup');
        Schema::dropIfExists('makeup_price');
    }
}
