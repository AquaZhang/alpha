package handlers

import (
	"main/models"
	"main/services"

	"github.com/gin-gonic/gin"
)

func GetAllToDoItems(c *gin.Context) {
	db := DB

	dao := models.NewToDoItemDAO(db)
	svc := services.NewToDoItemService(dao)
	items, _ := svc.GetAllToDoItems()
	c.JSON(200, gin.H{
		"message": "success!",
		"data":    items,
	})
}
