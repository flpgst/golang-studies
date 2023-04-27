package main

import "github.com/flpgst/golang-studies/45-API/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
