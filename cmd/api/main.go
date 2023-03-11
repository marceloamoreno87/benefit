package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/marceloamoreno87/benefit/api/docs"
	"github.com/marceloamoreno87/benefit/api/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Konsi API - Benefit
// @version         1.0
// @description     Sistema para capturar beneficios de usuários usando o CPF.

// @contact.name   Marcelo Moreno
// @contact.email  marceloamoreno87@gmail.com

// @host      localhost:8080
// @BasePath  /api
func main() {
	godotenv.Load()

	r := routes.SetupRouter()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(os.Getenv("PORT"))
}
