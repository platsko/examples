<?php

use Illuminate\Database\Seeder;

class WalletsTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        DB::table('wallets')->insert(
            [
                [
                    'type' => 'vm',
                ],
                [
                    'type' => 'user',
                ],
                [
                    'type' => 'deposit',
                ],
            ]
        );
    }
}
