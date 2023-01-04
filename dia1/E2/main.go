package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre   string `json:"nombre" form:"nombre"`
	Apellido string `json:"apellido" form:"apellido"`
}

func main() {

	router := gin.Default()

	router.POST("/saludar", func(ctx *gin.Context) {
		var json Persona
		ctx.Bind(&json)
		var saludo string = "Hola " + json.Nombre + " " + json.Apellido
		ctx.String(http.StatusCreated, saludo)

	})

	router.Run()

}
