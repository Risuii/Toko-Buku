package main

import (
	"Toko/controllers"
	"Toko/services"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	tokoService    services.Kontrak
	tokoController controllers.ControllerToko
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection established")

	usercollection = mongoclient.Database("store").Collection("Toko")
	tokoService = services.CallHandlerToko(usercollection, ctx)
	tokoController = controllers.New(tokoService)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	tokoController.Routing(basepath)
	log.Fatal(server.Run(":3000"))
}
