package loaders

import (
	"context"
	"log"

	"github.com/disgoorg/snowflake/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	dbName = "Bubbles"
)

func ConnectToDatabase(uri string) {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetReplicaSet("toastyy").SetAuth(options.Credential{
		Username:      TokenLoader("mongoUser"),
		Password:      TokenLoader("mongoPass"),
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    "admin",
	}))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Connected to database")
	}
}

func CreateSettings(guildID snowflake.ID, panelUrl string) *mongo.InsertOneResult {
	coll := client.Database(dbName).Collection("serverSettings")
	result, err := coll.InsertOne(context.Background(), bson.D{{Key: "_id", Value: guildID}, {Key: "server", Value: panelUrl}})
	if err != nil {
		log.Fatalf("Failed to create server settings: %v", err)
	}
	return result
}

/* func UpdateServerSettings(guildID snowflake.ID, settings bson.D) *mongo.UpdateResult {
	coll := client.Database(dbName).Collection("serverSettings")
	result, err := coll.UpdateOne(context.Background(), bson.D{{Key: "_id", Value: guildID}}, bson.D{{Key: "$set", Value: bson.D{{Key: "linkedServers", Value: settings}}}})
	if err != nil {
		log.Fatalf("Failed to update server settings: %v", err)
	}
	return result
} */

func DeleteSettings(guildID snowflake.ID) *mongo.DeleteResult {
	coll := client.Database(dbName).Collection("serverSettings")
	result, err := coll.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: guildID}})
	if err != nil {
		log.Fatalf("Failed to delete server settings: %v", err)
	}
	return result
}

func RetrieveFSServerURL(guildID snowflake.ID) string {
	coll := client.Database(dbName).Collection("serverSettings")
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: guildID}}).Decode(&result)
	if err != nil {
		log.Fatalf("Failed to retrieve FS server URL: %v", err)
	}
	log.Printf("%v requested \"%v\"", guildID, result["server"].(string))
	if result["server"] != nil {
		return result["server"].(string)
	} else {
		return mongo.ErrNoDocuments.Error()
	}
}
