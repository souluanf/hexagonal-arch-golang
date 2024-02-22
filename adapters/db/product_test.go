package db_test

import (
	"database/sql"
	"github.com/souluanf/hexagonal-arch-golang/adapters/db"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	createProductTable := `CREATE TABLE products (
		id VARCHAR(36) PRIMARY KEY,
		name VARCHAR(255),
		price FLOAT,
		status VARCHAR(255)
	);`
	stmt, err := db.Prepare(createProductTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, _ = stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products (id, name, price, status) VALUES ("abc", "Product test", 0, "disabled")`
	stmt, err := Db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setup()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())

}
