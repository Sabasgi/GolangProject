package db

import (
	"context"
	"fmt"
	"repogin/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoConnection() (*mongo.Client, error) {
	Coptions := options.Client().ApplyURI("mongodb://localhost:27017")
	Client, CError := mongo.Connect(context.TODO(), Coptions)
	if CError != nil {
		fmt.Println("ERROR : GetMongoConnection(", CError)
		return &mongo.Client{}, CError
	}
	pingerror := Client.Ping(context.TODO(), nil)
	if pingerror != nil {
		fmt.Println("ERROR : GetMongoConnection", pingerror)
		return &mongo.Client{}, CError

	}
	fmt.Println("SUCCESS : GetMongoConnection - MONGO CONNECTED")
	return Client, nil

}

type MongoRepo struct {
	Client *mongo.Client
	DBInfo models.DBInfo
}

// inheritance
//
//	func (m *MongoRepo) CreteConnection() {
//		mongo := NewMongoDBRepo()
//	}
func NewMongoDBRepo(mm models.DBInfo) *MongoRepo {
	clientOptions := options.Client().ApplyURI(mm.DSN)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, Cerror := mongo.Connect(ctx, clientOptions)
	if Cerror != nil {
		fmt.Println("ERROR : newMongoDB", Cerror)
	}
	return &MongoRepo{
		Client: client,
		DBInfo: mm,
	}
}
func (m *MongoRepo) CloseConnection() error {
	return m.Client.Disconnect(context.Background())
}
