package models

import (
	"github.com/globalsign/mgo"
	"log"
	"publisher-app/common"
)

type DbConfig struct {
	Server   string
	Database string
}

var Db *mgo.Database
var DbConfigInstance = DbConfig{}

// Establish a connection to database
func (m *DbConfig) Connect() {
	dbConnection := common.GetAppConfig().GetString("DB_MONGO_CONNECTION")
	log.Println("ConnectionString:" + dbConnection)
	session, err := mgo.Dial(dbConnection)
	if err != nil {
		log.Fatal(err)
	}
	Db = session.DB(m.Database)
}
