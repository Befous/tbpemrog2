package controllers

import (
	"github.com/Befous/tbpemrog2/initializers"
	"github.com/Befous/tbpemrog2/models"
	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	//get the post
	var products []models.Product
	initializers.DB.Find(&products)

	//response to them
	c.JSON(200, gin.H{
		"products": products,
	})
}

func GetProductID(c *gin.Context) {
	//get id url
	id := c.Param("id")

	//get the post
	var product models.Product
	initializers.DB.First(&product, id)

	//response to them
	c.JSON(200, gin.H{
		"products": product,
	})
}

func CreateProduct(c *gin.Context) {
	//get data off req bod
	var body struct {
		Name        string
		Category    string
		Stock       int
		Description string
	}

	c.Bind(&body)

	//create a post
	product := models.Product{
		Name:        body.Name,
		Category:    body.Category,
		Stock:       body.Stock,
		Description: body.Description,
	}

	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it

	c.JSON(200, gin.H{
		"products": product,
	})
}

func UpdateProduct(c *gin.Context) {
	//get id url
	id := c.Param("id")

	//get data of request body
	var body struct {
		Name        string
		Category    string
		Stock       int
		Description string
	}

	c.Bind(&body)

	//find post we are updating
	var product models.Product
	initializers.DB.First(&product, id)

	//update it
	initializers.DB.Model(&product).Updates(models.Product{
		Name:        body.Name,
		Category:    body.Category,
		Stock:       body.Stock,
		Description: body.Description,
	})

	//respon with it
	c.JSON(200, gin.H{
		"products": product,
	})
}

func DeleteProduct(c *gin.Context) {
	//get id url
	id := c.Param("id")

	//delete the post
	initializers.DB.Delete(&models.Product{}, id)

	//response to them
	c.Status(200)
}
