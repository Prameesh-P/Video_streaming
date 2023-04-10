package controllers

import (
	// "encoding/json"
	// "fmt"
	// "fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Videos struct {
	Id     int
	Videos string
}

func AllSongs(c *gin.Context) {
	
	tmpl, err := template.ParseFiles("template/list_songs.html")
	if err != nil {
		panic(err)
	}
	var vidoes Videos
	vidoes.Id=0
	folderAbsPath, err := filepath.Abs("songs")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	files, err := os.ReadDir(folderAbsPath)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not read directory..",
		})
		return
	}

	filesNames := make([]string, 0, len(files))

	for _, file := range files {
		//extracting video names
		baseName := filepath.Base(file.Name())
		Extention := filepath.Ext(file.Name())
		filesNames = append(filesNames, baseName[0:len(baseName)-len(Extention)])
	}
	err = tmpl.Execute(c.Writer, filesNames)
	if err != nil {
		panic(err)
	}

}
