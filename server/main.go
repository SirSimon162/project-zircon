package main

import (
	"os"

	"github.com/SirSimon162/project-zircon/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id", routes.EntryById)
	router.GET("/deadline/:deadline", routes.GetEntryByDeadline)
	router.GET("/assignedTo/:id", routes.GetEntryByAssignedTo)

	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("/deadline/update/:id", routes.UpdateDeadline)

	router.DELETE("/entry/delete/:id", routes.DeleteEntry)
	router.Run(":" + port)
}
