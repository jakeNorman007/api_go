package controllers

import(
    "github.com/gin-gonic/gin"
    "github.com/JakeNorman007/api_go/initializers"
    "github.com/JakeNorman007/api_go/models"
)

func PostsCreate(c *gin.Context) {
    // get data off of request body
    var body struct {
        Body    string
        Title   string
    }

    c.Bind(&body)

    //create a post
    post := models.Post{Title: body.Title, Body: body.Body}

    result := initializers.DB.Create(&post) // pass pointer of data to Create

    if result.Error != nil {
        c.Status(400)
        return
    }

    //return post

    c.JSON(200, gin.H{
        "post": post,
    })
}
