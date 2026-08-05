package main

import (
	"log"
	"os"
	"time"

	"portfolio-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal("Erro ao configura proxies: ", err)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			// "http://localhost:3000",
			// "http://localhost:5000",
			// "http://localhost:8080",
			"https://thedevjeffportfolio.web.app",
		},
		AllowMethods: []string{
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
		},
		MaxAge: 12 * time.Hour,
	}))

	routes.SetupRoutes(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	err = router.Run(":" + port)

	if err != nil {
		log.Fatal("Erro ao iniciar servidor: ", err)
	}
}
