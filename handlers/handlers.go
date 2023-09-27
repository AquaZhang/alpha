package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	dsn := "root:001000@tcp(127.0.0.1:3306)/alpha?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接错误：%v", err)
	}
	DB = db
}
func GetResource(c *gin.Context) {
	getType := c.Param("type")
	switch getType {
	case "getAllToDoItems":
		GetAllToDoItems(c)
	}
}
func CreateResource(c *gin.Context) {
	// 处理创建资源的请求
}

func UpdateResource(c *gin.Context) {
	// 处理更新资源的请求
}

func DeleteResource(c *gin.Context) {
	// 处理删除资源的请求
}
