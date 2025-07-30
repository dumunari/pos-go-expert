//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"di/product"

	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.ProductUsecase {
	wire.Build(
		setRepositoryDependency,
		product.NewProductUsecase,
	)
	return &product.ProductUsecase{}
}
