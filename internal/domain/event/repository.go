package event

import (
	"github.com/upper/db/v4"
	"log"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id string) (*Event, error)
	Create(event Event) (*Event, error)
	Update(event Event) (*Event, error)
	Delete(id string) (bool, error)
}

type repository struct {
	client db.Session
}

func NewRepository(client *db.Session) Repository {
	return &repository{
		client: *client,
	}
}

func (r *repository) FindAll() ([]Event, error) {

	var eventArr []Event
	eventCol := r.client.Collection("event")
	err := eventCol.Find().All(&eventArr)
	if err != nil {
		return nil, err
	}
	log.Printf("Records in the %q collection:\n", eventCol.Name())
	for i := range eventArr {
		log.Printf("record #%d: %#v\n", i, eventArr[i])
	}
	return eventArr, nil
}

func (r *repository) FindOne(id string) (*Event, error) {

	var eventStr Event
	eventCol := r.client.Collection("event")
	err := eventCol.Find(db.Cond{"id": id}).One(&eventStr)
	if err != nil {
		return nil, err
	}
	return &eventStr, nil
}
func (r *repository) Create(event Event) (*Event, error) {
	eventCol := r.client.Collection("event")
	err := eventCol.InsertReturning(&event)
	if err != nil {
		return nil, err
	}
	log.Printf("New Event Added to db: %v", event)
	return &event, nil
}
func (r *repository) Update(event Event) (*Event, error) {

	eventCol := r.client.Collection("event")
	err := eventCol.Find(db.Cond{"id": event.ID}).Update(&event)
	if err != nil {
		return nil, err
	}
	if event.Name == "" || event.Description == "" || event.DateAndTime == "" {
		err := eventCol.Find(db.Cond{"id": event.ID}).One(&event)
		if err != nil {
			return nil, err
		}
	}
	log.Printf("Event updated to db: %v", event)
	return &event, nil
}
func (r *repository) Delete(id string) (bool, error) {
	eventCol := r.client.Collection("event")
	err := eventCol.Find(db.Cond{"id": id}).Delete()
	if err != nil {
		return false, err
	}
	log.Printf("Event (ID = %s) deleted in db", id)
	return true, nil
}
