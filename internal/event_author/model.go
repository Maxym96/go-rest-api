package event_author

import "rest-app/internal/event"

type EventAuthor struct {
	ID       string      `db:"id"`
	Name     string      `db:"name"`
	AuthorID event.Event `db:"author_id"`
}
