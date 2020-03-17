### TAXI TEST APP

## Description
Simple test application which generating random orders and shows it by calling endpoints

## Deploy

1) Install docker
2) Insta docker-compose

Every that you need is deploying via command:

`$ docker-compose up --build -d`

The environment variables are in docker compose file.

## Endpoints

#### Test
```
localhost:8081/
```

#### Get one order
```
localhost:8081/order
```

#### Get all orders
```
localhost:8081/orders
```