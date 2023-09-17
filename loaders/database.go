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

func CreateServerSettings(guildID snowflake.ID, settings string) *mongo.InsertOneResult {
	coll := client.Database(dbName).Collection("serverSettings")
	result, err := coll.InsertOne(context.TODO(), bson.D{{Key: "_id", Value: guildID}, {Key: "server", Value: settings}})
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

func DeleteServerSettings(guildID snowflake.ID) *mongo.DeleteResult {
	coll := client.Database(dbName).Collection("serverSettings")
	result, err := coll.DeleteOne(context.Background(), ServerSettings{GuildID: guildID})
	if err != nil {
		log.Fatalf("Failed to delete server settings: %v", err)
	}
	return result
}

func RetrieveFSServerURL(guildID snowflake.ID) string {
	coll := client.Database(dbName).Collection("serverSettings")
	var result bson.M
	err := coll.FindOne(context.Background(), ServerSettings{GuildID: guildID}).Decode(&result)
	if err != nil {
		log.Fatalf("Failed to retrieve FS server URL: %v", err)
	}
	log.Print(result["panelURL"].(string))
	return result["panelURL"].(string)
}

type ServerSettings struct {
	GuildID snowflake.ID `bson:"_id"`
	Server  FSServerURL  `bson:"server"`
	// Server  []FSServerURL `bson:"server"`
}

/* type ServerSettings struct {
	GuildID snowflake.ID           	 `bson:"_id"`
	Server  map[string]FSServerURL{} `bson:"server"`
} */
//Q: Should it be map[string]FSServerURL or map[string]FSServerURL{} if I want numbers after "Server"? (e.g. server1, server2, server3)
//A: It should be map[string]FSServerURL{}.
//Q: It doesn't work. It says "expected ';', found '{'"
//A: You need to use a map[string]FSServerURL{} instead of a map[string]FSServerURL.
//Q: But I am already using it.
//A: No, you are not. You are using a map[string]FSServerURL. You need to use a map[string]FSServerURL{}.
//Q: YES I AM USING IT, STOP BEING MEAN TO ME!
//A: NO, YOU ARE NOT USING IT. YOU ARE USING A map[string]FSServerURL. YOU NEED TO USE A map[string]FSServerURL{}.

//Q: Create a struct for server object and re-use it on multiple objects inside same document
//A: Yes, but it's not that simple. You need to use a map[string]struct{} instead of a struct{}.
//Q: Why?
//A: Because Go doesn't allow you to use the same struct{} on multiple objects inside the same document.

type FSServerURL map[string]struct {
	PanelURL string `bson:"_id"`
}

/* type ServerSettings struct {
	GuildID       snowflake.ID     `bson:"_id"`
	LinkedServers FSServers_Object `bson:"linkedServers"`
}

type FSServers_Object struct {
	Servers map[string]FSServer_Schema `bson:"servers"`
}

type FSServer_Schema struct {
	PanelURL string `bson:"_id"`
} */
