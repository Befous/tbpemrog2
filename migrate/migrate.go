package main

import (
	"github.com/Befous/tbpemrog2/initializers"
	"github.com/Befous/tbpemrog2/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Product{})
}
