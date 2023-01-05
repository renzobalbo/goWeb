package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/renzobalbo/goWeb/cmd/server/handlers"
	"github.com/renzobalbo/goWeb/global"
)

func LoadJson() {
	data, err := ioutil.ReadFile("./products.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, &global.Productos)
	if err != nil {
		fmt.Println(err)
		return
	}

	global.LastId = len(global.Productos)
}

func main() {
	LoadJson()

	router := gin.Default()

	router.GET("/ping", handlers.Pong)
	router.GET("/products", handlers.Products)
	router.GET("/products/:id", handlers.ProductId)
	router.GET("/products/search", handlers.ProductsPriceGt)

	//Metodo Post
	router.POST("/products", handlers.AddProduct)

	router.Run()
}
