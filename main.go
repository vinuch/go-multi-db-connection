package main

import (
    "context"
    "fmt"
    "log"

    "multi-db-connection/db"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
    // Connect to MySQL
    mysqlDB, err := db.ConnectMySQL()
    if err != nil {
        log.Fatalf("Could not connect to MySQL: %v", err)
    }
    defer mysqlDB.Close()

    // Connect to MongoDB
    mongoClient, err := db.ConnectMongoDB()
    if err != nil {
        log.Fatalf("Could not connect to MongoDB: %v", err)
    }
    defer mongoClient.Disconnect(context.Background())

    // Query MySQL
    rows, err := mysqlDB.Query("SELECT id, name FROM users")
    if err != nil {
        log.Fatalf("Error querying MySQL: %v", err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            log.Fatalf("Error scanning MySQL row: %v", err)
        }
        fmt.Printf("MySQL User: %d, %s\n", id, name)
    }

    // Query MongoDB
    collection := mongoClient.Database("testdb").Collection("users")
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        log.Fatalf("Error querying MongoDB: %v", err)
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var user bson.M
        if err := cursor.Decode(&user); err != nil {
            log.Fatalf("Error decoding MongoDB document: %v", err)
        }
        fmt.Printf("MongoDB User: %v\n", user)
    }
}
