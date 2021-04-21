package main

import (
	"github.com/meateam/go-service-template/server"
)

func main() {
	server.NewServer(nil).Serve(nil)
}
