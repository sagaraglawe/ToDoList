package TaskProgress

import (
	"github.com/jinzhu/gorm"
	"github.com/todoList/Task"
)

func CreateTaskProgress(task Task.Task_Model, db *gorm.DB){
	var taskProgress Task_Progress_Model

	taskProgress.Task = task

	db.Create(&taskProgress)
}