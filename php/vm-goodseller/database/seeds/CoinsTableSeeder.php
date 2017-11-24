<?php

use Illuminate\Database\Seeder;

class CoinsTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        DB::table('coins')->insert(
            [
                [
                    'cost' => 1,
                ],
                [
                    'cost' => 2,
                ],
                [
                    'cost' => 5,
                ],
                [
                    'cost' => 10,
                ],
            ]
        );
    }
}
