package TaskAssignment

import (
	"github.com/jinzhu/gorm"
	"github.com/todoList/Employee"
	"github.com/todoList/Task"
)

type TASK_ASSIGNMENT_MODEL struct {
	gorm.Model
	Task Task.Task_Model `gorm:"foreignkey:Task_id"`
	Task_id int
	Employee Employee.Employee_Model `gorm:"foreignkey:Employee_id"`
	Employee_id int
}
