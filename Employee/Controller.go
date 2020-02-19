package Employee

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
)

func EmployeeCreate(w http.ResponseWriter, r *http.Request, db *gorm.DB){
	employee:= &Employee_Model{}

	err:= json.NewDecoder(r.Body).Decode(&employee)

	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db.Create(&employee)
}
