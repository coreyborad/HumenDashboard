<?php

use Illuminate\Http\Request;

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

// Route::middleware('auth:api')->get('/user', function (Request $request) {
//     return $request->user();
// });

Route::group(['prefix' => 'v1'], function () {
    Route::post('signup', 'AuthController@signup');
    Route::post('login', 'AuthController@login');
    Route::middleware('auth:api')->group(function () {
        Route::post('logout', 'AuthController@logout');
        // User
        Route::group(['prefix' => 'user'], function () {
            Route::get('', 'UserController@getUserInfo');
        });
        // Stock
        // User
        Route::group(['prefix' => 'stock'], function () {
            Route::get('', 'StockController@getStockList');
        });
        Route::group(['prefix' => 'user_stock'], function () {
            Route::get('', 'StockController@getUserStock');
            Route::post('', 'StockController@createUserStock');
            Route::delete('{id}', 'StockController@deleteUserStock');
        });

        // Makeup
        Route::group(['prefix' => 'makeup'], function () {
            Route::get('', 'StockController@getMakeup');
            Route::post('', 'StockController@createMakeup');
            Route::delete('{id}', 'StockController@deleteMakeup');
        });
    });
});
