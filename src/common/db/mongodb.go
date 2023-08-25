package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var MongoClient *mongo.Client
var MongoDB *mongo.Database
var UserCollection *mongo.Collection


func InitMongoDB()error{
	  // MongoDB 연결 정보 설정
    clientOptions := options.Client().ApplyURI("mongodb://AGscK7APmSeDkftc:YL6BWYsf5uaK0vXt@localhost:27018/?authSource=admin&replicaSet=rs0&w=majority&readPreference=primaryPreferred&retryWrites=true&directConnection=true&ssl=false")
	var err error
    // MongoDB 클라이언트 생성
    MongoClient, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // 연결 확인
    err = MongoClient.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB!")
	databaseName := "ryan_dev"
	mongo := MongoClient.Database(databaseName)
	UserCollection = mongo.Collection("user")
	
	return nil
}