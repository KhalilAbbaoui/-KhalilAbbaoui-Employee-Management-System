package config

import (
    "context"
    //"log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var employeeCollection *mongo.Collection

//InitDB initializes the connection to MongoDB
func InitDB() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    var err error
    client, err = mongo.Connect(ctx, clientOptions)
    if err != nil {
        return err
    }

    //Checking the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        return err
    }

    employeeCollection = client.Database("employee_management").Collection("employees")
    return nil
}

func GetEmployeeCollection() *mongo.Collection {
    return employeeCollection
}
