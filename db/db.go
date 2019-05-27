package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
	bsonPrimitive "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	// "encoding/json"
	"log"
	"context"
	"os"
	"fmt"
)

var client *mongo.Client
var dbDetails map[string]string = make(map[string]string)
var dbContext context.Context


func Connect(ctx context.Context) (*mongo.Database, error){
	
	dbContext = ctx

	getConnectionDetails()

	var err error

	client, err = mongo.NewClient(mongoOptions.Client().ApplyURI(dbDetails["dbConn"]))

	if err != nil { 
		log.Fatal("error")
		log.Fatal(err) 
	}

	err = client.Connect(ctx)
	
	if err != nil { log.Fatal(err) }

	
	return client.Database(dbDetails["dbName"]), nil
}

func getConnectionDetails() {
	var dbName string = ""
	var dbConn string = ""

	env := os.Getenv("ENV")

	switch env {
	case "local":
		dbName = os.Getenv("LOCAL_DB_NAME")
		dbConn = fmt.Sprintf("mongodb://%s:%s", os.Getenv("LOCAL_DB_SERVER_URL"), os.Getenv("LOCAL_DB_PORT"))
	}

	dbDetails["dbName"] = dbName
	dbDetails["dbConn"] = dbConn

}


func FindOne(collection string, filter interface{}) (interface{}, error) {
	var v map[string]interface{}

	err := client.Database(dbDetails["dbName"]).Collection(collection).FindOne(dbContext, filter).Decode(&v)

	return v, err
}

func FindByAggregate(collection string, aggregation interface{}) ([]interface{}, error) {
	cur, err := client.Database(dbDetails["dbName"]).Collection(collection).Aggregate(dbContext, aggregation)
	defer cur.Close(dbContext)

	jsonArray := []interface{}{}

	for cur.Next(dbContext) {
	   var result bson.M
	   err := cur.Decode(&result)
	   if err != nil { log.Println("aggregate error:"); log.Fatal(err) }

	   jsonArray = append(jsonArray, result)
	   // do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Println("cur error:");
	  log.Fatal(err)
	}

	return jsonArray, err
}

func GetItems(collection string, filter interface{}) ([]bson.M) {

	cur, err := client.Database(dbDetails["dbName"]).Collection(collection).Find(dbContext, filter)

	if err != nil { log.Fatal(err) }
	defer cur.Close(dbContext)

	jsonArray := []bson.M{}


	for cur.Next(dbContext) {
	   var result bson.M
	   err := cur.Decode(&result)
	   if err != nil { log.Fatal(err) }

	   jsonArray = append(jsonArray, result)
	   // do something with result....
	}
	if err := cur.Err(); err != nil {
	  log.Fatal(err)
	}

	return jsonArray
}

func InsertObj(collection string, jsonBody interface{}) (*mongo.InsertOneResult, error, bsonPrimitive.ObjectID){
	objectId := bsonPrimitive.NewObjectID()

	var bsonObj map[string]interface{}

	var body1, err = bson.Marshal(jsonBody)

	body2, err := bson.Marshal(bson.M{"_id": objectId })
	
	bson.Unmarshal(body1, &bsonObj)
	bson.Unmarshal(body2, &bsonObj)

	log.Println(bsonObj)

	res, err := client.Database(dbDetails["dbName"]).Collection(collection).InsertOne(dbContext, bsonObj)

	return res, err, objectId
}

func InsertMany(collection string, documents []interface{}, options *mongoOptions.InsertManyOptions) (*mongo.InsertManyResult, error){
	res, err := client.Database(dbDetails["dbName"]).Collection(collection).InsertMany(dbContext, documents, options)
	return res,err
}

func UpdateObj(collection string,filter interface{}, jsonBody interface{}) (*mongo.UpdateResult, error){
	return client.Database(dbDetails["dbName"]).Collection(collection).UpdateOne(dbContext, filter, jsonBody)
}

func FindOneAndReplace(collection string,filter interface{}, jsonBody interface{}, options *mongoOptions.FindOneAndReplaceOptions) *mongo.SingleResult {
	return client.Database(dbDetails["dbName"]).Collection(collection).FindOneAndReplace(dbContext, filter, jsonBody, options)
}

func Client() (*mongo.Database, error){

	getConnectionDetails()

	if(client == nil){
		return nil, fmt.Errorf("Client is not connected")
	}

	return client.Database(dbDetails["dbName"]), nil
}