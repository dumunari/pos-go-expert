package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           string       `gorm:"primaryKey"`
	Name         string       `gorm:"size:255"`
	Price        float64      `gorm:"type:decimal(10,2)"`
	Categories   []Category   `gorm:"many2many:product_categories;"`
	SerialNumber SerialNumber `gorm:"foreignKey:ProductID;references:ID"`
	gorm.Model
}

type Category struct {
	ID       string    `gorm:"primaryKey"`
	Name     string    `gorm:"size:255"`
	Products []Product `gorm:"many2many:product_categories;"`
}

type SerialNumber struct {
	ID        string `gorm:"primaryKey"`
	Serial    string `gorm:"size:255"`
	ProductID string `gorm:"size:255"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	cleanup(db)
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	createCategories(db)
	createProducts(db)
	getAllProductsWithCategoriesAndSerialNumbers(db)
	getAllCategoriesWithProductsAndSerialNumbers(db)
}

func getAllCategoriesWithProductsAndSerialNumbers(db *gorm.DB) {
	var categories []Category
	db.Preload("Products").Preload("Products.SerialNumber").Find(&categories)
	for _, category := range categories {
		println("Category:", category.Name)
		for _, product := range category.Products {
			println("- Product:", product.Name)
			println("- Serial Number:", product.SerialNumber.Serial)
		}
	}
}

func getAllProductsWithCategoriesAndSerialNumbers(db *gorm.DB) {
	var products []Product
	db.Preload("Categories").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		println("Product:", product.Name)
		for _, category := range product.Categories {
			println("Category:", category.Name)
		}
		println("Serial Number:", product.SerialNumber.Serial)
	}
}

func createCategories(db *gorm.DB) {
	categories := []Category{
		{ID: "1", Name: "Electronics"},
		{ID: "2", Name: "Books"},
		{ID: "3", Name: "Clothing"},
	}

	for _, category := range categories {
		db.Create(&category)
	}
}

func createProducts(db *gorm.DB) {
	products := []Product{
		{ID: "1", Name: "Laptop", Price: 1000.00,
			Categories:   []Category{{ID: "1"}},
			SerialNumber: SerialNumber{ID: "1", Serial: "SN12345"}},
		{ID: "2", Name: "Smartphone", Price: 500.00,
			Categories:   []Category{{ID: "1"}},
			SerialNumber: SerialNumber{ID: "2", Serial: "SN67890"}},
		{ID: "3", Name: "Novel", Price: 20.00,
			Categories:   []Category{{ID: "2"}},
			SerialNumber: SerialNumber{ID: "3", Serial: "SN11223"}},
		{ID: "4", Name: "T-shirt", Price: 15.00,
			Categories:   []Category{{ID: "3"}},
			SerialNumber: SerialNumber{ID: "4", Serial: "SN44556"}},
		{ID: "5", Name: "Jeans", Price: 40.00,
			Categories:   []Category{{ID: "3"}},
			SerialNumber: SerialNumber{ID: "5", Serial: "SN77889"}},
	}

	for _, product := range products {
		db.Create(&product)
	}
}

func cleanup(db *gorm.DB) {
	db.Migrator().DropTable("product_categories")
	db.Migrator().DropTable("categories")
	db.Migrator().DropTable("products")
	db.Migrator().DropTable("serial_numbers")
}
