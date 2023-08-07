package db

import (
	"github.com/SharanM7/hotelreservation/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStorer struct {
	Client *mongo.Client
}

func (ms *MongoStorer) AddUser() error {
	return nil
}

func (ms *MongoStorer) GetUser(id string) (*types.User, error) {
	ms.Client.Database("")
	return &types.User{}, nil
}
