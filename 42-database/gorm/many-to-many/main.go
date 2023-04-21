package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	eletronicosCategory := Category{
		Name: "Eletronicos",
	}
	db.Create(&eletronicosCategory)

	cozinhaCategory := Category{
		Name: "Cozinha",
	}
	db.Create(&cozinhaCategory)

	db.Create(&Product{
		Name:       "Panela Eletrica",
		Price:      399.00,
		Categories: []Category{eletronicosCategory, cozinhaCategory},
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println(" - ", product.Name)
		}
	}

}
