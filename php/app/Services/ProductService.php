<?php

declare(strict_types=1);

namespace App\Services;

use App\Repositories\ProductRepository;
use Ramsey\Uuid\Uuid;

class ProductService
{
    public function __construct(private ProductRepository $repo) {}

    public function getAll(): array
    {
        return $this->repo->findAll();
    }

    public function getById(string $id): ?array
    {
        return $this->repo->findById($id);
    }

    public function create(array $body, string $userId): array
    {
        $product = [
            'id'         => Uuid::uuid4()->toString(),
            'name'       => trim($body['name']),
            'desc'       => trim($body['desc'] ?? ''),
            'price'      => (int) ($body['price'] ?? 0),
            'user_id'    => $userId,
            'created_at' => (new \DateTime())->format(\DateTime::ATOM),
        ];

        return $this->repo->create($product);
    }

    public function update(string $id, array $body): bool
    {
        $existing = $this->repo->findById($id);
        if (!$existing) {
            return false;
        }

        $fields = array_filter([
            'name'  => isset($body['name']) ? trim($body['name']) : null,
            'desc'  => isset($body['desc']) ? trim($body['desc']) : null,
            'price' => isset($body['price']) ? (int) $body['price'] : null,
        ], fn($v) => $v !== null);

        if (empty($fields)) {
            return true;
        }

        return $this->repo->update($id, $fields);
    }

    public function delete(string $id): bool
    {
        return $this->repo->delete($id);
    }
}
