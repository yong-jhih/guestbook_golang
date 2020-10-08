package main

import (
	"guestbook_golang/router"
	"guestbook_golang/middleware"
	"guestbook_golang/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	engine := gin.Default()
	engine.LoadHTMLGlob("view/*")

	////////////////////////////////////////////
	// view
	////////////////////////////////////////////
	var a = []int{1,2,3,4,5,6}

	engine.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "Index.html", gin.H{
			"title": "Main website",
			"content":"Hello!",
			"array":a,
        })
    })

	v1 := engine.Group("/v1", middleware.Middleware1)
	{
		v1.GET("/register",router.Register)
	}

	v2 := engine.Group("/v2", middleware.Middleware2)
	{
		v2.GET("/login",router.Login)
	}

	
	////////////////////////////////////////////
	// api
	////////////////////////////////////////////

	engine.POST("/form_post", api.Form_post)

	engine.Run()
}