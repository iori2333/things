package mongo

import (
	"context"

	db "things/persistence/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	name       string
	ctx        context.Context
	collection *mongo.Collection
}

func (c *Collection) Name() string {
	return c.name
}

func (c *Collection) Insert(obj any) (id db.Key, err error) {
	result, err := c.collection.InsertOne(c.ctx, obj, nil)
	if err != nil {
		return "", err
	}
	return result.InsertedID, nil
}

func (c *Collection) InsertMany(obj []any) (ids []db.Key, err error) {
	result, err := c.collection.InsertMany(c.ctx, obj, nil)
	if err != nil {
		return nil, err
	}
	return result.InsertedIDs, nil
}

func (c *Collection) Modify(filter db.Document, obj db.Document) (modified db.Key, err error) {
	result, err := c.collection.UpdateOne(c.ctx, filter, obj)
	if err != nil {
		return
	}
	return result.UpsertedID, nil
}

func (c *Collection) ModifyKey(key string, obj db.Document) (modified db.Key, err error) {
	result, err := c.collection.UpdateByID(c.ctx, key, obj)
	if err != nil {
		return
	}
	return result.UpsertedID, nil
}

func (c *Collection) Query(filter db.Document, limit int64) (objs []db.Document, err error) {
	var lim *int64
	if limit <= 0 {
		lim = nil
	} else {
		lim = &limit
	}
	result, err := c.collection.Find(c.ctx, filter, &options.FindOptions{
		Limit: lim,
	})
	if err != nil {
		return nil, err
	}
	err = result.All(c.ctx, objs)
	return
}

func (c *Collection) QueryKey(key string) (obj db.Document, err error) {
	result := c.collection.FindOne(c.ctx, bson.M{"_id": key}, nil)
	queryErr := result.Err()
	if queryErr != nil {
		if queryErr == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, queryErr
	}
	err = result.Decode(&obj)
	return
}

func (c *Collection) Replace(key string, obj db.Document) (err error) {
	_, err = c.collection.ReplaceOne(c.ctx, bson.M{"_id": key}, obj, nil)
	return
}

func (c *Collection) Delete(filter db.Document) (deleted []db.Document, err error) {
	deleted, err = c.Query(filter, 0)
	if err != nil || len(deleted) == 0 {
		return
	}
	_, err = c.collection.DeleteMany(c.ctx, filter, nil)
	return
}

func (c *Collection) DeleteKey(key string) (deleted db.Document, err error) {
	deleted, err = c.QueryKey(key)
	if err != nil || deleted == nil {
		return
	}
	_, err = c.collection.DeleteOne(c.ctx, bson.M{"_id": key}, nil)
	return
}
