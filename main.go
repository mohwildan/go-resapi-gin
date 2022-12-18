package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mohwildan/go-resapi-gin/controllers/productController"
	"github.com/mohwildan/go-resapi-gin/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/product", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product/:id", productcontroller.Delete)

	r.Run(":3000")
}
