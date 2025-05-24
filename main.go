package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haru210/go-TODO/internal/database"
	"github.com/haru210/go-TODO/internal/handlers"
)

func main() {
	database.Init()

	r := gin.Default()

	r.LoadHTMLGlob("internal/templates/*.html")
	r.Static("/static", "/internal/static")

	r.POST("/add", handlers.CreateTodo)
	r.GET("/", handlers.GetTodos)
	r.POST("/update/:id", handlers.UpdateTodo)
	r.POST("/delete/:id", handlers.DeleteTodo)
	r.Run(":8080")
}
