package main

import (
	"guestbook_golang/api"
	"guestbook_golang/middleware"
	"guestbook_golang/router"
	// "net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	////////////////////////////////////////////
	// view
	////////////////////////////////////////////

	engine.LoadHTMLGlob("view/*")
	engine.Static("/static", "./static")
	loginStatus := engine.Group("/", middleware.Middleware1)
	{
		loginStatus.GET("", router.GetIndex)
		loginStatus.GET("register", router.GetRegister)
		loginStatus.GET("member", router.GetMember)
	}

	////////////////////////////////////////////
	// api
	////////////////////////////////////////////

	engine.POST("/form_post", api.Form_post)

	engine.Run()
}
