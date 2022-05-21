package db

import "github.com/test_server/internal/domain/event"

type Storage interface {
	FindAll() ([]event.Event, error)
	FindOne(id string) (*event.Event, error)
	Create() (*event.Event, error)
	Update(event event.Event) (event.Event, error)
	Delete(id string) (bool, error)
}
