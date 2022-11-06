package main

import (
	"go-todolist/api" // path package module

	"go-todolist/database" // path package module
	"net/http"

	"github.com/gin-contrib/cors" // connect cross platform (angular, react)
	"github.com/gin-gonic/gin"    // make API
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	database.ConnectDatabase()

	r.GET("/todo-lists", api.FindAll)
	r.GET("/todo-list", api.FindByUsername)
	r.POST("/new-todo-list", api.CreateTodoList)
	r.DELETE("delete/:id", api.DeleteList)
	r.POST("/upload", api.Upload)
	r.StaticFS("/file", http.Dir("public"))

	r.Run()
}
