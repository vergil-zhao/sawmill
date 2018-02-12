package connection

import (
	"log"

	"gopkg.in/mgo.v2"
	"vergil.com/practice/sawmill/app/config"
)

var database *mgo.Database

// DB return mongo database
func DB() *mgo.Database {
	if database == nil {
		session, err := mgo.Dial(config.DBHost)
		if err != nil {
			log.Fatalln("Connect to mongodb failed.")
		}
		database = session.DB(config.DBName)
	}
	return database
}
