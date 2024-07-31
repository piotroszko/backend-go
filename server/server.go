package server

import (
	"github.com/gin-gonic/gin"
	db "github.com/piotroszko/backend-go/database"
	"github.com/piotroszko/backend-go/modules/router"
)

func Init() {
	r := gin.Default()
	db.DbConnect()
	db.DbAutoMigrate()
	router.AddV1Routes(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
