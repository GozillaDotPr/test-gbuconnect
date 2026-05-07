<?php

declare(strict_types=1);

namespace App\Repositories;

use MongoDB\Database;
use MongoDB\BSON\UTCDateTime;

class ProductRepository
{
    private $collection;

    public function __construct(Database $db)
    {
        $this->collection = $db->selectCollection('products');
    }

    public function findAll(): array
    {
        $cursor = $this->collection->find([], ['sort' => ['created_at' => -1]]);
        $results = [];
        foreach ($cursor as $doc) {
            $results[] = $this->normalize($doc);
        }
        return $results;
    }

    public function findById(string $id): ?array
    {
        $doc = $this->collection->findOne(['id' => $id]);
        return $doc ? $this->normalize($doc) : null;
    }

    public function create(array $data): array
    {
        $this->collection->insertOne($data);
        return $data;
    }

    public function update(string $id, array $data): bool
    {
        $result = $this->collection->updateOne(
            ['id' => $id],
            ['$set' => $data]
        );
        return $result->getMatchedCount() > 0;
    }

    public function delete(string $id): bool
    {
        $result = $this->collection->deleteOne(['id' => $id]);
        return $result->getDeletedCount() > 0;
    }

    private function normalize($doc): array
    {
        $arr = (array) $doc->bsonSerialize();
        unset($arr['_id']);
        return $arr;
    }
}
