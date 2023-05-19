package main

import (
	"context"
	"database/sql"

	"github.com/flpgst/golang-studies/52-SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Go",
	// 	Description: sql.NullString{String: "Go course", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "5bd914ba-fb01-4987-b608-52b06fe210cb",
		Name:        "updated",
		Description: sql.NullString{String: "updated", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	err = queries.DeleteCategory(ctx, "5bd914ba-fb01-4987-b608-52b06fe210cb")
	if err != nil {
		panic(err)
	}
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
