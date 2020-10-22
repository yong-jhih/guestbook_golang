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

	engine.LoadHTMLGlob("view/*") // 設定模板路徑
	//engine.LoadHTMLFiles("view/Index.html", "view/Login.html") // 設定模板檔案

	engine.Static("/static", "./static") // 設定靜態資源路徑
	
	getLoginStatus := engine.Group("/", middleware.GetLoginStatus())
	{
		getLoginStatus.GET("", router.GetIndex)
		getLoginStatus.GET("register", router.GetRegister)
		getLoginStatus.GET("member", router.GetMember)
	}

	////////////////////////////////////////////
	// api
	////////////////////////////////////////////

	engine.POST("/api/form_post", api.FormPost)
	engine.POST("/api/check_account", api.CheckAccount)
	engine.POST("/api/add_account", api.AddAccount)

	engine.Run()
}
