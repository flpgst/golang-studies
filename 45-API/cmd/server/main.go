package main

import (
	"net/http"

	"github.com/flpgst/golang-studies/45-API/configs"
	"github.com/flpgst/golang-studies/45-API/internal/entity"
	"github.com/flpgst/golang-studies/45-API/internal/entity/infra/database"
	"github.com/flpgst/golang-studies/45-API/internal/entity/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExpiresin)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8000", r)
}