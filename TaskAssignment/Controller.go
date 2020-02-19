package TaskAssignment

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
)

func AssignTask (w http.ResponseWriter, r *http.Request, db *gorm.DB){

	taskAssign := &TASK_ASSIGNMENT_MODEL{}

	err:= json.NewDecoder(r.Body).Decode(&taskAssign)

	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db.Create(&taskAssign)
}
