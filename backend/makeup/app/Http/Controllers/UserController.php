<?php

namespace App\Http\Controllers;

use Auth;
use App\Http\Controllers\Controller;
use App\Services\UserService;
use Illuminate\Http\Request;
use App\Exceptions\ErrorException;

class UserController extends Controller
{
    protected $service;

    public function __construct(UserService $service)
    {
        $this->service = $service;
    }

    public function getUserInfo(Request $request){
        try {
            $user_id = Auth::id();
            $data = $this->service->fetchUser($user_id);
        } catch (\Exception $e) {
            throw $e;
        }

        return response()->json($data);
    }

}
