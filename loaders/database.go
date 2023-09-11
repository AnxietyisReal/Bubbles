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
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    "admin",
	}))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Connected to database")
	}
}

func FindDoc(collection string, filter bson.D) []bson.M {
	coll := client.Database(dbName).Collection(collection)
	var results []bson.M
	cur, err := coll.Find(context.TODO(), filter)
	if err == mongo.ErrNoDocuments {
		return nil
	}
	if err != nil {
		log.Fatalf("Failed to find document: %v", err)
	}
	if err = cur.All(context.Background(), &results); err != nil {
		log.Fatalf("Failed to decode document: %v", err)
	}
	return results
}

func InsertDoc(collection string, doc bson.D) *mongo.InsertOneResult {
	coll := client.Database(dbName).Collection(collection)
	result, err := coll.InsertOne(context.Background(), doc)
	if err != nil {
		log.Fatalf("Failed to insert document: %v", err)
	}
	return result
}

func InsertMultipleDocs(collection string, docs []interface{}) *mongo.InsertManyResult {
	coll := client.Database(dbName).Collection(collection)
	result, err := coll.InsertMany(context.Background(), docs)
	if err != nil {
		log.Fatalf("Failed to insert documents: %v", err)
	}
	return result
}

func UpdateDoc(collection string, filter bson.D, update bson.D) *mongo.UpdateResult {
	coll := client.Database(dbName).Collection(collection)
	result, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatalf("Failed to update document: %v", err)
	}
	return result
}

func UpdateMultipleDocs(collection string, filter bson.D, update bson.D) *mongo.UpdateResult {
	coll := client.Database(dbName).Collection(collection)
	result, err := coll.UpdateMany(context.Background(), filter, update)
	if err != nil {
		log.Fatalf("Failed to update documents: %v", err)
	}
	return result
}

func DeleteDoc(collection string, filter bson.D) *mongo.DeleteResult {
	coll := client.Database(dbName).Collection(collection)
	result, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatalf("Failed to delete document: %v", err)
	}
	return result
}

func DeleteMultipleDocs(collection string, filter bson.D) *mongo.DeleteResult {
	coll := client.Database(dbName).Collection(collection)
	result, err := coll.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatalf("Failed to delete documents: %v", err)
	}
	return result
}

func FindOneAndUpdate(collection string, filter bson.D, update bson.D) *mongo.SingleResult {
	coll := client.Database(dbName).Collection(collection)
	result := coll.FindOneAndUpdate(context.Background(), filter, update)
	return result
}

func FindOneAndDelete(collection string, filter bson.D) *mongo.SingleResult {
	coll := client.Database(dbName).Collection(collection)
	result := coll.FindOneAndDelete(context.Background(), filter)
	return result
}

func FindOneAndReplace(collection string, filter bson.D, replacement bson.D) *mongo.SingleResult {
	coll := client.Database(dbName).Collection(collection)
	result := coll.FindOneAndReplace(context.Background(), filter, replacement)
	return result
}

func CreateServerSettings(guildID snowflake.ID) *mongo.InsertOneResult {
	coll := client.Database(dbName).Collection("serverSettings")
	result, err := coll.InsertOne(context.Background(), bson.D{{Key: "_id", Value: guildID}})
	if err != nil {
		log.Fatalf("Failed to create server settings: %v", err)
	}
	return result
}

func UpdateServerSettings(guildID snowflake.ID, settings bson.D) *mongo.UpdateResult {
	coll := client.Database(dbName).Collection("serverSettings")
	result, err := coll.UpdateOne(context.Background(), bson.D{{Key: "_id", Value: guildID}}, bson.D{{Key: "$set", Value: settings}})
	if err != nil {
		log.Fatalf("Failed to update server settings: %v", err)
	}
	return result
}

func DeleteServerSettings(guildID snowflake.ID) *mongo.DeleteResult {
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
	err := coll.FindOne(context.Background(), bson.D{{Key: "_id", Value: guildID}}).Decode(&result)
	if err != nil {
		log.Fatalf("Failed to retrieve FS server URL: %v", err)
	}
	return result["panelURL"].(string)
}

type ServerSettings struct {
	GuildID       snowflake.ID     `bson:"_id"`
	LinkedServers FSServers_Object `bson:"linkedServers"`
}

type FSServers_Object struct {
	Servers FSServer_Schema `bson:"urls"`
}

type FSServer_Schema struct {
	PanelURL string `bson:"_id"`
	IsPublic bool   `bson:"isPublic"`
	// If guild admin/manager linked this server to the bot but not ready for public viewing, this can be toggled.
}
