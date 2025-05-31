package database

import (
	"apis/internal/entity"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	return db
}

func TestNewProduct(t *testing.T) {
	db := initDB(t)

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Produto Test", 100.0)
	productDB := NewProduct(db)

	err := productDB.Create(product)
	assert.Nil(t, err)

	var foundProduct entity.Product
	err = db.First(&foundProduct, "id = ?", product.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, product.ID, foundProduct.ID)
	assert.Equal(t, product.Name, foundProduct.Name)
	assert.Equal(t, product.Price, foundProduct.Price)
}

func TestFindAll(t *testing.T) {
	db := initDB(t)

	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)

	var products []entity.Product
	for i := 1; i < 24; i++ {
		product, _ := entity.NewProduct(fmt.Sprintf("Produto Test [%d]", i), rand.Float64()*10)
		productDB.Create(product)
		products = append(products, *product)
	}

	foundProducts, _ := productDB.FindAll(1, 10, "asc")
	assert.Equal(t, products[0].Name, foundProducts[0].Name)
	assert.Equal(t, products[9].Name, foundProducts[9].Name)

	foundProducts, _ = productDB.FindAll(2, 10, "asc")
	assert.Equal(t, products[10].Name, foundProducts[0].Name)
	assert.Equal(t, products[19].Name, foundProducts[9].Name)

	foundProducts, _ = productDB.FindAll(3, 10, "asc")
	assert.Equal(t, products[20].Name, foundProducts[0].Name)
	assert.Equal(t, products[22].Name, foundProducts[2].Name)
}

func TestFindProductById(t *testing.T) {
	db := initDB(t)

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Produto Test", 100.0)
	productDB := NewProduct(db)

	err := productDB.Create(product)
	assert.Nil(t, err)

	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Produto Test", product.Name)
}

func TestUpdateProduct(t *testing.T) {
	db := initDB(t)

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Produto Test", 100.0)
	productDB := NewProduct(db)

	err := productDB.Create(product)
	assert.Nil(t, err)

	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Produto Test", product.Name)

	product.Name = "Produto Test 2"

	err = productDB.Update(product)
	assert.NoError(t, err)

	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Produto Test 2", product.Name)
}

func TestDeleProduct(t *testing.T) {
	db := initDB(t)

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Produto Test", 100.0)
	productDB := NewProduct(db)

	err := productDB.Create(product)
	assert.Nil(t, err)

	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Produto Test", product.Name)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
	assert.Equal(t, "record not found", err.Error())
}
