package Database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

const (
	host = "localhost"
	port = 5432
	user = "sagargovindaaglawe"
	dbName = "todolist"
)

func InitDB() (*gorm.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbName)

	db, err := gorm.Open("postgres", psqlInfo)

	if err!=nil {
		log.Fatal(err)
	}
	return db, err
}

