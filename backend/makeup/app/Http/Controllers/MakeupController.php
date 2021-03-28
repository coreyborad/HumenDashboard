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
            $param = ['brand', 'name', 'color_name'];
            $query_string = $request->only($param);
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
        $param = ['price', 'order_date'];
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
        $param = ['price', 'sold_date'];
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

    public function getReports(Request $request){
        $data = [];
        // Option-a 月份範圍收支額(全部商品)
        $param = ['type', 'target', 'date_start', 'date_end'];
        if ($request->has($param) && $request->input('type') === 'a') {
            try {
                $query_string = $request->only($param);
                $start = Carbon::parse($query_string['date_start']);
                $end = Carbon::parse($query_string['date_end'])->endOfMonth();
                $diff_months = $start->diffInMonths($end);
                if ($diff_months >= 12 || $diff_months <= 0){
                    throw new ErrorException(400, $diff_months);
                }
                switch ($query_string['target']) {
                    case 'cost':
                        $data = $this->service->getMakeupCostByDate($start, $end);
                        break;
                    case 'sale':
                        $data = $this->service->getMakeupSaleByDate($start, $end);
                        break;
                }
                
            } catch (\Exception $e) {
                throw $e;
            }
        }
        // Option-b 單月份銷售狀況(以商品分類)
        $param = ['type', 'target', 'date'];
        if ($request->has($param) && $request->input('type') === 'b') {
            try {
                $query_string = $request->only($param);
                $date = Carbon::parse($query_string['date']);
                switch ($query_string['target']) {
                    case 'cost':
                        $data = $this->service->getMakeupCostGroupByItemOnDate($date);
                        break;
                    case 'sale':
                        $data = $this->service->getMakeupCostGroupByItemOnDate($date);
                        break;
                }
            } catch (\Exception $e) {
                throw $e;
            }
        }
        // Option-c 月份範圍商品淨損狀況
        $param = ['type', 'date_start', 'date_end'];
        if ($request->has($param) && $request->input('type') === 'c') {
            try {
                $query_string = $request->only($param);
                $start = Carbon::parse($query_string['date_start']);
                $end = Carbon::parse($query_string['date_end'])->endOfMonth();
                $diff_months = $start->diffInMonths($end);
                if ($diff_months >= 12 || $diff_months <= 0){
                    throw new ErrorException(400, $diff_months);
                }
                $data = $this->service->getMakeupRealSaleReportByDate($start, $end);
                
            } catch (\Exception $e) {
                throw $e;
            }
        }
        // Option-d 月份範圍商品淨損狀況
        $param = ['type', 'date'];
        if ($request->has($param) && $request->input('type') === 'd') {
            try {
                $query_string = $request->only($param);
                $date = Carbon::parse($query_string['date']);
                $data = $this->service->getMakeupSaleCountReportByMonth($date);
            } catch (\Exception $e) {
                throw $e;
            }
        }
        return response()->json($data);
    }

    public function getMakeupInventory(Request $request, int $id){

        try {
            $data = $this->service->getMakeupInventory($id);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }
    
}
