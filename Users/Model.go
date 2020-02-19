package Users

import "github.com/jinzhu/gorm"

type User_Model struct {
	gorm.Model
	Password string
	Username string
	Role 	 string
}