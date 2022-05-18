package author

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"rest-app/internal/author"
	"rest-app/package/client/postgresql"
	"rest-app/package/logging"
	"strings"
)

type db struct {
	client postgresql.Client
	logger *logging.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (d *db) Create(ctx context.Context, author *author.Author) error {
	q := `
		INSERT INTO author (name) 
		VALUES ($1)
		RETURNING id
	`
	d.logger.Tracef("SQL Query: %s", formatQuery(q))
	row := d.client.QueryRow(ctx, q, author.Name)
	if err := row.Scan(&author.ID); err != nil {
		if pgErr, ok := err.(*pgx.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Where: %s, Detail: %s, Code: %s, SQL State: %s", pgErr.Message, pgErr.Where, pgErr.Detail, pgErr.Code, pgErr.SQLState()))
			d.logger.Error(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (d *db) FindAll(ctx context.Context) ([]author.Author, error) {
	q := `
		SELECT id, name FROM public.author;
	`
	d.logger.Tracef("SQL Query: %s", formatQuery(q))
	rows, err := d.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	authors := make([]author.Author, 0)

	for rows.Next() {
		var ath author.Author
		err := rows.Scan(&ath.ID, &ath.Name)
		if err != nil {
			return nil, err
		}

		authors = append(authors, ath)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return authors, nil

}

func (d *db) FindOne(ctx context.Context, id string) (author.Author, error) {
	q := `
		SELECT id, name FROM public.author WHERE id = $1;
	`
	d.logger.Tracef("SQL Query: %s", formatQuery(q))

	var ath author.Author
	row := d.client.QueryRow(ctx, q, id)
	err := row.Scan(&ath.ID, &ath.Name)
	if err != nil {
		return author.Author{}, err
	}
	return ath, nil
}

func (d *db) Update(ctx context.Context, author author.Author) error {
	//TODO implement me
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) author.Repository {
	return &db{
		client: client,
		logger: logger,
	}
}
