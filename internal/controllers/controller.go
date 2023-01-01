package controllers

import (
	"extractor/internal/configs"
	"extractor/internal/db_client"
	"extractor/internal/internal_state"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func Info() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, internal_state.InternalState())
	}
}

func Extract() gin.HandlerFunc {
	return func(c *gin.Context) {
		go internal_state.Scrape()

		c.JSON(http.StatusAccepted, "")
	}
}
