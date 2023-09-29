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

func GetToDoItemById(c *gin.Context) {
	db := DB
	dao := models.NewToDoItemDAO(db)
	svc := services.NewToDoItemService(dao)

	var queryId string
	c.ShouldBindJSON(&queryId)
	res, err := svc.GetToDoItemByID(queryId)
	if err != nil {

	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    res,
	})
}
