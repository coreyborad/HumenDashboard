<?php

namespace App\Services;

use App\Repositories\UserRepository;

class UserService
{
    protected $userRepository;

    public function __construct(GroupRepository $userRepository)
    {
        $this->userRepository = $userRepository;
    }

    public function fetchUser(int $id)
    {
        return $this->userRepository->find($id)->user()->get();
    }
}
