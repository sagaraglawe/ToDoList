package Task

import "github.com/jinzhu/gorm"

type Task_Model struct {
	gorm.Model
	Estimate int
	Description string
}
