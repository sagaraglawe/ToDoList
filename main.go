package main

import (
	database "./Database"
	"github.com/jinzhu/gorm"
	"github.com/todoList/Employee"
	"github.com/todoList/Task"
	"github.com/todoList/TaskAssignment"
	"github.com/todoList/TaskProgress"
	"github.com/todoList/Users"
	"log"
	"net/http"
)

func testingTest(w http.ResponseWriter, r *http.Request){

}
type DbConn struct {
	Db *gorm.DB
}

func main(){

	db, err := database.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	dbConn := DbConn{Db:db}

	dbConn.Migrations()


	http.HandleFunc("/signin",dbConn.SignIn)
	http.HandleFunc("/signup", dbConn.SignUp)

	http.HandleFunc("/task/create", dbConn.CreateTask)
	http.HandleFunc("/employee/create", dbConn.EmployeeCreate)
	http.HandleFunc("/task/assign", dbConn.AssignTask)
	http.HandleFunc("/task/modify/description",dbConn.ModifyTaskDescription)
	http.HandleFunc("/task/modify/estimate", dbConn.ModifyTaskEstimate)
	http.HandleFunc("/task/delete", dbConn.DeleteTask)


	log.Fatal(http.ListenAndServe(":8000", nil))
}

func (db *DbConn) Migrations(){
	db.Db.AutoMigrate(&Task.Task_Model{})
	db.Db.AutoMigrate(&Employee.Employee_Model{})
	db.Db.AutoMigrate(&TaskAssignment.TASK_ASSIGNMENT_MODEL{})
	db.Db.AutoMigrate(&TaskProgress.Task_Progress_Model{})
	db.Db.AutoMigrate(&Users.User_Model{})
}

func (db *DbConn) SignIn(w http.ResponseWriter, r *http.Request){
	Users.SignIn(w, r, db.Db)
}

func (db *DbConn) SignUp(w http.ResponseWriter, r *http.Request){
	Users.Signup(w, r, db.Db)
}

// Admin Route
func (db *DbConn) CreateTask(w http.ResponseWriter, r *http.Request){
	Task.CreateTask(w, r, db.Db)
}

func (db *DbConn) EmployeeCreate(w http.ResponseWriter, r *http.Request){
	Employee.EmployeeCreate(w, r, db.Db)
}

func (db *DbConn) AssignTask(w http.ResponseWriter, r *http.Request){
	TaskAssignment.AssignTask(w, r, db.Db)
}

func (db *DbConn) ModifyTaskDescription(w http.ResponseWriter, r *http.Request){
	Task.ModifyTaskDescription(w, r, db.Db)
}

func (db *DbConn) ModifyTaskEstimate(w http.ResponseWriter, r *http.Request){
	Task.ModifyTaskEstimate(w, r, db.Db)
}

func (db *DbConn) DeleteTask(w http.ResponseWriter, r *http.Request){
	Task.DeleteTask(w, r, db.Db)
}

func (db *DbConn) DeleteCompletedTask (w http.ResponseWriter, r *http.Request){

}

func (db *DbConn) GetTaskStatus (w http.ResponseWriter, r *http.Request) {

}

func (db *DbConn) GetAllTaskStatus (w http.ResponseWriter, r *http.Request){

}

func (db *DbConn) GetTaskPerUser (w http.ResponseWriter, r *http.Request) {

}


//User Route

func (db *DbConn) ChangeTaskStatus (w http.ResponseWriter, r *http.Request) {

}

func (db *DbConn) LogTime (w http.ResponseWriter, r *http.Request) {

}

func (db *DbConn) GetAssignedTask (w http.ResponseWriter, r *http.Request){

}

func (db *DbConn) GetAssignedTaskStatus (w http.ResponseWriter, r *http.Request) {

}

