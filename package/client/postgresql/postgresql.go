package postgresql

import (
	"context"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"rest-app/internal/config"
)

func NewClient(ctx context.Context, st config.StorageConfig) (db *db.Session, err error) {
	settings := postgresql.ConnectionURL{
		Database: st.Database,
		Host:     st.Host,
		User:     st.Username,
		Password: st.Password,
	}
	sess, err := postgresql.Open(settings)

	if err != nil {
		log.Fatal("postgresql.Open: ", err)
	}
	//defer sess.Close()

	if err := sess.Ping(); err != nil {
		log.Fatal("Ping: ", err)
	}
	log.Printf("Successfully connected to database: %q", sess.Name())

	return &sess, nil
}
