<?php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use Slim\Factory\AppFactory;
use DI\Container;
use Dotenv\Dotenv;

$dotenv = Dotenv::createImmutable(__DIR__ . '/..');
$dotenv->load();

$container = new Container();
AppFactory::setContainer($container);

// Register MongoDB
$container->set('db', function () {
    $client = new MongoDB\Client($_ENV['MONGO_URI']);
    return $client->selectDatabase($_ENV['MONGO_DB']);
});

$app = AppFactory::create();
$app->addBodyParsingMiddleware();
$app->addErrorMiddleware(true, true, true);

require __DIR__ . '/../app/Routes/api.php';

$app->run();
