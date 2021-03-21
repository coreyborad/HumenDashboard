<?php

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', function () {
    return 'API';
});

// Route::get('/{path}', function () {
//     return view('index');
// })->where('path', '^((?!api).)*$');

// Route::any('/.well-known/acme-challenge/{all}', function () {
//     return 'ok';
// });

// Route::any('{all}', function () {
//     return view('index');
// })->where(['all' => '.*']);
