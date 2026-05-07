<?php

declare(strict_types=1);

namespace App\Services;

use Firebase\JWT\JWT;

class AuthService
{
    public function attempt(string $username, string $password): ?string
    {
        $validUsername = $_ENV['APP_USERNAME'] ?? 'admin';
        $validPassword = $_ENV['APP_PASSWORD'] ?? 'admin';

        if ($username !== $validUsername || $password !== $validPassword) {
            return null;
        }

        return $this->generateToken($username);
    }

    private function generateToken(string $username): string
    {
        $now     = time();
        $expires = (int) ($_ENV['JWT_EXPIRES'] ?? 3600);
        $secret  = $_ENV['JWT_SECRET'];

        $payload = [
            'iat'      => $now,
            'exp'      => $now + $expires,
            'username' => $username,
        ];

        return JWT::encode($payload, $secret, 'HS256');
    }
}
