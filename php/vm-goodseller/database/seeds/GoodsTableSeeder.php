<?php

use Illuminate\Database\Seeder;

class GoodsTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        DB::table('goods')->insert(
            [
                [
                    'name'     => 'Чай',
                    'quantity' => 10,
                    'cost'     => 13,
                ],
                [
                    'name'     => 'Кофе',
                    'quantity' => 20,
                    'cost'     => 18,
                ],
                [
                    'name'     => 'Кофе с молоком',
                    'quantity' => 20,
                    'cost'     => 21,
                ],
                [
                    'name'     => 'Сок',
                    'quantity' => 15,
                    'cost'     => 35,
                ],
            ]
        );
    }
}
