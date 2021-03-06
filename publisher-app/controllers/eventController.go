package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"publisher-app/models"
	"strings"
)

type eventServiceHandler struct{}

var event models.Event

func NewEventHandler() *eventServiceHandler {
	return &eventServiceHandler{}
}

func (eh *eventServiceHandler) FindEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search criteria found, you can either search by id via /id/4
                   to search by name via /name/coldplayconcert}`)
		return
	}
	searchkey, ok := vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search keys found, you can either search by id via /id/4
                   to search by name via /name/coldplayconcert}`)
		return
	}
	var err error
	switch strings.ToLower(criteria) {
	case "name":
		err = event.FindEventByName(searchkey)
	case "id":
		err = event.FindEventById(searchkey)
	}
	if err != nil {
		fmt.Fprintf(w, "{error %s}", err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)
}

func (eh *eventServiceHandler) AllEventHandler(w http.ResponseWriter, r *http.Request) {

	err, users := event.FindEvent()

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying to find all available events %s}", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf8")

	err = json.NewEncoder(w).Encode(&users)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying encode events to JSON %s}", err)
	}
}

func (eh *eventServiceHandler) NewEventHandler(w http.ResponseWriter, r *http.Request) {

	event := models.NewEvent()
	err := json.NewDecoder(r.Body).Decode(&event)
	if nil != err {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "error occured while decoding event data %s"}`, err)
		return
	}

	err = event.AddEvent()

	if nil != err {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "error occured while persisting event %d %s"}`, event.ID, err)
		return
	}

	fmt.Fprint(w, `{"id":%d}`, event.ID)
}
