package mongodb

import (
	"context"
	"fmt"
	"github.com/pritunl/mongo-go-driver/mongo"
	"github.com/pritunl/mongo-go-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDb string) (db *mongo.Database, err error) {
	var mongoDBURL string
	var isAuth bool

	if username == "" && password == "" {
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		isAuth = true
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	}
	clientOption := options.Client().ApplyURI(mongoDBURL)
	if isAuth {
		if authDb == "" {
			authDb = database
		}
		clientOption.SetAuth(options.Credential{
			AuthSource: authDb,
			Username:   username,
			Password:   password,
		})
	}

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		return nil, fmt.Errorf("failed	 to connect to Mongo DB error: %v", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping to MongoDB due to error: %v", err)
	}
	return client.Database(database), nil
}
