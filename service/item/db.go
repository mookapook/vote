package vote

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MapOfClient map[string]*mongo.Client

var clients = make(MapOfClient)

func setHost(host *string) {
	if strings.HasPrefix(*host, "mongodb://") {
		*host = strings.Replace(*host, "mongodb://", "", -1)
	}
}

func setKey(key string, host *string) string {
	setHost(host)
	return key + ":" + *host
}

func ClientV2(host string, keys ...string) (*mongo.Client, error) {
	// log.Printf("Client.clients: %#v", clients)
	key := ""
	if len(keys) > 0 {
		key = keys[0]
		// } else {
		// key = fmt.Sprint(time.Now().UnixMilli())
	}
	k := setKey(key, &host)

	// log.Printf("Client.k: %s", k)
	if _, ok := clients[k]; ok {
		// log.Printf("MatchClient[%s]: %p", k, clients[k])
		return clients[k], nil
	}

	// if _, ok := clients[k]; !ok {

	// 	// log.Println("host:", host)
	dbhost := fmt.Sprintf("mongodb://%s", host)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client()
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(30)
	clientOptions.SetMaxConnIdleTime(0)
	clientOptions.SetMaxConnecting(3)
	clientOptions.ApplyURI(dbhost)
	tM := reflect.TypeOf(bson.M{})
	registry := bson.NewRegistryBuilder().RegisterTypeMapEntry(bsontype.EmbeddedDocument, tM).Build()
	// registry := bson.NewRegistryBuilder().Build()
	clientOptions.SetRegistry(registry)
	// clientOpts := options.Client().ApplyURI(SOMEURI).SetAuth(authVal).SetRegistry(reg)
	// client, err := mongo.Connect(ctx, clientOpts)

	// var err error
	// var c *mongo.Client
	if c, err := mongo.Connect(ctx, clientOptions); err != nil {
		// log.Printf("Connect.error:%s of %p", err.Error(), c)
		return nil, err
	} else {
		if key == "" {
			k = fmt.Sprintf("%p", c)
		}
		// log.Printf("SetClient[%s] to %p", k, c)
		clients[k] = c
	}
	return clients[k], nil

	// }

	// return clients[k], nil
	// return nil, nil
}

func CloseV2(c *mongo.Client) error {
	for k, v := range clients {
		// if v.NumberSessionsInProgress()
		if c == v {
			// log.Printf("v[%s]: %p\n", k, v)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := clients[k].Disconnect(ctx); err != nil {
				return err
			}
			c = nil
			delete(clients, k)
		}
	}
	return nil
}
