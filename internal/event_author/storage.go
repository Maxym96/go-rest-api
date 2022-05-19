package event_author

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, eventAuthor *EventAuthor) error
	FindAll(ctx context.Context) ([]EventAuthor, error)
	FindOne(ctx context.Context, id string) (EventAuthor, error)
	Update(ctx context.Context, eventAuthor EventAuthor) error
	Delete(ctx context.Context, id string) error
}
