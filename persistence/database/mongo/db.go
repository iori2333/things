package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	db "things/persistence/database"
)

type Database struct {
	ctx      context.Context
	client   *mongo.Client
	database *mongo.Database
	options  *options.ClientOptions
}

func (db *Database) Connect() error {
	if db.client != nil {
		return nil
	}
	client, err := mongo.NewClient(db.options)
	if err != nil {
		return err
	}
	db.client = client
	if err := client.Connect(db.ctx); err != nil {
		return err
	}
	if err := db.client.Ping(db.ctx, nil); err != nil {
		return err
	}
	return nil
}

func (db *Database) Context() context.Context {
	return db.ctx
}

func (db *Database) Collection(name string) db.Collection {
	coll := db.database.Collection(name, nil)
	return &Collection{
		name:       name,
		collection: coll,
	}
}

func New(ctx context.Context, uri string) db.Interface {
	opts := &options.ClientOptions{}
	return &Database{
		ctx:     ctx,
		options: opts.ApplyURI(uri),
	}
}
