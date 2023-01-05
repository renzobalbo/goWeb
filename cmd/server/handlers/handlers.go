package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renzobalbo/goWeb/global"
)

// Metodos GET

func Pong(ctx *gin.Context) {
	ctx.String(200, "pong")
}

func Products(ctx *gin.Context) {
	productos := global.Productos
	ctx.JSON(200, gin.H{
		"products": productos,
	})
}

func ProductId(ctx *gin.Context) {
	tempId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "failed to parse id",
			"data":    nil,
		})
		return
	}
	actualProduct := global.Producto{}
	for i, a := range global.Productos {
		if a.Id == tempId {
			actualProduct = global.Productos[i]
			break
		}
	}

	if actualProduct.Id != 0 {
		ctx.JSON(200, gin.H{
			"message": "Found it!",
			"data":    actualProduct,
		})
		return
	} else {
		ctx.JSON(404, gin.H{
			"message": "Id not found!",
			"data":    nil,
		})
		return
	}
}

func ProductsPriceGt(ctx *gin.Context) {
	priceQuery, err := strconv.ParseFloat(ctx.Query("price"), 64)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "Couldn't find a match",
			"data":    nil,
		})
		return
	}

	var filteredProducts = make([]global.Producto, 0)
	for _, w := range global.Productos {
		if priceQuery != 0 && w.Price > float64(priceQuery) {
			filteredProducts = append(filteredProducts, w)
		}
	}
	ctx.JSON(200, gin.H{
		"message": "Products filtered:",
		"data":    filteredProducts,
	})
}

// Metodos POST

type request struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func AddProduct(ctx *gin.Context) {
	var req request
	err := ctx.ShouldBind(&req)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	global.LastId++
	newProd := global.Producto{Id: global.LastId, Name: req.Name, Quantity: req.Quantity, CodeValue: req.CodeValue, IsPublished: req.IsPublished, Expiration: req.Expiration, Price: req.Price}
	global.Productos = append(global.Productos, newProd)
	ctx.JSON(201, req)
}
