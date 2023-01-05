package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/renzobalbo/goWeb/dia1/E2.2/global"
	"github.com/renzobalbo/goWeb/dia1/E2.2/handlers"
)

func LoadJson() {
	data, err := ioutil.ReadFile("./dia1/products.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, &global.Productos)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	LoadJson()

	router := gin.Default()

	router.GET("/ping", handlers.Pong)
	router.GET("/products", handlers.Products)
	router.GET("/products/:id", handlers.ProductId)
	router.GET("/products/search", handlers.ProductsPriceGt)

	router.Run()
}
