package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	cleanup(db)

	product := NewProduct("Product 1", 10.0)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product2 := NewProduct("Product 2", 30.0)
	err = insertProduct(db, product2)
	if err != nil {
		panic(err)
	}
	product2.Name = "Product 2 Updated"
	updateProduct(db, product2)

	p, err := getProduct(db, product2.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %s, %s, %f\n", p.ID, p.Name, p.Price)

	ps, err := getProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range ps {
		fmt.Printf("Product: %s, %s, %f\n", p.ID, p.Name, p.Price)
	}

}

func cleanup(db *sql.DB) {
	_, err := db.Exec("DELETE FROM products")
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func getProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	product := &Product{}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func getProducts(db *sql.DB) ([]*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		product := &Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
