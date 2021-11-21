package main

import (
	"fmt"
	"os"

	"github.com/Poche-1/entrust_test/server_"
)

func main() {
	server := server_.NewServer(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	server.Handle("GET", "/", server_.HandleRoot)
	server.Handle("POST", "/cars", server.AddMiddleware(server_.CarPostRequest, server_.CheckAuth(), server_.CheckBodyCar(), server_.Loggin()))
	server.Handle("GET", "/cars/:id", server.AddMiddleware(server_.CarGetRequest, server_.CheckAuth(), server_.Loggin()))
	server.Handle("DELETE", "/cars/:id", server.AddMiddleware(server_.CarDeleteRequest, server_.CheckAuth(), server_.Loggin()))
	server.Listen()
}
