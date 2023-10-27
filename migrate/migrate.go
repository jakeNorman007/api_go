package main

import (
    "github.com/JakeNorman007/api_go/initializers"
    "github.com/JakeNorman007/api_go/models"
)

func init(){
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
}

func main(){
    initializers.DB.AutoMigrate(&models.Post{})
}
