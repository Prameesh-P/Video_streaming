package controllers

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func StreamVideo(c *gin.Context) {
	//taking video name as parameter
	videoFiles:=c.Param("name")
    
	videoPath :="videos/"+videoFiles+".mp4"

	file, err := os.Open(videoPath)
	
	if err != nil {
		c.AbortWithError(500, err)
		return
	
	}
	defer file.Close()


    c.Header("Content-Type", "video/mp4")

    // Set the transfer encoding to "chunked"
    c.Header("Transfer-Encoding", "chunked")

    // Read the file in 1 MB chunks and write them to the response writer
    buffer := make([]byte, 1024*1024) // 1 MB buffer
    for {
        n, err := file.Read(buffer)
        if err != nil {
            if err == io.EOF {
                break
            }
            c.AbortWithError(http.StatusInternalServerError, err)
            return
        }

        // Write the chunk to the response writer
        if _, err := c.Writer.Write(buffer[:n]); err != nil {
            return
        }
        // Flush the response writer to ensure the data is sent immediately
        c.Writer.Flush()
    }


}
func DisplayVideo(c *gin.Context){
	//taking video name as parameter
	param:=c.Param("name")

	//parsing html template
	tpl,err:=template.ParseFiles("template/video.html")

	if err !=nil{
		c.AbortWithError(500, err)
		return 
	}
	//sending params into html and render that 
	err = tpl.Execute(c.Writer,param)
	if err !=nil{
		c.AbortWithError(500, err)
		return 
	}
}