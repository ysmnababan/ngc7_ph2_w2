package main

import (
	"ngc/config"
	"ngc/handler"
	"ngc/middleware"
	"ngc/repo"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Recovery())

	db := config.Connect()
	defer db.Close()

	Repo := &repo.MysqlRepo{DB: db}
	h := &handler.ProductHandler{Repo: Repo}
	userhandler := handler.UserHandler{Repo: Repo}

	r.POST("/users/register", userhandler.Register)
	r.POST("/users/login", userhandler.Login)

	product := r.Group("/")
	product.Use(middleware.Auth())
	{
		product.GET("/products", h.GetProducts)
		product.GET("/product/:id", h.GetProductById)
		product.POST("/product", h.CreateProduct)
		product.PUT("/product/:id", h.UpdateProduct)
		product.DELETE("/product/:id", h.DeleteProduct)
	}

	r.Run(":8080")
}
