<?php

namespace App\Http\Controllers;

use Auth;
use App\Http\Controllers\Controller;
use App\Services\StockService;
use Illuminate\Http\Request;

class StockController extends Controller
{
    protected $service;

    public function __construct(StockService $service)
    {
        $this->service = $service;
    }

    public function getStock(Request $request){
        try {
            $user_id = Auth::id();
            $data = $this->service->getUserStockList($user_id);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }
}
