package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jgreat/things/db"
	"github.com/jgreat/things/models"
	"github.com/jgreat/things/routes"
)

func main() {
	router := gin.Default()

	db, err := db.Connect()
	if err != nil {
		log.Fatalf("Database Error: %v", err)
	}

	err = models.CreateTable(db, "thing", &models.Thing{})
	if err != nil {
		log.Fatalf("Failed to Create Table: %v", err)
	}

	// Add DB to context for routes/controllers
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Routes - list of routes
	router.GET("/", routes.GetWelcome)
	router.GET("/things", routes.GetThings)
	router.POST("/thing", routes.CreateThing)

	// listen and serve on 0.0.0.0:8080
	log.Fatal(router.Run())
}
