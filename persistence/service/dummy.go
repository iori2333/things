//go:build !mongo

// This is a template file for configuring service database backend

package service

import db "things/persistence/database"

func (a *Actor) GetDatabase(config map[string]any) db.Interface {
	a.Logger.Printf("Created a dummy database with config %v", config)
	return nil
}
