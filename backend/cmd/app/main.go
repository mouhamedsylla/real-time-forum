package main

import (
	"real-time-forum/server/gateway"
)

func main() {
	app := gateway.NewGateway()
	app.BootstrapApp()
}
