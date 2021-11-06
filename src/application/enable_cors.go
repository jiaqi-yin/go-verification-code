package application

import "github.com/gin-contrib/cors"

func enableCors() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}
	config.AddAllowMethods("OPTIONS")
	router.Use(cors.New(config))
}
