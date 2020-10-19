package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Middleware1(c *gin.Context) {
	fmt.Println("exec middleware1")
	//你可以写一些逻辑代码
	// 执行该中间件之后的逻辑
	c.Next()
}

func GetLoginStatus() gin.HandlerFunc {
	fmt.Println("exec GetLoginStatus")
	return func(c *gin.Context) {
		cookie, _ := c.Request.Cookie("account")
		if cookie == nil {
			
			return
			
		} else {

		}
		c.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Request.Cookie("account")
		if cookie == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "請先登陸"})
			c.Abort()
		}
		c.Next()
	}
}
