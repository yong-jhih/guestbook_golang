package main

import (
	"guestbook_golang/router"
	"guestbook_golang/middleware"
	"guestbook_golang/api"

	"github.com/gin-gonic/gin"
)

func main()  {
	engine := gin.Default()
	engine.LoadHTMLGlob("view/*")

	////////////////////////////////////////////
	// view
	////////////////////////////////////////////

	engine.GET("/",router.Index)

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