package controllers

import (
	"fmt"
	"net/http"
)

type bookingService struct {}

func NewBookingHandler() *bookingService {
	return &bookingService{}
}

func (eh *bookingService) AddBookingHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, `{"id":%d}`, 2222)
}
