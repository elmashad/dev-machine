package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/streadway/amqp"
)

func init() {

	connection, err := amqp.Dial("amqp://test1:test1@192.168.99.100:5672")
	if err != nil {
		panic("could not establish AMQP connection: " + err.Error())
	}

	channel, err := connection.Channel()
	if err != nil {
		panic("could not open channel: " + err.Error())
	}

	_, err = channel.QueueDeclare("my_queue", true, false, false, false, nil)
	if err != nil {
		panic("error while declaring the queue: " + err.Error())
	}

	err = channel.QueueBind("my_queue", "#", "events", false, nil)
	if err != nil {
		panic("error while binding the queue: " + err.Error())
	}

	msgs, err := channel.Consume("my_queue", "", false, false, false, false, nil)
	if err != nil {
		panic("error while consuming the queue: " + err.Error())
	}

	for msg := range msgs {
		fmt.Println("message received: " + string(msg.Body))
		msg.Ack(false)
	}

}

func main() {


	fmt.Println("goood")
	//ventHandler := controllers.NewEventHandler()
	r := mux.NewRouter()
	//eventsrouter := r.PathPrefix("/events").Subrouter()
	//eventsrouter.Meethods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(eventHandler.FindEventHandler)
	//eventsrouter.Methods("GET").Path("").HandlerFunc(eventHandler.AllEventHandler)
	//eventsrouter.Methods("POST").Path("").HandlerFunc(eventHandler.NewEventHandler)
	http.ListenAndServe(":8033", r)
}

