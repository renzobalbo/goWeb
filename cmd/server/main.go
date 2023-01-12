package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/renzobalbo/goWeb/cmd/server/handler"
	"github.com/renzobalbo/goWeb/internal/product"
	"github.com/renzobalbo/goWeb/pkg/store"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}

	db := store.NewStorage(os.Getenv("DB_FILE_NAME"))
	repo := product.NewRepository(db)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAll())
		products.GET(":id", productHandler.GetByID())
		products.GET("/search", productHandler.Search())
		products.POST("", productHandler.Post())
		products.PUT("/:id", productHandler.Update())
		products.PATCH("/:id", productHandler.UpdatePrice())
		products.DELETE("/:id", productHandler.Delete())
	}
	r.Run(":8080")
}

// loadProducts carga los productos desde un archivo json
// func loadProducts(path string, list *[]domain.Product) {
// 	file, err := os.ReadFile(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = json.Unmarshal([]byte(file), &list)
// 	if err != nil {
// 		panic(err)
// 	}
// }
