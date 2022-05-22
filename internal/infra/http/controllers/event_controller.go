package controllers

import (
	"fmt"
	"github.com/test_server/internal/domain/event"
	"log"
	"net/http"
)

type EventController struct {
	service *event.Service
}

func NewEventController(s *event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			events, err := (*c.service).FindAll()
			if err != nil {
				log.Printf("EventController.FindAll(): %s", err)
				err = internalServerError(w, err)
				if err != nil {
					log.Printf("EventController.FindAll(): %s", err)
				}
				return
			}
			err = success(w, events)
			if err != nil {
				log.Printf("EventController.FindAll(): %s", err)
			}
		} else {
			log.Printf("http.Request.Method is not GET")
			return
		}
	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("id")
		if id == "" {
			panic("Headers parameter id = \"\". Please, repeat Request with value")
		}
		log.Printf("id:%s", id)
		eventFind, err := (*c.service).FindOne(id)
		if err != nil {
			log.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				log.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		err = success(w, eventFind)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}
func (c *EventController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		name := r.PostFormValue("name")
		description := r.PostFormValue("description")
		dateAndTime := r.PostFormValue("date_and_time")

		if name == "" || description == "" || dateAndTime == "" {
			panic("Name or description or date_and_time don`t have value. All fields ARE REQUIRED!. Please, repeat Request with value")
		}

		eventNew := event.Event{
			ID:          "",
			Name:        name,
			Description: description,
			DateAndTime: dateAndTime,
		}
		events, err := (*c.service).Create(eventNew)
		if err != nil {
			log.Printf("EventController.Create(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				log.Printf("EventController.Create(): %s", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			log.Printf("EventController.Create(): %s", err)
		}
	}
}

func (c *EventController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PostFormValue("id")
		name := r.PostFormValue("name")
		description := r.PostFormValue("description")
		dateAndTime := r.PostFormValue("date_and_time")

		if id == "" {
			panic("id = \"\". ID is REQUIRED. Please, repeat Request with value")
		}

		event := event.Event{
			ID:          id,
			Name:        name,
			Description: description,
			DateAndTime: dateAndTime,
		}
		events, err := (*c.service).Update(event)
		if err != nil {
			log.Printf("EventController.Update(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				log.Printf("EventController.Update(): %s", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			log.Printf("EventController.Update(): %s", err)
		}
	}
}
func (c *EventController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("id")

		if id == "" {
			panic("id = \"\". ID is REQUIRED. Please, repeat Request with value")
		}

		events, err := (*c.service).Delete(id)
		if err != nil {
			log.Printf("EventController.Delete(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				log.Printf("EventController.Delete(): %s", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			log.Printf("EventController.Delete(): %s", err)
		}
	}
}
