package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haru210/go-TODO/internal/database"
	"github.com/haru210/go-TODO/internal/models"
)

func GetTodos(c *gin.Context) {
	var todos []models.ToDo
	database.DB.Find(&todos)
	c.HTML(http.StatusOK, "index.html", gin.H{"Todos": todos})
}

func CreateTodo(c *gin.Context) {
	var form struct {
		Title     string `form:"title"`
		Completed bool   `form:"completed"`
	}
	if err := c.Bind(&form); err == nil && form.Title != "" {
		todo := models.ToDo{
			Title:     form.Title,
			Completed: form.Completed,
		}
		database.DB.Create(&todo)
	}

	c.Redirect(http.StatusFound, "/")
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.ToDo

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.String(http.StatusNotFound, "Todo not found")
		return
	}

	method := c.PostForm("method")
	if method == "PUT" {
		todo.Completed = c.PostForm("completed") == "true"
		database.DB.Save(&todo)
	}

	c.Redirect(http.StatusFound, "/")
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.ToDo

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.String(http.StatusNotFound, "Todo not found")
		return
	}

	if c.PostForm("_method") == "DELETE" {
		database.DB.Delete(&todo)
	}

	c.Redirect(http.StatusFound, "/")
}
