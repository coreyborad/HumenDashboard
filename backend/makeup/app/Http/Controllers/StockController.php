<?php

namespace App\Http\Controllers;

use Auth;
use App\Http\Controllers\Controller;
use App\Services\StockService;
use Illuminate\Http\Request;
use App\Exceptions\ErrorException;

class StockController extends Controller
{
    protected $service;

    public function __construct(StockService $service)
    {
        $this->service = $service;
    }

    public function getUserStock(Request $request){
        try {
            $user_id = Auth::id();
            $data = $this->service->getUserStockList($user_id);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function createUserStock(Request $request){
        $param = ['stock_number', 'shares', 'cost'];
        if (!$request->has($param)) {
            throw new ErrorException(400, 'error');
        }

        try {
            $user_id = Auth::id();
            $stock = $request->only($param);
            $data = $this->service->createUserStock($user_id, $stock);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function deleteUserStock(Request $request, int $id){

        try {
            $data = $this->service->deleteUserStock($id);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function getStockList(Request $request) {
        try {
            $data = $this->service->getStockList();
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function updateUserStock(Request $request, int $id){
        $param = ['stock_number', 'shares', 'cost'];
        if (!$request->has($param)) {
            throw new ErrorException(400, 'error');
        }
        try {
            $info = $request->only($param);
            $data = $this->service->updateUserStock($id, $info);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }
}
