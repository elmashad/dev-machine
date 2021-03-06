package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"publisher-app/common"
	"publisher-app/controllers"
	"publisher-app/models"
)

func init() {
	common.SetConfig()
	models.DbConfigInstance.Connect()

	fmt.Println("good")
	//connection, err := amqp.Dial("amqp://test1:test1@192.168.99.100:5672")
	//if err != nil {
	//	panic("could not establish AMQP connection: " + err.Error())
	//}
	//
	//channel, err := connection.Channel()
	//if err != nil {
	//	panic("could not open channel: " + err.Error())
	//}
	//
	//err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)
	//if err != nil {
	//	panic(err)
	//}
	//
	//message := amqp.Publishing{
	//	Body: []byte("Hello World"),
	//}
	//
	//err = channel.Publish("events", "some-routing-key", false, false, message)
	//if err != nil {
	//	panic("error while publishing message: " + err.Error())
	//}


}

func main() {

	eventHandler := controllers.NewEventHandler()
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(eventHandler.FindEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(eventHandler.AllEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(eventHandler.NewEventHandler)
	http.ListenAndServe(":"+common.GetAppPort(), r)
}
