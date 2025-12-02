package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	

}
