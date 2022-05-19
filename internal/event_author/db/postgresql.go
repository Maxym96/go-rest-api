package event_author

import (
	"context"
	db2 "github.com/upper/db/v4"
	"log"
	"rest-app/internal/event_author"
	"rest-app/package/logging"
)

type db struct {
	client db2.Session
	logger *logging.Logger
}

func (d *db) Create(ctx context.Context, author *event_author.EventAuthor) error {
	d.logger.Trace("SQL Query: %s")
	return nil
}

func (d *db) FindAll(ctx context.Context) ([]event_author.EventAuthor, error) {
	eventAuthorCol := d.client.Collection("event_author")

	// Uncomment the following line (and the github.com/upper/db import path) to
	// write SQL statements to os.Stdout:
	// db.LC().SetLevel(db.LogLevelDebug)

	// Find().All() maps all the records from the books collection.
	var eventAuthorStr []event_author.EventAuthor
	err := eventAuthorCol.Find().All(&eventAuthorStr)
	if err != nil {
		log.Fatal("EventAuthors.Find: ", err)
	}

	// Print the queried information.
	log.Println(eventAuthorStr)
	log.Printf("Records in the %q collection:\n", eventAuthorCol.Name())
	for i := range eventAuthorStr {
		log.Printf("record #%d: %#v\n", i, eventAuthorStr[i])
	}
	return eventAuthorStr, nil
}

func (d *db) FindOne(ctx context.Context, id string) (event_author.EventAuthor, error) {
	return event_author.EventAuthor{}, nil
}

func (d *db) Update(ctx context.Context, author event_author.EventAuthor) error {
	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	collections, err := d.client.Collections()
	if err != nil {
		log.Fatal("Collections: ", err)
	}
	for i := range collections {
		// Name returns the name of the collection.
		log.Printf("-> %q\n", collections[i].Name())
	}
	return nil
}

func NewRepository(client *db2.Session, logger *logging.Logger) event_author.Repository {
	return &db{
		client: *client,
		logger: logger,
	}
}
