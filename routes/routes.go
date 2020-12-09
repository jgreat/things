package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jgreat/things/models"

	"github.com/go-pg/pg/v10"
)

// GetWelcome - default get route
func GetWelcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "world",
	})
	return
}

// GetThings - Get all the things
func GetThings(c *gin.Context) {
	db := c.MustGet("db").(*pg.DB)
	things := &[]models.Thing{}

	err := db.Model(things).Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": things})
	return
}

/*
curl --header "Content-Type: application/json" \
  --request POST --data '{"name":"plumbus"}' \
  http://localhost:8080/thing
*/

// CreateThing - add a thing to the db
func CreateThing(c *gin.Context) {
	db := c.MustGet("db").(*pg.DB)
	thing := &models.Thing{}

	// Validate input
	err := c.ShouldBindJSON(thing)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Thing
	thing = &models.Thing{
		Name: thing.Name,
	}

	// probably needs a lot more to check for unique and other constraints...
	_, err = db.Model(thing).Insert()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": thing})
	return
}
