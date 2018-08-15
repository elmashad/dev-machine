package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
	"github.com/twinj/uuid"
)

type Event struct {
	ID        bson.ObjectId `bson:"_id"`
	Id        string         `json:"id"`
	Name      string        `json:"name"`
	Duration  int           `json:"duration"`
	CreatedAt time.Time     `json:"created_at"`
	StartDate int64         `json:"startDate"`
	EndDate   int64         `json:"endDate"`
	Location  Location      `json:"location"`
}
type Location struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Country   string `json:"country"`
	OpenTime  int    `json:"opentime"`
	CloseTime int    `json:"clostime"`
	Halls     []Hall `json:"hall"`
}
type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	Capacity int    `json:"capacity"`
}

func NewEvent() *Event {
	return &Event{
		Id: uuid.NewV4().String(),
	}
}
