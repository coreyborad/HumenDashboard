<?php

namespace App\Http\Controllers;

use Auth;
use App\Http\Controllers\Controller;
use App\Services\MakeupService;
use Illuminate\Http\Request;
use App\Exceptions\ErrorException;

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

    public function createMakeup(Request $request){
        $param = ['Makeup_number', 'shares', 'cost'];
        if (!$request->has($param)) {
            throw new ErrorException(400, 'error');
        }

        try {
            $user_id = Auth::id();
            $Makeup = $request->only($param);
            $data = $this->service->createUserMakeup($user_id, $Makeup);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

    public function deleteMakeup(Request $request, int $id){

        try {
            $data = $this->service->deleteMakeup($id);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }
}
