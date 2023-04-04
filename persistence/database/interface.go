package db

import "context"

type Key = any

type Document map[string]any

type Interface interface {
	Connect() error
	Context() context.Context
	Collection(name string) Collection
}

type Collection interface {
	Name() string

	Insert(obj any) (id Key, err error)
	InsertMany(obj []any) (ids []Key, err error)

	Modify(filter Document, obj Document) (modified Key, err error)
	ModifyKey(key string, obj Document) (modified Key, err error)

	Query(filter Document, limit int64) (objs []Document, err error)
	QueryKey(key string) (obj Document, err error)

	Replace(key string, obj Document) (err error)

	Delete(filter Document) (deleted []Document, err error)
	DeleteKey(key string) (deleted Document, err error)
}
