package main

import (
    "github.com/gin-gonic/gin"
    "github.com/JakeNorman007/api_go/initializers"
    "github.com/JakeNorman007/api_go/controllers"
)

func init(){
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
}

func main(){
    r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
    r.GET("/posts/:id", controllers.PostShow)
    r.PUT("/posts/:id", controllers.PostUpdate)
    r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run()
}
