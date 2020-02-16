package datastore

import (
	"fmt"
	"golang-poc-postgress/config"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func NewDB() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", config.C.Database.Host, config.C.Database.Port, config.C.Database.User, config.C.Database.Password, config.C.Database.Dbname)

	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}
