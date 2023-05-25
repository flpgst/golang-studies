package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/flpgst/golang-studies/55-CleanArch/configs"
	"github.com/flpgst/golang-studies/55-CleanArch/internal/event/handler"
	"github.com/flpgst/golang-studies/55-CleanArch/internal/infra/graph"
	"github.com/flpgst/golang-studies/55-CleanArch/internal/infra/web/webserver"
	"github.com/flpgst/golang-studies/55-CleanArch/pkg/events"
	"github.com/streadway/amqp"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel(configs)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create)
	fmt.Println("Listening on port", configs.WebServerPort)
	go webserver.Start()

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", configs.GraphQLServerPort)
	log.Fatal(http.ListenAndServe(":"+configs.GraphQLServerPort, nil))

}

func getRabbitMQChannel(configs *configs.Conf) *amqp.Channel {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", configs.RABBITMQ_USER, configs.RABBITMQ_PASSWORD, configs.RABBITMQ_HOST, configs.RABBITMQ_PORT))
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
