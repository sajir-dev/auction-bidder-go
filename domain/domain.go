package domain

import (
	"context"
	"fmt"

	"github.com/sajir-dev/auction-bidder-go/config"
	"go.mongodb.org/mongo-driver/bson"
)

func collections() (*config.DbCollections, error) {
	return config.ConnectWithDB()
}

func InsertOne() error {
	collections, _ := collections()

	res, err := collections.BiddersCollection.InsertMany(context.TODO(), []interface{}{
		bson.D{{"id", "a1"}},
		bson.D{{"id", "a2"}},
		bson.D{{"id", "b1"}},
		bson.D{{"id", "b2"}},
		bson.D{{"id", "c1"}},
		bson.D{{"id", "c2"}},
	})

	if err != nil {
		return err
	}

	fmt.Printf("inserted document with ID %v\n", res.InsertedIDs...)
	return nil
}

func List() {

}
