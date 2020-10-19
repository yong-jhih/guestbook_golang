package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Logger(c *gin.Context) {
	fmt.Println("exec Logger")

	//你可以写一些逻辑代码

	// 执行该中间件之后的逻辑
	c.Next()
}