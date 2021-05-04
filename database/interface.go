package database

import "go.mongodb.org/mongo-driver/mongo"

type DB interface {
	InstantiateClient(uri string, userName string, password string, database string, authSource string, readTimeout uint, writeTimeout uint) error
	CheckConnection() error
	DisconnectClient() error
	GetCollection(collection string) *mongo.Collection
}
