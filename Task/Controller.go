package Task

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
)

func CreateTask (w http.ResponseWriter, r *http.Request, db *gorm.DB){
	task:= &Task_Model{}

	err:= json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	db.Create(&task)
}

func ModifyTaskDescription (w http.ResponseWriter, r *http.Request, db *gorm.DB) {

	task:= &Task_Model{}

	err:= json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db.Model(&task).Update("description", task.Description)
}

func ModifyTaskEstimate (w http.ResponseWriter, r *http.Request, db *gorm.DB){
	task:= &Task_Model{}

	err:= json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db.Model(&task).Update("estimate", task.Estimate)
}

func DeleteTask (w http.ResponseWriter, r *http.Request, db *gorm.DB){
	task := &Task_Model{}

	err:= json.NewDecoder(r.Body).Decode(&task)

	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db.Delete(&task)
}

