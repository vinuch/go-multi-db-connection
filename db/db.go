package db

import (
	"context"
	"log"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMySQL() (*sql.DB, error) {
	dsn := "root:rootpassword@tcp(localhost:3306)/mydatabase"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	log.Println("Connected to MySQL successfully!")
	return db, nil
}

func ConnectMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	// if err != nil {
	// 	log.Fatalf("Error creating MongoDB client: %v", err)
	// 	return nil, err
	// }


	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)

		return nil, err
	}

	log.Println("Connected to MongoDB successfully!")
	return client, nil
}