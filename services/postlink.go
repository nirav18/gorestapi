package services

import (
	"github.com/gin-gonic/gin"
)

func PostLinkHandler(c *gin.Context) {
    c.JSON(200, "Downloading MP3")
} 
