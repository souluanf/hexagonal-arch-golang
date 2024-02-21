# Hexagonal Architecture in Go

### Generate mocks
```bash
mockgen -destination=application/mocks/application.go -source=application/product.go application
```

### Run tests with coverage and generate html report
```bash
go test ./... -coverprofile=coverage.out
grep -v "mocks/" coverage.out > filtered_coverage.out
rm coverage.out
go tool cover -html=filtered_coverage.out -o coverage.html
rm filtered_coverage.out
```