package main

import (
	"encoding/json"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func LoadData() []Product {
	file, err := os.Open("./dia1/products.json")

	if err != nil {
		panic(err)
	}

	byteJson, _ := io.ReadAll(file)

	defer file.Close()

	products := []Product{}

	json.Unmarshal(byteJson, &products)

	return products
}

func exmain() {

	catalog := LoadData()

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "Pong")
	})

	router.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, catalog)
	})

	router.GET("/products/:id", func(ctx *gin.Context) {
		prId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, "That product does not exist!")
			return
		}
		var actualPr Product
		for _, a := range catalog {
			if a.ID == prId {
				actualPr = a
				break
			}
		}

		ctx.JSON(200, actualPr)

	})

	router.GET("/products/search", func(ctx *gin.Context) {

	})

	router.Run()

}
