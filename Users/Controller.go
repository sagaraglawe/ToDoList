package Users

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)


func SignIn(w http.ResponseWriter, r *http.Request, db *gorm.DB){

	creds := &User_Model{}

	err:= json.NewDecoder(r.Body).Decode(creds)

	if err !=nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storedCred := &User_Model{}

	db.Where("Username = ?", creds.Username).Find(&storedCred)

	if storedCred.DeletedAt != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if storedCred.Role != creds.Role {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedCred.Password), []byte(creds.Password))

	if err!=nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, err =w.Write([]byte("SUCCESSFULLY SIGNED IN"))

	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Signup (w http.ResponseWriter, r *http.Request, db *gorm.DB){

	creds := &User_Model{}

	err:= json.NewDecoder(r.Body).Decode(creds)

	if err!=nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password),8)

	user := User_Model{Username:creds.Username, Password:string(hashedPassword), Role:creds.Role}

	db.Create(&user)

	_, err= w.Write([]byte("SUCCESSFULLY SIGNED UP"))

	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
