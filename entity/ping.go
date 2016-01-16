package entity

import (
	"time"
	"errors"

	"github.com/Nainterceptor/MiniProject-Ping/config"
	"github.com/asaskevich/govalidator"

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

func (p *Ping) Normalize() {
	p.Origin = govalidator.Trim(p.Origin, "")
}

func (p *Ping) Validate() error {
	p.Normalize()
	if p.Origin == "" {
		return errors.New("`origin` is empty")
	}
	return nil
}

func (p *Ping) Insert() error {
	return getPingCollection().Insert(&p)
}