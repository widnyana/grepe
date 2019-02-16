package dal

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/widnyana/grepe/config"
)

//Mongo is a datasource.
type Mongo struct {
	Client   *mongo.Client
	Database *mongo.Database
	context  context.Context
}

//NewMongo creates a newinstance of Mongo
func NewMongo(cnf config.Config) (*Mongo, error) {

	var (
		uri string
		ctx = context.TODO()
	)
	if len(cnf.Mongo.Username) > 0 && len(cnf.Mongo.Password) > 0 {
		uri = fmt.Sprintf(`mongodb://%s:%s@%s:%d/%s`,
			cnf.Mongo.Username,
			cnf.Mongo.Password,
			cnf.Mongo.Host,
			cnf.Mongo.Port,
			cnf.Mongo.Database,
		)
	} else {
		uri = fmt.Sprintf(`mongodb://%s:%d/%s`,
			cnf.Mongo.Host,
			cnf.Mongo.Port,
			cnf.Mongo.Database,
		)
	}

	if len(cnf.Mongo.Options) > 0 {
		uri = fmt.Sprintf(`%s?%s`, uri, cnf.Mongo.Options)
	}

	client, err := mongo.NewClient(uri)
	if err != nil {
		log.Printf("The mongodb URL is incorrect: %s", uri)
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Print("Could not use context")
		return nil, err
	}

	db := client.Database(cnf.Mongo.Database)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Could not contact %v via port %d", cnf.Mongo.Host, cnf.Mongo.Port)
		return nil, err
	}

	return &Mongo{Client: client, Database: db, context: ctx}, nil
}

// Disconnect a Mongo client
func (m *Mongo) Disconnect() error {
	err := m.Client.Disconnect(m.context)
	if err != nil {
		log.Printf("Can not close the connection")
		return err
	}
	return nil
}
