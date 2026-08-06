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
			"https://thedevjeffportfolio.web.app",
			"http://localhost:3000",
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

	log.Printf("Servidor iniciado na porta %s", port)

	err = router.Run("0.0.0.0:" + port)

	if err != nil {
		log.Fatal("Erro ao iniciar servidor: ", err)
	}
}
