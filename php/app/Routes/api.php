<?php

declare(strict_types=1);

use App\Handlers\AuthHandler;
use App\Handlers\ProductHandler;
use App\Middleware\JwtMiddleware;
use App\Repositories\ProductRepository;
use App\Services\ProductService;
use App\Services\AuthService;

// Auth
$app->post('/api/v1/login', function ($request, $response) use ($app) {
    $handler = new AuthHandler(new AuthService());
    return $handler->login($request, $response);
});

// Products (protected)
$app->group('/api/v1/products', function ($group) {
    $group->get('', function ($request, $response) {
        $handler = new ProductHandler(new ProductService(new ProductRepository($this->get('db'))));
        return $handler->getAll($request, $response);
    });

    $group->get('/{id}', function ($request, $response, $args) {
        $handler = new ProductHandler(new ProductService(new ProductRepository($this->get('db'))));
        return $handler->getById($request, $response, $args);
    });

    $group->post('', function ($request, $response) {
        $handler = new ProductHandler(new ProductService(new ProductRepository($this->get('db'))));
        return $handler->create($request, $response);
    });

    $group->put('/{id}', function ($request, $response, $args) {
        $handler = new ProductHandler(new ProductService(new ProductRepository($this->get('db'))));
        return $handler->update($request, $response, $args);
    });

    $group->delete('/{id}', function ($request, $response, $args) {
        $handler = new ProductHandler(new ProductService(new ProductRepository($this->get('db'))));
        return $handler->delete($request, $response, $args);
    });
})->add(new JwtMiddleware());
