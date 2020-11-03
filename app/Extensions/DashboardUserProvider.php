<?php

namespace App\Extensions;

use Illuminate\Support\Str;
use Illuminate\Contracts\Auth\UserProvider;
use Illuminate\Contracts\Support\Arrayable;
use Illuminate\Contracts\Auth\Authenticatable as UserContract;

class DashboardUserProvider implements UserProvider
{
    protected $model;

    public function __construct($model)
    {
        $this->model = $model;
    }

    public function retrieveById($identifier)
    {
        $model = $this->createModel();

        return $this->newModelQuery($model)
                    ->where($model->getAuthIdentifierName(), $identifier)
                    ->first();
    }

    public function retrieveByToken($identifier, $token)
    {
        return null;
    }

    public function updateRememberToken(UserContract $user, $token)
    {
        return;
    }

    public function retrieveByCredentials(array $credentials)
    {
        if (empty($credentials) ||
           (count($credentials) === 1 &&
            array_key_exists('password', $credentials))) {
            return;
        }

        $query = $this->newModelQuery();

        foreach ($credentials as $key => $value) {
            if (Str::contains($key, 'password')) {
                continue;
            }

            if (is_array($value) || $value instanceof Arrayable) {
                $query->whereIn($key, $value);
            } else {
                $query->where($key, $value);
            }
        }

        return $query->first();
    }

    public function validateCredentials(UserContract $user, array $credentials)
    {
        $cypher = sha1($credentials['password']);

        return hash_equals($cypher, strtolower($user->getAuthPassword()));
    }

    protected function newModelQuery($model = null)
    {
        return is_null($model)
                ? $this->createModel()->newQuery()
                : $model->newQuery();
    }

    public function createModel()
    {
        $class = '\\'.ltrim($this->model, '\\');

        return new $class;
    }

    public function getModel()
    {
        return $this->model;
    }

    public function setModel($model)
    {
        $this->model = $model;

        return $this;
    }
}
