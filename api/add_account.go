package api

import (
	"fmt"
	// "database/sql"
	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-gonic/gin"
)

var MysqlDB *gorm.DB

func AddAccount(c *gin.Context) {
	memberAC := c.PostForm("memberAC")
    memberPW := c.PostForm("memberPW")
    // memberName := c.DefaultPostForm("memberName","")
	// memberMail := c.DefaultPostForm("memberMail","")

	MysqlDB, err := gorm.Open("mysql", "root:@tcp(192.168.99.84:3306)")
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return
	}else{
	    fmt.Println("connect database success")
	    MysqlDB.SingularTable(true)
	    fmt.Println("create table success")
	}
	defer MysqlDB.Close()

	c.JSON(200, gin.H{
		"memberAC": memberAC,
		"memberPW":    memberPW,
	})
}
