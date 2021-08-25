# apitask

REST API for creating orders and sending email reminders of unfinished orders, written in Go.

## Initializing database (development)

```bash
cd server
go run . init
```

## Starting application (development)

```bash
cd server
go run . start
```

## Starting application (production)

```bash
docker build -t apitask .
docker run -d -p 3000:3000 apitask
```

## API Documentation

### POST /customer/new

Creates a new customer

#### Example request

``POST /customer/new``

Request body:
```json
{
    "name":"James",
    "email":"hello@example.com"
}
```

#### Example response

```json
{
    "customer":{
        "customerID":7,
        "name":"James",
        "email":"hello@example.com"
    },
    "status":201,
    "success":true
}
```

### POST /product/new

Creates a new product

#### Example request

``POST /product/new``

Request body:
```json
{
    "name":"Computer",
    "price":10000
}
```

#### Example response

```json
{
    "product":{
        "productID":1,
        "name":"Computer",
        "priceCZK":10000
    },
    "status":201,
    "success":true
}
```

### POST /order/new

Creates a new order

#### Example request

``POST /order/new``

Request body:
```json
{
    "customer":7,
}
```

#### Example response

```json
{
    "order":{
        "orderID":1,
        "items":[
            
        ],
        "customer":{
            "customerID":7,
            "name":"James",
            "email":"hello@example.com"
        },
        "createdAt":"2021-08-25T17:40:48.8443813+02:00",
        "confirmed":false,
        "reminded":false
    },
    "status":201,
    "success":true
}
```

### PUT /order/add

Adds a product into an order

#### Example request

``PUT /order/add``

Request body:
```json
{
    "orderid":1,
    "product":1,
    "amount":5
}
```

#### Example response

```json
{
    "orderItem":{
        "orderItemID":1,
        "product":{
            "productID":1,
            "name":"Computer",
            "priceCZK":10000
        },
        "amount":5
    },
    "status":201,
    "success":true
}
```

### GET /order/{orderid}

Returns complete order information

#### Example request

``GET /order/1``

#### Example response

```json
{
    "order":{
        "orderID":1,
        "items":[
            {
                "orderItemID":1,
                "product":{
                    "productID":1,
                    "name":"Computer",
                    "priceCZK":10000
                },
                "amount":5
            }
        ],
        "customer":{
            "customerID":7,
            "name":"James",
            "email":"hello@example.com"
        },
        "createdAt":"2021-08-25T17:40:48.8443813+02:00",
        "confirmed":false,
        "reminded":false
    },
    "status":200,
    "success":true
}
```

### PUT /orderitem/amount

Changes order item amount

#### Example request

``PUT /orderitem/amount``

Request body:
```json
{
    "itemid":1,
    "amount":2
}
```

#### Example response

```json
{
    "status":200,
    "success":true
}
```

### PUT /order/confirm

Confirms order

#### Example request

``PUT /order/confirm``

Request body:
```json
{
    "orderid":1
}
```

#### Example response

```json
{
    "status":200,
    "success":true
}
```

### DELETE /orderitem

Deletes order item

#### Example request

``DELETE /orderitem``

Request body:
```json
{
    "itemid":1
}
```

#### Example response

```json
{
    "status":200,
    "success":true
}
```