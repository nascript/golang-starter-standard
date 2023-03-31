package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const maxTimeout = 60

func (c *Config) InitMongoDBConn() {
	log.Println("Init database connection pool . . .")

	mongoSingleton.Do(func() {
		ctxWT, cancel := context.WithTimeout(context.Background(), maxTimeout*time.Second)
		defer cancel()

		option := options.Client().
			ApplyURI(c.MongoURI).
			SetMinPoolSize(uint64(c.MongoPoolMin)).
			SetMaxPoolSize(uint64(c.MongoPoolMax)).
			SetMaxConnIdleTime(time.Duration(c.MongoMaxIdleTimeSecond) * time.Second)

		var client *mongo.Client
		var err error
		if client, err = mongo.NewClient(option); err != nil {
			panic(fmt.Sprintf("DATABASE_ERROR: %s", err.Error()))
		}

		if err := client.Connect(ctxWT); err != nil {
			panic(fmt.Sprintf("DATABASE_ERROR: %s", err.Error()))
		}

		DbPool = client.Database(c.MongoDB)

		log.Print("DB conn pool: ready (mongo driver)")
	})
}
