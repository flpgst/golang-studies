//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/flpgst/golang-studies/54-DI/product"
	"github.com/google/wire"
)

var setRepositoryDependencies = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(dc *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependencies,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
