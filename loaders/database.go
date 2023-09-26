package loaders

import (
	"bubbles/bot/toolbox"
	"context"
	"fmt"
	"log"
	"strconv"

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

type ServerSettings struct {
	GuildID       snowflake.ID   `bson:"_id"`
	LinkedServers map[int]string `bson:"linkedServers"`
}

func AddServer(guildID snowflake.ID, serverID int, serverURL string) (*mongo.UpdateResult, error) {
	collection := client.Database(dbName).Collection("servers")
	filter := bson.M{"_id": guildID}
	update := bson.M{"$set": bson.M{"linkedServers." + strconv.Itoa(serverID): serverURL}}
	opts := options.Update().SetUpsert(true)
	return collection.UpdateOne(context.Background(), filter, update, opts)
}

func DeleteServer(guildID snowflake.ID, serverID int) (*mongo.UpdateResult, error) {
	collection := client.Database(dbName).Collection("servers")
	filter := bson.M{"_id": guildID}
	update := bson.M{"$unset": bson.M{"linkedServers." + strconv.Itoa(serverID): ""}}
	return collection.UpdateOne(context.Background(), filter, update)
}

func UpdateServer(guildID snowflake.ID, serverID int, serverURL string) (*mongo.UpdateResult, error) {
	collection := client.Database(dbName).Collection("servers")
	server := ServerSettings{
		GuildID:       guildID,
		LinkedServers: map[int]string{serverID: serverURL},
	}
	return collection.UpdateOne(context.Background(), bson.M{"_id": guildID}, bson.M{"$set": server})
}

func GetServer(guildID snowflake.ID, serverID int) (string, error) {
	collection := client.Database(dbName).Collection("servers")
	var server ServerSettings
	filter := bson.M{"_id": guildID}
	err := collection.FindOne(context.Background(), filter).Decode(&server)
	if err != nil {
		return "", err
	}
	serverURL, ok := server.LinkedServers[serverID]
	if !ok {
		return "", fmt.Errorf("server ID %d not found for %s", serverID, toolbox.RESTGuild_Name(guildID, TokenLoader("bot")))
	}
	return serverURL, nil
}

func ListServersForThisGuild(guildID snowflake.ID) string {
	collection := client.Database(dbName).Collection("servers")
	var server ServerSettings
	filter := bson.M{"_id": guildID}
	err := collection.FindOne(context.Background(), filter).Decode(&server)
	if err != nil {
		return ""
	}
	var serverList string
	for serverID := range server.LinkedServers {
		serverList += fmt.Sprintf("%d\n", serverID)
		fmt.Printf("Choosen ID: %v", serverList)
	}
	return serverList
}
