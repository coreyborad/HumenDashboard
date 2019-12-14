<?php

namespace App\Http\Middleware;

use Auth;
use Closure;
use Tymon\JWTAuth\Exceptions\JWTException;
use Tymon\JWTAuth\Http\Middleware\BaseMiddleware;
use Tymon\JWTAuth\Exceptions\TokenExpiredException;
use App\Exceptions\UnauthorizedTokenException;

class AuthenticateToken extends BaseMiddleware
{
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure  $next
     * @return mixed
     */
    public function handle($request, Closure $next)
    {
        $this->checkForToken($request);

        try {
            if ($this->auth->parseToken()->authenticate()) {
                return $next($request);
            }

            throw new UnauthorizedTokenException('User not found');
        } catch (TokenExpiredException $exception) {
            try {
                $token = $this->auth->refresh();

                Auth::onceUsingId($this->auth->manager()->getPayloadFactory()->buildClaimsCollection()->toPlainArray()['sub']);
            } catch (JWTException $exception) {
                throw new UnauthorizedTokenException($exception->getMessage());
            }
        }

        // Send the refreshed token back to the client.
        return $this->setAuthenticationHeader($next($request), $token);
    }
}
