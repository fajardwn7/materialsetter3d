package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStuffs ... Get all stuff
func GetStuffs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "It's Work",
	})
}
