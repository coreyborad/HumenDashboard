<?php

namespace App\Http\Controllers;

use Auth;
use App\Http\Controllers\Controller;
use App\Services\MakeupService;
use Illuminate\Http\Request;
use App\Exceptions\ErrorException;
use Carbon\Carbon;

class MakeupController extends Controller
{
    protected $service;

    public function __construct(MakeupService $service)
    {
        $this->service = $service;
    }

    public function getMakeup(Request $request){
        try {
            $user_id = Auth::id();
            $data = $this->service->getUserMakeupList($user_id);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function getMakeupByQuery(Request $request){
        try {
            $query_string = $request->query();
            $data = $this->service->getMakeupByQuery($query_string);
        } catch (\Exception $e) {
            throw $e;
        }
        return response()->json($data);
    }

    public function createMakeupInfo(Request $request){
        $param = ['brand', 'name', 'color_name'];
        if (!$request->has($param)) {
            throw new ErrorException(400, 'error');
        }

        try {
            $makeup_info = $request->only($param);
            $data = $this->service->createMakeupInfo($makeup_info);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function deleteMakeupInfo(Request $request, int $id){

        try {
            $data = $this->service->deleteMakeupInfo($id);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function updateMakeupInfo(Request $request, int $id){
        $param = ['brand', 'name', 'color_name'];
        if (!$request->has($param)) {
            throw new ErrorException(400, 'error');
        }
        try {
            $makeup_info = $request->only($param);
            $data = $this->service->updateMakeupInfo($id, $makeup_info);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function createMakeupCost(Request $request){
        $param = ['makeup_id', 'price', 'count', 'order_date'];
        if (!$request->has($param)) {
            throw new ErrorException(400, 'error');
        }

        try {
            $cost_info = $request->only($param);
            $data = $this->service->createMakeupCost($cost_info);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function deleteMakeupCost(Request $request, int $id){

        try {
            $data = $this->service->deleteMakeupCost($id);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function updateMakeupCost(Request $request, int $id){
        $param = ['price', 'count', 'order_date'];
        if (!$request->has($param)) {
            throw new ErrorException(400, 'error');
        }
        try {
            $cost_info = $request->only($param);
            $data = $this->service->updateMakeupCost($id, $cost_info);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function createMakeupSale(Request $request){
        $param = ['makeup_id', 'price', 'count', 'sold_date'];
        if (!$request->has($param)) {
            throw new ErrorException(400, 'error');
        }

        try {
            $makeup_info = $request->only($param);
            $data = $this->service->createMakeupSale($makeup_info);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function deleteMakeupSale(Request $request, int $id){

        try {
            $data = $this->service->deleteMakeupSale($id);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function updateMakeupSale(Request $request, int $id){
        $param = ['price', 'count', 'sold_date'];
        if (!$request->has($param)) {
            throw new ErrorException(400, 'error');
        }
        try {
            $sale_info = $request->only($param);
            $data = $this->service->updateMakeupSale($id, $sale_info);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function getCostByMonthRange(Request $request){
        $param = ['date_start', 'date_end'];
        if (!$request->has($param)) {
            throw new ErrorException(400, 'error');
        }
        try {
            $query_string = $request->only($param);
            $start = Carbon::parse($query_string['date_start']);
            $end = Carbon::parse($query_string['date_end']);
            $diff_months = $start->floatDiffInMonths($end);
            if ($diff_months >= 12 || $diff_months <= 0){
                throw new ErrorException(400, $diff_months);
            }
            $data = $this->service->getMakeupCostByDate($start, $end);
        } catch (\Exception $e) {
            throw $e;
        }
        return response()->json($data);
    }

    public function getSaleByMonthRange(Request $request){
        $param = ['date_start', 'date_end'];
        if (!$request->has($param)) {
            throw new ErrorException(400, 'error');
        }
        try {
            $query_string = $request->only($param);
            $start = Carbon::parse($query_string['date_start']);
            $end = Carbon::parse($query_string['date_end']);
            $diff_months = $start->floatDiffInMonths($end);
            if ($diff_months >= 12 || $diff_months <= 0){
                throw new ErrorException(400, $diff_months);
            }
            $data = $this->service->getMakeupSaleByDate($start, $end);
        } catch (\Exception $e) {
            throw $e;
        }
        return response()->json($data);
    }
}
