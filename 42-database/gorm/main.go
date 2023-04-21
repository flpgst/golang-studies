package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// var category Category
	// db.Find(&category, 1)
	// fmt.Println(category)

	// db.Create(&Product{
	// 	Name:     "Keyboard",
	// 	Price:    100.00,
	// 	Category: category,
	// })

	// db.Create(&SerialNumber{
	// 	Number:    "12345",
	// 	ProductID: 2,
	// })

	// db.Create(&Category{
	// 	Name: "Eletronicos",
	// })

	// products := []Product{
	// 	{Name: "Notebook", Price: 1000.00},
	// 	{Name: "Mouse", Price: 10.00},
	// 	{Name: "Keyboard", Price: 100.00},
	// }

	// db.Create(&products)

	// var product Product
	// db.First(&product, 1)
	// fmt.Println(product)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Where("price > ?", 100).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Where("name like ?", "%book%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Product"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2.Name)

	// db.Delete(&p2)

	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product.ID, product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println(" - ", product.Name, product.SerialNumber.Number)
		}
	}

}
