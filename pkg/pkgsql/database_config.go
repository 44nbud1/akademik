package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type SqlDB struct {
	databaseModel DatabaseModel
}

func NewSqlDB(model DatabaseModel) *SqlDB {
	return &SqlDB{
		databaseModel: model,
	}
}

func (sqb *SqlDB) GetDatabasePostgres() (*gorm.DB, error) {

	log.Print("DB connecting...")
	db, err := gorm.Open("postgres", fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=%v", sqb.databaseModel.User, sqb.databaseModel.Password, sqb.databaseModel.Host, sqb.databaseModel.Port, sqb.databaseModel.Name, "disable"))
	if err != nil {
		log.Print("Error connecting postgres: ", err)
		return nil, err
	}
	log.Print("DB connected")

	return db, nil
}
