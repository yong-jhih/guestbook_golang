package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
	// "os"
)

////////////////////////////////
// view
////////////////////////////////

func GetIndex(ctx *gin.Context) {
	// fmt.Println(os.Args[0])
	ctx.HTML(http.StatusOK, "Index.html", nil)
}

func GetRegister(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "Register.html", nil)
}

func GetMember(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "Member.html", nil)
}

////////////////////////////////
// api
////////////////////////////////

func Index(c *gin.Context) {
	// 使用 c.Param(key) 获取 url 参数
	name := c.Param("name")
	c.String(http.StatusOK, "index %s", name)
}

func Register(c *gin.Context) {
	// 使用 c.Param(key) 获取 url 参数
	name := c.Param("name")
	c.String(http.StatusOK, "register %s", name)
}

func Login(c *gin.Context) {
	// 使用 c.Param(key) 获取 url 参数
	firstname := c.DefaultQuery("firstname", "123")
	lastname := c.Query("lastname")
	// name := c.Param("name")
	c.String(http.StatusOK, "login , %s %s", firstname, lastname)
}
