package controllers

import (
	"log"
	"net/http"

	"github.com/multimoml/extractor/internal/configs"
	"github.com/multimoml/extractor/internal/db_client"
	"github.com/multimoml/extractor/internal/internal_state"

	"github.com/gin-gonic/gin"
)

// Self is a simple endpoint to check if the server is alive and working
// @Summary Get liveness and readiness status of the microservice
// @Description Get liveness ad readiness status of the microservice
// @Tags Kubernetes
// @Success 200 {} nil
// @Failure 503 {} string
// @Router /self [get]
func Self() gin.HandlerFunc {
	return func(c *gin.Context) {
		dbClient := db_client.DBClient(c)

		if err := dbClient.Ping(c, nil); err != nil {
			c.JSON(http.StatusServiceUnavailable, "")
			log.Println(err)
		} else if val, err := configs.GetConfig("broken"); err == nil && val == "1" {
			c.String(http.StatusServiceUnavailable, "dead by config")
			return
		} else {
			c.JSON(http.StatusOK, "")
		}
	}
}

// info returns a extractors internal state
// @Summary Get internal state
// @Description Get extractors internal state
// @Tags Info
// @Produce json
// @Success 200 {} object
// @Router /info [get]
func Info() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, internal_state.InternalState())
	}
}

// Extract signals internal state to initiate a new extraction
// @Summary Initiate extraction
// @Description Initiate extraction
// @Tags Extract
// @Success 202 {} nil
// @Router /extract [post]
func Extract() gin.HandlerFunc {
	return func(c *gin.Context) {
		go internal_state.Scrape()

		c.JSON(http.StatusAccepted, "")
	}
}
