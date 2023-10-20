package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/jakeNorman007/api_go/initializers"
)

func init(){
    initializers.LoadEnvVariables()
}

func main(){
    r := gin.Default()

	r.GET("/", func(c *gin.Context) {
	    c.JSON(200, gin.H{ 
            "message": "pong",
        })
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
