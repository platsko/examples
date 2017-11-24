<?php

use Faker\Generator as Faker;

/*
|--------------------------------------------------------------------------
| Model Factories
|--------------------------------------------------------------------------
|
| This directory should contain each of the model factory definitions for
| your application. Factories provide a convenient way to generate new
| model instances for testing / seeding your application's database.
|
*/

$factory->define(App\Models\Goods::class, function (Faker $faker) {
    return [
        'id'       => $faker->numberBetween(1, 1111),
        'name'     => $faker->name,
        'quantity' => $faker->numberBetween(1, 1111),
        'cost'     => $faker->numberBetween(1, 1111),
    ];
});
