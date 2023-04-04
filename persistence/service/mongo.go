//go:build mongo

package service

import (
	"context"
	db "things/persistence/database"
	"things/persistence/database/mongo"
)

func (a *Actor) GetDatabase(config map[string]any) db.Interface {
	once.Do(func() {
		uri, ok := config["uri"]
		if !ok {
			a.Logger.Fatal("MongoDB: uri not specified in configuration file")
		}
		instance = mongo.New(context.TODO(), uri.(string))
		if err := instance.Connect(); err != nil {
			a.Logger.Fatal(err.Error())
		}
		a.Logger.Printf("Created mongodb connection to %s\n", uri)
	})
	return instance
}
