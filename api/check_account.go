package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func CheckAccount(c *gin.Context) {
	account := c.DefaultPostForm("account", "")

	var (
		rows  *sql.Rows
		row   *sql.Row
		err   error
		db    *sql.DB
		count int
	)

	rows, err = db.QueryContext(c, "SELECT account FROM members_account WHERE account=?", account)
	row = db.QueryRowContext(c, "SELECT COUNT(0) FROM members_account WHERE account=?", account)
	row.Scan(&count)
	if count != 0 {
		c.JSON(200, gin.H{
			"errorCode": "0",
			"errorMsg":  "error",
		})
	}
	if err != nil {
		c.JSON(200, gin.H{
			"errorCode": "0",
			"errorMsg":  "error",
		})
	}

	for rows.Next() {
		
	}
	

	c.JSON(200, gin.H{
		"errorCode": "0",
		"errorMsg":  "success",
	})
}
