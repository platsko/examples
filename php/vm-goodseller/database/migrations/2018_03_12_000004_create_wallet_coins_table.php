<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateWalletCoinsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('wallet_coins', function (Blueprint $table) {

            $table->unsignedInteger('wallet_id');
            $table->unsignedInteger('coin_id');
            $table->unsignedInteger('quantity');

            $table->foreign('wallet_id')->references('id')->on('wallets');
            $table->foreign('coin_id')->references('id')->on('coins');

        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('wallet_coins');
    }
}
