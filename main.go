package main 

import (
	"github.com/gin-gonic/gin"
	"video/routes"
)

func main(){
router:=gin.Default()
routes.Backend(router)
router.Run(":8008")
}




