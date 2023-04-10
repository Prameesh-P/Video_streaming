package controllers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func AllVideos(c *gin.Context) {

	tmpl, err := template.ParseFiles("template/list_video.html")
	if err != nil {
		panic(err)
	}
	folderAbsPath, err := filepath.Abs("videos")
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

		baseName := filepath.Base(file.Name())
		Extention := filepath.Ext(file.Name())
		filesNames = append(filesNames, baseName[0:len(baseName)-len(Extention)])

	}
	
	err = tmpl.Execute(c.Writer, filesNames)
		if err != nil {
			panic(err)
		}

}
