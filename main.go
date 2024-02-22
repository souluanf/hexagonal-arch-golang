package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/souluanf/hexagonal-arch-golang/adapters/db"
	"github.com/souluanf/hexagonal-arch-golang/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite3.db")
	productDbAdapter := db2.NewProductDb(db)

	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product 1", 10)

	productService.Enable(product)

}
