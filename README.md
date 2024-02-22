# Hexagonal Architecture in Go

## Overview
This is an example of an application implementing Hexagonal Architecture in Go.

## Used Technologies
- [Go](https://golang.org/): Main programming language.
- [Docker](https://www.docker.com/): For managing application containers.
- [SQLite](https://www.sqlite.org/index.html): Database used for storing products.
- [Gorilla Mux](https://github.com/gorilla/mux): HTTP router used for managing API routes.
- [Cobra](https://github.com/spf13/cobra): Library for easily creating CLI applications.
- [Testify](https://github.com/stretchr/testify): Testing framework for Go.
- [Negroni](https://github.com/urfave/negroni): HTTP middleware used for the application layer.
- [Mockery](https://github.com/vektra/mockery): Tool for generating mocks in Go.
- [UUID](https://github.com/google/uuid): For generating unique identifiers.
- [GoValidator](https://github.com/asaskevich/govalidator): Library for data validation in Go.
- [Uber Go Mock](https://github.com/uber-go/goleak): Library for creating mocks in tests.

## Usage Instructions
Follow these instructions to run and test the application locally.

### Prepare environment
```bash
docker-compose up -d
docker exec -it  appproduct bash
```

### Create sqlite database
```bash
touch sqlite3.db
sqlite3 sqlite3.db
CREATE TABLE products (id VARCHAR(36) PRIMARY KEY,name VARCHAR(255),price FLOAT,status VARCHAR(255));
.tables
```


### Generate mocks
```bash
mockgen -destination=application/mocks/application.go -source=application/product.go application
```

### Run tests with coverage and generate html report
```bash
go test ./... -coverprofile=coverage.out && grep -v "mocks/" coverage.out | go tool cover -html=/dev/stdin -o coverage.html && rm coverage.out
``` 

### Execute web server
```bash
go run main.go http
```

### Request examples


#### Create a product
```bash
curl --location 'http://localhost:8080/products' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Macbook Pro 2021",
    "price": 2000.00,
}'
```

#### GET /products/{id}
```bash
curl --location 'http://localhost:8080/products/ad10a8bc-c9e7-4d5d-8a4f-7611cdba3c95'
```

#### GET /products/{id}/enable
```bash
curl --location 'http://localhost:8080/products/0b8bd385-f222-4969-99c6-dc31071c0657/enable'
```

#### GET /products/{id}/disable
```bash
curl --location 'http://localhost:8080/products/0b8bd385-f222-4969-99c6-dc31071c0657/disable'
```