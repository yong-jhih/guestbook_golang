package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Middleware1(c *gin.Context) {
	fmt.Println("exec middleware1")

	//你可以写一些逻辑代码

	// 执行该中间件之后的逻辑
	c.Next()
}
