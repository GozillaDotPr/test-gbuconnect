<?php

declare(strict_types=1);

namespace App\Handlers;

use App\Services\ProductService;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

class ProductHandler
{
    public function __construct(private ProductService $productService) {}

    public function getAll(Request $request, Response $response): Response
    {
        $products = $this->productService->getAll();
        return $this->json($response, ['success' => true, 'data' => $products]);
    }

    public function getById(Request $request, Response $response, array $args): Response
    {
        $product = $this->productService->getById($args['id']);

        if (!$product) {
            return $this->json($response, ['success' => false, 'message' => 'Product not found'], 404);
        }

        return $this->json($response, ['success' => true, 'data' => $product]);
    }

    public function create(Request $request, Response $response): Response
    {
        $body = (array) $request->getParsedBody();

        if (empty($body['name']) || !isset($body['price'])) {
            return $this->json($response, ['success' => false, 'message' => 'name and price are required'], 400);
        }

        // Get user_id from JWT payload stored in request attribute
        $token  = $request->getAttribute('token');
        $userId = $token->username ?? 'unknown';

        $product = $this->productService->create($body, $userId);

        return $this->json($response, ['success' => true, 'data' => $product], 201);
    }

    public function update(Request $request, Response $response, array $args): Response
    {
        $body = (array) $request->getParsedBody();

        $updated = $this->productService->update($args['id'], $body);

        if (!$updated) {
            return $this->json($response, ['success' => false, 'message' => 'Product not found'], 404);
        }

        $product = $this->productService->getById($args['id']);
        return $this->json($response, ['success' => true, 'data' => $product]);
    }

    public function delete(Request $request, Response $response, array $args): Response
    {
        $deleted = $this->productService->delete($args['id']);

        if (!$deleted) {
            return $this->json($response, ['success' => false, 'message' => 'Product not found'], 404);
        }

        return $this->json($response, ['success' => true, 'data' => null]);
    }

    private function json(Response $response, array $data, int $status = 200): Response
    {
        $response->getBody()->write(json_encode($data));
        return $response->withHeader('Content-Type', 'application/json')->withStatus($status);
    }
}
