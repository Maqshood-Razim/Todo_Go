package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var Todos = []Todo{
	{Id: "1", Item: "clean room", Completed: false},
	{Id: "2", Item: "read book", Completed: false},
	{Id: "3", Item: "record video", Completed: false},
}

func getTodos(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{
		"Todos": Todos,
	})
}

func addTodo(context *gin.Context) {
	item := context.PostForm("item")
	newTodo := Todo{
		Id:        generateID(), 
		Item:      item,
		Completed: false,
	}
	Todos = append(Todos, newTodo)
	context.Redirect(http.StatusSeeOther, "/")
}

func toggleTodoCompletion(context *gin.Context) {
	id := context.Param("id")
	for i, todo := range Todos {
		if todo.Id == id {
			Todos[i].Completed = !Todos[i].Completed
			break
		}
	}
	context.Redirect(http.StatusSeeOther, "/")
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	for i, todo := range Todos {
		if todo.Id == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			break
		}
	}
	context.Redirect(http.StatusSeeOther, "/")
}

func generateID() string {
	return fmt.Sprintf("%d", len(Todos)+1) 
}

func main() {
	router := gin.Default()

	
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", getTodos)
	router.POST("/Todos", addTodo)
	router.POST("/Todos/:id/toggle", toggleTodoCompletion)
	router.POST("/Todos/:id/delete", deleteTodo)

	router.Run("localhost:9090")
}
