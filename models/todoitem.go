package models

import "gorm.io/gorm"

type ToDoItem struct {
	ID     string `json:"id"`
	Label  string `json:"label"`
	Detail string `json:"detail"`
	IsDone bool   `json:"done"`
}
type ToDoItemDAO struct {
	db *gorm.DB
}

func NewToDoItemDAO(db *gorm.DB) *ToDoItemDAO {
	return &ToDoItemDAO{db: db}
}

func (dao *ToDoItemDAO) GetAllToDoItems() (*[]ToDoItem, error) {
	var items []ToDoItem
	dao.db.Find(&items)
	return &items, nil

}

func (dao *ToDoItemDAO) CreateToDoItem(item *ToDoItem) error {
	return dao.db.Create(item).Error
}

func (dao *ToDoItemDAO) GetToDoItemByID(id int) (*ToDoItem, error) {
	var item ToDoItem
	err := dao.db.First(&item, id).Error
	return &item, err
}
