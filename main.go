package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nirav18/gorestapi/services"
	"github.com/nirav18/gorestapi/utils"
	// "github.com/rs/zerolog"
)

func main() {
    // gin.SetMode(gin.ReleaseMode)
    r := gin.New() //new engine, you can use gin.Default for default use
    r.Use(utils.DefaultStructuredLogger())
    r.Use(gin.Recovery())
    r.GET("/", services.PostLinkHandler) 
    r.Run()
}
