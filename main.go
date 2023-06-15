package main

import (
	"github.com/Befous/tbpemrog2/controllers"
	"github.com/Befous/tbpemrog2/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./templates/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.tmpl", map[string]string{"title": ""})
	})

	r.PUT("/product/:id", controllers.UpdateProduct)
	r.GET("/product", controllers.GetProduct)
	r.GET("/product/:id", controllers.GetProductID)
	r.POST("/product", controllers.CreateProduct)
	r.DELETE("/product/:id", controllers.DeleteProduct)

	r.Run() // listen and serve on 0.0.0.0:8080
}
