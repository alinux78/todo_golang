package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alinux78/todo/pkg/model"
	"github.com/docker/distribution/uuid"
	"github.com/gin-gonic/gin"
)

var items = []model.TodoItem{
	{ID: "1", Title: "list items endpoint", Description: "list all todos", CreatedAt: time.Now().UTC()},
	{ID: "2", Title: "create item endpoint", Description: "endpoint for creating an item", CreatedAt: time.Now().UTC()},
	{ID: "3", Title: "delete item endpoint", Description: "endpoint for deleting an item", CreatedAt: time.Now().UTC()},
	{ID: "4", Title: "mark item done endpoint", CreatedAt: time.Now().Add(-24 * time.Hour)},
}

const port = 8080

func main() {
	fmt.Println("starting server")

	router := gin.Default()

	//DISCUSS functions
	/*
		In Go you can assign functions to variables, pass functions to functions,
		and even write functions that return functions.
		Functions are first-class—they work in all the places that integers, strings,
		and other types work
	*/
	router.GET("/items", logRequest, getItems)
	router.POST("/items", logRequest, createItem)

	listenAddress := fmt.Sprintf("localhost:%v", port)
	router.Run(listenAddress)
}

func logRequest(c *gin.Context) {
	fmt.Printf("new request %v:%v\n", c.Request.Method, c.Request.URL)
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func createItem(c *gin.Context) {
	var input model.TodoItemInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Printf("cannot add item %v", err)
		return
	}
	newItem := model.TodoItem{
		Title:       input.Title,
		Description: input.Description,
		ID:          uuid.Generate().String(),
		CreatedAt:   time.Now(),
	}
	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}
