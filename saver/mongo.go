package main

import (
	"github.com/globalsign/mgo"
	"github.com/go-pkgz/mongo"
)

type Mongo struct {
	conn *mongo.Connection
}

const (
	collectionRecords = "records"
)

func NewMongo(conn *mongo.Connection) (*Mongo, error) {
	result := Mongo{conn: conn}
	return &result, nil
}

func (m *Mongo) saveRecord(record Record) error {
	return m.conn.WithCustomCollection(collectionRecords, func(coll *mgo.Collection) error {
		_, err := coll.Upsert(record.email, record) // update record by email but maybe it should be composing key
		return err
	})
}
