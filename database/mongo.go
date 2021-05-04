package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client       *mongo.Client
	ReadTimeout  uint
	WriteTimeout uint
	Database     string
}

var GlobalDB MongoDB = MongoDB{}

// InstantiateClient intialises the client. URI should include mongodb://
func (m *MongoDB) InstantiateClient(uri string, userName string, password string, database string, authsource string, readTimeout uint, writeTimeout uint) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(readTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{AuthSource: authsource,
		Username: userName,
		Password: password}).ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("InstantiateClient: %v", err)
	}
	m.Client = client
	m.ReadTimeout = readTimeout
	m.WriteTimeout = writeTimeout
	m.Database = database
	return nil
}

func (m *MongoDB) CheckConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.ReadTimeout)*time.Second)
	defer cancel()
	return m.Client.Ping(ctx, readpref.Primary())
}

func (m *MongoDB) DisconnectClient() error {
	return m.Client.Disconnect(context.Background())
}

func (m *MongoDB) GetCollection(collection string) *mongo.Collection {
	return m.Client.Database(m.Database).Collection(collection)
}
