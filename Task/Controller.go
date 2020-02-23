package Task

import (
	"encoding/json"
	"fmt"
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
	var task Task_Model

	err:= json.NewDecoder(r.Body).Decode(&task)

	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var findTask Task_Model

	rowsAffected := db.First(&findTask, task.ID).RowsAffected

	if rowsAffected == 0 {

		w.WriteHeader(http.StatusBadRequest)

		_,err=w.Write([]byte(fmt.Sprintf("The task with given Id %d is do not exists",task.ID)))

		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	err=db.Delete(&task).Error

	if err!=nil {
		w.WriteHeader(http.StatusOK)

		_,err = w.Write([]byte(fmt.Sprintf("Successfully deleted the task with Id = %d", task.ID)))

		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
}

func DeleteCompletedTask(w http.ResponseWriter, r *http.Request, db *gorm.DB){

}

