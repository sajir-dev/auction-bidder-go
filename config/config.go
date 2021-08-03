package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	CONN_STRING = "mongodb://localhost:27017/auctioneer-go"
	DB          = "auctioneer-go"
	auction     = "auctions"
	bidders     = "bidders"
)

type DbCollections struct {
	AuctionCollection *mongo.Collection
	BiddersCollection *mongo.Collection
}

func ConnectWithDB() (*DbCollections, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONN_STRING))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	auctionCollection := client.Database(DB).Collection(auction)
	biddersCollection := client.Database(DB).Collection(bidders)

	return &DbCollections{
		AuctionCollection: auctionCollection,
		BiddersCollection: biddersCollection,
	}, nil
}
