package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/test_server/internal/domain/event"
	"io/ioutil"
	"log"
	"net/http"
)

type EventController struct {
	service *event.Service
}

var params event.Event

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
			err := errors.New("headers parameter id = \"\". Please, repeat Request with value")
			err = badRequest(w, err)
			if err != nil {
				return
			}
			return
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

		bodyRead, err := ioutil.ReadAll(r.Body)

		if l := len(bodyRead); l == 0 {
			err := errors.New("body does`t have Json. Please, repeat Request with Body (Json)")
			err = badRequest(w, err)
			if err != nil {
				return
			}
			return
		}

		err = json.Unmarshal(bodyRead, &params)
		if err != nil {
			return
		}

		if params.Name == "" || params.Description == "" || params.DateAndTime == "" {
			err := errors.New("name or description or date_and_time don`t have value. All fields ARE REQUIRED!. Please, repeat Request with value")
			err = badRequest(w, err)
			if err != nil {
				return
			}
			return
		}

		events, err := (*c.service).Create(params)
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

		bodyRead, err := ioutil.ReadAll(r.Body)

		if l := len(bodyRead); l == 0 {
			err := errors.New("body does`t have Json. Please, repeat Request with Body (Json)")
			err = badRequest(w, err)
			if err != nil {
				return
			}
			return
		}

		err = json.Unmarshal(bodyRead, &params)
		log.Println(params)
		if err != nil {
			return
		}

		if params.ID == "" {
			err := errors.New("id = \"\". ID is REQUIRED. Please, repeat Request with value")
			err = badRequest(w, err)
			if err != nil {
				return
			}
			return
		}

		events, err := (*c.service).Update(params)
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
			err := errors.New("id = ''. ID is REQUIRED. Please, repeat Request with value")
			err = badRequest(w, err)
			if err != nil {
				return
			}
			return
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
