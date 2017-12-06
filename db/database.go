package db

import (
	"github.com/kbence/rendr/util"
	mgo "gopkg.in/mgo.v2"
)

var defaultConnectionString = "mongodb://localhost:27017/rendr"

func GetDatabase() *mgo.Database {
	connection, err := mgo.Dial(util.GetEnvWithDefault("RENDR_DB_URL", defaultConnectionString))

	if err != nil {
		panic(err)
	}

	return connection.DB("")
}
