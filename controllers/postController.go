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

    // Bind methid in Gin serailizes the body struct into JSON in this case
    c.Bind(&body)

    //create a post
    post := models.Post{Title: body.Title, Body: body.Body}

    result := initializers.DB.Create(&post) // pass pointer of data to Create

    if result.Error != nil {
        c.Status(400)
        return
    }

    //returns post
    c.JSON(200, gin.H{
        "post": post,
    })
}

func PostsIndex(c *gin.Context){
    // get all of the posts
    var posts []models.Post
    initializers.DB.Find(&posts)

    // respond from the request with all of the posts
    c.JSON(200, gin.H{
        "posts": posts,
    })
}

func PostShow(c *gin.Context){
    // get the id of the post off the url, which is linked in main.go i.e. "/post/:id"
    id := c.Param("id")

    // get all of the posts
    var post models.Post
    initializers.DB.First(&post, id )

    // respond from the request with a single post, based of the id
    c.JSON(200, gin.H{
        "post": post,
    })
}

func PostUpdate(c *gin.Context){
    // get the id off of the url
    id := c.Param("id")

    // get the data of the request body
    var body struct {
        Body    string
        Title   string
    }

    // Bind methid in Gin serailizes the body struct into JSON in this case
    c.Bind(&body)

    // find the post we are wanting to update
    var post models.Post
    initializers.DB.First(&post, id )

    // update the post
    initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body,})

    // response
    c.JSON(200, gin.H{
        "post": post,
    })
}

func PostDelete(c *gin.Context){
    // get id of post you want to delete off of the url
    id := c.Param("id")

    // delete the post
    initializers.DB.Delete(&models.Post{}, id)

    // response
    c.Status(200)
}
