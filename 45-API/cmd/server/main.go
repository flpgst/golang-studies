package main

import (
	"log"
	"net/http"

	"github.com/flpgst/golang-studies/45-API/configs"
	_ "github.com/flpgst/golang-studies/45-API/docs"
	"github.com/flpgst/golang-studies/45-API/internal/entity"
	"github.com/flpgst/golang-studies/45-API/internal/infra/database"
	"github.com/flpgst/golang-studies/45-API/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title 				GoExpert Example API
// @version 			1.0
// @description 		This is a sample server Petstore server.
// @termsOfService 		http://swagger.io/terms/

// @contact.name 		Filipe Augusto Gonçalves
// @contact.url 		http://www.swagger.io/support
// @contact.email 		flpgst@gmail.com

// @license.name 		NoLicense
// @license.url 		http://noLicense.nolicense

// @host 				localhost:8000
// @BasePath 			/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpiresin))

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

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/docs/doc.json"),
	))

	http.ListenAndServe(":8000", r)
}

// Raw middleware exaple
// r.Use(LogRequest) for activate it
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println()
		next.ServeHTTP(w, r)
	})
}
