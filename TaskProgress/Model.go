package TaskProgress

import (
	"github.com/jinzhu/gorm"
	"github.com/todoList/Task"
)

type Task_Progress_Model struct {
	gorm.Model
	Task Task.Task_Model `gorm:"foreignkey:Task_id"`
	Task_id int
	Status string
	Log_time int
}
