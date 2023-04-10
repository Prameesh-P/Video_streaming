package routes

import (
	"video/controllers"
	"github.com/gin-gonic/gin"
)
func Backend(c *gin.Engine){

	c.GET("/videos",controllers.AllVideos)
	c.GET("/videos/:name",controllers.DisplayVideo)
	c.GET("/songs/:name",controllers.AudioStreaming)
	c.GET("/songs",controllers.AllSongs)
	c.GET("/displays/:name",controllers.StreamVideo)

}