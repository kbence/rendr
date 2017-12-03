package db

import (
	"fmt"
	"strings"

	"github.com/kbence/rendr/util"
)

type Connection interface {
	GetObject(keyName, keyValue string) interface{}
	PutObject(object interface{})
	DeleteObject(keyName, keyValue string)
}

var defaultConnectionString = "mongodb://localhost/rendr"

func GetConnection() Connection {
	connectionString := util.GetEnvWithDefault("RENDR_DB_URL", defaultConnectionString)

	switch {
	case strings.HasPrefix(connectionString, "mongodb://"):
		return GetMongoDBConnection(connectionString)
	}

	panic(fmt.Sprintf("Cannot interpret connection string '%s'", connectionString))
}
