<?php

namespace App\Exceptions;

use Symfony\Component\HttpKernel\Exception\HttpException;
use Throwable;

class ErrorException extends HttpException
{
    private $errorCode;

    public function __construct(int $statusCode, string $errorCode = null, string $message = null, Throwable $previous = null, ?int $code = 0, array $headers = [])
    {
        $this->errorCode = $errorCode;

        parent::__construct($statusCode, $message, $previous, $headers, $code);
    }

    public function render($request)
    {
        return response([
                'error' => $this->errorCode,
                'error_description' => $this->getMessage(),
            ], $this->getStatusCode())
            ->withHeaders([
                'Cache-Control' =>'no-store',
                'Pragma' => 'no-cache',
            ]);
    }
}
