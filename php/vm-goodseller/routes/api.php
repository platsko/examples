<?php

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::get('/refresh',     'VmController@refresh')
    ->name('refresh');

Route::post('/putcoin',     'VmController@putCoin')
    ->name('putcoin');

Route::post('/purchase',    'VmController@purchase')
    ->name('purchase');

Route::post('/refund',      'VmController@refund')
    ->name('refund');
