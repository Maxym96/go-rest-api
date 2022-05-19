package event_author

import "rest-app/internal/event"

type CreateEventAuthorDTO struct {
	Name     string      `db:"name"`
	AuthorID event.Event `db:"author_id"`
}
