package event

import (
	"context"
	db2 "github.com/upper/db/v4"
	"log"
	"rest-app/internal/event"
	"rest-app/package/logging"
)

type db struct {
	client db2.Session
	logger *logging.Logger
}

func (d *db) Create(ctx context.Context, event event.CreateEventDTO) error {
	eventCol := d.client.Collection("event")
	err := eventCol.InsertReturning(&event)
	if err != nil {
		log.Fatal("EventAuthors.Find: ", err)
	}
	log.Printf("Account: %#v", event)
	return nil
}

func (d *db) FindAll(ctx context.Context) ([]event.Event, error) {
	eventCol := d.client.Collection("event")
	var eventStr []event.Event
	err := eventCol.Find().All(&eventStr)
	if err != nil {
		log.Fatal("EventAuthors.Find: ", err)
	}
	log.Printf("Records in the %q collection:\n", eventCol.Name())
	for i := range eventStr {
		log.Printf("record #%d: %#v\n", i, eventStr[i])
	}
	return eventStr, nil
}

func (d *db) FindOne(ctx context.Context, id string) (event.Event, error) {
	eventCol := d.client.Collection("event")
	var eventStr event.Event
	err := eventCol.Find(db2.Cond{"id": id}).One(&eventStr)
	if err != nil {
		log.Fatal("EventAuthors.Find: ", err)
	}
	return eventStr, nil
}

func (d *db) Update(ctx context.Context, event event.Event) error {
	eventCol := d.client.Collection("event")
	err := eventCol.Find("id", event.ID).Update(event)
	if err != nil {
		log.Fatal("EventAuthors.Find: ", err)
	}
	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	eventCol := d.client.Collection("event")
	err := eventCol.Find("id", id).Delete()
	if err != nil {
		log.Fatal("EventAuthors.Find: ", err)
	}
	return nil
}

func NewRepository(client *db2.Session, logger *logging.Logger) event.Repository {
	return &db{
		client: *client,
		logger: logger,
	}
}
