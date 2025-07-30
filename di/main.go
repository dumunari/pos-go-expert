package main

import (
	"database/sql"
	"di/product"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

func main() {
	db, err := sql.Open("sqlite3", "./di.db")
	if err != nil {
		log.Fatal(err)
	}

	repo := product.NewProductRepository(db)
	usecase := product.NewProductUsecase(repo)

	product, err := usecase.GetProduct(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Product ID: %+v", product.ID)
	log.Printf("Product Name: %+v", product.Name)
}
