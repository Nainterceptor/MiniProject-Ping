package entity

import (
	"time"

	"github.com/Nainterceptor/MiniProject-Ping/config"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	getPingCollection().EnsureIndex(mgo.Index{
		Key:      []string{"origin"},
	})
}

func getPingCollection() *mgo.Collection {
	return config.MongoDB.C("pings")
}

type Ping struct {
	Id               bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedAt        time.Time     `json:"created_at"`
	Origin           string        `json:"origin,omitempty" bson:",omitempty"`
	NameLookupTimeMs string        `json:"name_lookup_time_ms,omitempty" bson",omitempty"`
	ConnectTimeMs    string        `json:"connect_time_ms,omitempty" bson",omitempty"`
	TransferTimeMs   string        `json:"transfer_time_ms,omitempty" bson",omitempty"`
	TotalTimeMs      string        `json:"total_time_ms,omitempty" bson",omitempty"`
	Status           string        `json:"status,omitempty" bson",omitempty"`
}

func PingNew() *Ping {
	ping := new(Ping)
	ping.Id = bson.NewObjectId()
	ping.CreatedAt = time.Now()

	return ping
}
