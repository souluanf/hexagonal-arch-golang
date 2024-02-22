# Hexagonal Architecture in Go

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