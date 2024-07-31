package server

import (
	"github.com/gin-gonic/gin"
	db "github.com/piotroszko/backend-go/database"
	"github.com/piotroszko/backend-go/modules/router"
)

func Init() {

	db.DbConnect()
	db.DbAutoMigrate()

	r := gin.Default()
	router.AddV1Routes(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
