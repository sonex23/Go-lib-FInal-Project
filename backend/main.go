package main

import (
	"golib/handler"
	"golib/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(handler.CorsMiddleware())

	routes.UserRoute(r)
	routes.CategoryRoute(r)
	routes.BooksRoute(r)
	r.Run(":4444")
}
