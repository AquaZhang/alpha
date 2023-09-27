package services

import (
	. "main/models"
)

type ToDoItemService struct {
	dao *ToDoItemDAO
}

func NewToDoItemService(dao *ToDoItemDAO) *ToDoItemService {
	return &ToDoItemService{dao: dao}
}

func (s *ToDoItemService) GetAllToDoItems() (*[]ToDoItem, error) {
	return s.dao.GetAllToDoItems()
}

func (s *ToDoItemService) CreateToDoItem(item *ToDoItem) error {
	// 执行业务逻辑，创建待办事项
	// ...
	// 使用 s.dao 执行数据库操作
	if err := s.dao.CreateToDoItem(item); err != nil {
		return err
	}
	// ...
	return nil
}

func (s *ToDoItemService) GetToDoItemByID(id int) (*ToDoItem, error) {
	// 执行业务逻辑，获取待办事项
	// ...
	// 使用 s.dao 执行数据库操作
	return s.dao.GetToDoItemByID(id)
}
