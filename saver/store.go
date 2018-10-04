package main

import (
	"github.com/go-pkgz/mongo"
	"time"
)

type RecordAccessor interface {
	//can be implemented with different storage
	saveRecord(record Record) error
}

func makeDataStore() (result RecordAccessor, err error) {
	//can be extended by different types of store
	mgServer, err := mongo.NewServerWithURL(MONGO_URL, 10*time.Second) // todo get mongo url from env vars
	if err != nil {
		panic(err)
	}
	conn := mongo.NewConnection(mgServer, MONGO_URL_DB, "")
	return NewMongo(conn)
}
