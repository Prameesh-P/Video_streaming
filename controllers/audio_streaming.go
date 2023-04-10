package controllers

import (
	"fmt"

	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)


func AudioStreaming(c *gin.Context){
	params:=c.Param("name")
	AudioName:=fmt.Sprintf("%v",params)


	c.Header("Content-Type", "audio/mpeg")
	c.Header("Transfer-Encoding", "chunked")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	file, err := os.ReadFile("songs/"+AudioName+".mp3")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(file); i += 1024 {
		if i+1024 < len(file) {
			c.Writer.Write(file[i : i+1024])
		} else {
			c.Writer.Write(file[i:])
		}
		c.Writer.Flush()
		time.Sleep( 5* time.Millisecond) 
	}
}
