package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

func GetMyEventCollection() *mgo.Collection {
	return Db.C("myevents")
}

// Insert a row into database
func (row *Event) AddEvent() error {
	row.ID = bson.NewObjectId()
	row.CreatedAt = time.Now()
	return GetMyEventCollection().Insert(&row)

}

func (row *Event) FindEventByName(name string) error {

	err := GetMyEventCollection().Find(bson.M{"name": name}).One(&row)
	return err
}

func (row *Event) FindEventById(id string) error {

	err := GetMyEventCollection().Find(bson.M{"id": id}).One(&row)
	return err
}

func (row *Event) FindEvent() (error,[]Event) {
	var rows []Event
	err := GetMyEventCollection().Find(bson.M{}).All(&rows)
	return err,rows
}
