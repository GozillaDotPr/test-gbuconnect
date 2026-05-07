<?php

declare(strict_types=1);

namespace App\Handlers;

use App\Services\AuthService;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

class AuthHandler
{
    public function __construct(private AuthService $authService) {}

    public function login(Request $request, Response $response): Response
    {
        $body     = (array) $request->getParsedBody();
        $username = trim($body['username'] ?? '');
        $password = trim($body['password'] ?? '');

        if (!$username || !$password) {
            return $this->json($response, ['success' => false, 'message' => 'Username and password required'], 400);
        }

        $token = $this->authService->attempt($username, $password);

        if (!$token) {
            return $this->json($response, ['success' => false, 'message' => 'Invalid credentials'], 401);
        }

        return $this->json($response, [
            'success' => true,
            'data'    => ['token' => $token],
        ]);
    }

    private function json(Response $response, array $data, int $status = 200): Response
    {
        $response->getBody()->write(json_encode($data));
        return $response->withHeader('Content-Type', 'application/json')->withStatus($status);
    }
}
