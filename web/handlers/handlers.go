package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	data := gin.H{
		"title": "preved",
		"h1":    "medved",
	}

	c.HTML(http.StatusOK, "home.html", data)
}
