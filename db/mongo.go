package db

import mgo "gopkg.in/mgo.v2"

type MongoConnection struct {
	Session *mgo.Session
}

var mongoConnection *MongoConnection

func GetMongoDBConnection(connectionString string) Connection {
	if mongoConnection == nil {
		session, err := mgo.Dial(connectionString)

		if err != nil {
			panic(err)
		}

		mongoConnection = &MongoConnection{Session: session}
	}

	return mongoConnection
}

func (c *MongoConnection) GetObject(keyName, keyValue string) interface{} {
	return nil
}

func (c *MongoConnection) PutObject(object interface{}) {}

func (c *MongoConnection) DeleteObject(keyName, keyValue string) {}
