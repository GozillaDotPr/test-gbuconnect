<?php

declare(strict_types=1);

namespace App\Models;

class Product
{
    public string $id;
    public string $name;
    public string $desc;
    public int $price;
    public string $user_id;
    public string $created_at;

    public function __construct(array $data)
    {
        $this->id         = $data['id'] ?? '';
        $this->name       = $data['name'] ?? '';
        $this->desc       = $data['desc'] ?? '';
        $this->price      = (int) ($data['price'] ?? 0);
        $this->user_id    = $data['user_id'] ?? '';
        $this->created_at = $data['created_at'] ?? '';
    }

    public function toArray(): array
    {
        return [
            'id'         => $this->id,
            'name'       => $this->name,
            'desc'       => $this->desc,
            'price'      => $this->price,
            'user_id'    => $this->user_id,
            'created_at' => $this->created_at,
        ];
    }
}
