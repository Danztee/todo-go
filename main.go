package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Buy groceries", Completed: false},
	{ID: "2", Item: "Clean the house", Completed: false},
	{ID: "3", Item: "Pay bills", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")

	fmt.Println(id, "id")

	todo, err := getTodosById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Todo not found",
		})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleStatus(context *gin.Context) {
	id := context.Param("id")

	fmt.Println(id, "id")

	todo, err := getTodosById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Todo not found",
		})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodosById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
