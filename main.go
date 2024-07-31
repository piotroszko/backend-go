package main

import (
	"github.com/piotroszko/backend-go/helpers/env"
	"github.com/piotroszko/backend-go/server"
)

func main() {
	env.LoadEnv()
	server.Init()

}
