package Employee

import "github.com/jinzhu/gorm"

type Employee_Model struct {
	gorm.Model
	Name string
	Company_id string
}
