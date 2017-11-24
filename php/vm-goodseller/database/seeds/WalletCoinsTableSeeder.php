<?php

use Illuminate\Database\Seeder;

class WalletCoinsTableSeeder extends Seeder
{
    protected $coins;

    protected $wallets;

    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $this->coins = DB::table('coins')->get()->keyBy('cost');
        $this->wallets = DB::table('wallets')->get()->keyBy('type');

        $this->seedVmWallet();
        $this->seedUserWallet();
        $this->seedDepositWallet();
    }

    protected function seedVmWallet()
    {
        $id = $this->wallets->get('vm')->id;

        DB::table('wallet_coins')->insert(
            [
                [
                    'wallet_id' => $id,
                    'coin_id' => $this->coins->get('1')->id,
                    'quantity' => 100,
                ],
                [
                    'wallet_id' => $id,
                    'coin_id' => $this->coins->get('2')->id,
                    'quantity' => 100,
                ],
                [
                    'wallet_id' => $id,
                    'coin_id' => $this->coins->get('5')->id,
                    'quantity' => 100,
                ],
                [
                    'wallet_id' => $id,
                    'coin_id' => $this->coins->get('10')->id,
                    'quantity' => 100,
                ],
            ]
        );
    }

    protected function seedUserWallet()
    {
        $id = $this->wallets->get('user')->id;

        DB::table('wallet_coins')->insert(
            [
                [
                    'wallet_id' => $id,
                    'coin_id' => $this->coins->get('1')->id,
                    'quantity' => 10,
                ],
                [
                    'wallet_id' => $id,
                    'coin_id' => $this->coins->get('2')->id,
                    'quantity' => 30,
                ],
                [
                    'wallet_id' => $id,
                    'coin_id' => $this->coins->get('5')->id,
                    'quantity' => 20,
                ],
                [
                    'wallet_id' => $id,
                    'coin_id' => $this->coins->get('10')->id,
                    'quantity' => 15,
                ],
            ]
        );
    }

    protected function seedDepositWallet()
    {
        $id = $this->wallets->get('deposit')->id;

        DB::table('wallet_coins')->insert(
            [
                [
                    'wallet_id' => $id,
                    'coin_id' => $this->coins->get('1')->id,
                    'quantity' => 0,
                ],
                [
                    'wallet_id' => $id,
                    'coin_id' => $this->coins->get('2')->id,
                    'quantity' => 0,
                ],
                [
                    'wallet_id' => $id,
                    'coin_id' => $this->coins->get('5')->id,
                    'quantity' => 0,
                ],
                [
                    'wallet_id' => $id,
                    'coin_id' => $this->coins->get('10')->id,
                    'quantity' => 0,
                ],
            ]
        );
    }
}
