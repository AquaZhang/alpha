package handlers

import (
	"main/models"
	"main/services"

	"github.com/gin-gonic/gin"
)

func CreateToDoItem(c *gin.Context) {
	db := DB
	dao := models.NewToDoItemDAO(db)
	svc := services.NewToDoItemService(dao)
	// 解析请求体
	var postData models.ToDoItem
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := svc.CreateToDoItem(&postData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success!"})

	// 调用 svc函数 参数为 *ToDoItem

}
