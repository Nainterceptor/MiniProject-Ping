package entity

import (
	"errors"
	"time"

	"github.com/Nainterceptor/MiniProject-Ping/config"
	"github.com/asaskevich/govalidator"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	getPingCollection().EnsureIndex(mgo.Index{
		Key: []string{"origin"},
	})
}

func getPingCollection() *mgo.Collection {
	return config.MongoDB.C("pings")
}

type Ping struct {
	Id               bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedAt        time.Time     `json:"created_at"`
	Origin           string        `json:"origin,omitempty" bson:",omitempty"`
	NameLookupTimeMs int           `json:"name_lookup_time_ms,omitempty" bson",omitempty"`
	ConnectTimeMs    int           `json:"connect_time_ms,omitempty" bson",omitempty"`
	TransferTimeMs   int           `json:"transfer_time_ms,omitempty" bson",omitempty"`
	TotalTimeMs      int           `json:"total_time_ms,omitempty" bson",omitempty"`
	Status           int           `json:"status,omitempty" bson",omitempty"`
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

func AggregatePingOrigin(origin string, result *[]bson.M) error {
	aggregation := getPingCollection().Pipe([]bson.M{
		{"$match": bson.M{"origin": origin}},
		{"$project": bson.M{"transfer_time_ms": true, "created_at": true}},
		{"$group": bson.M{
			"_id": bson.M{
				"year":  bson.M{"$year": "$created_at"},
				"month": bson.M{"$month": "$created_at"},
				"day":   bson.M{"$dayOfMonth": "$created_at"},
				"hour":  bson.M{"$hour": "$created_at"},
			},
			"date": bson.M{"$first": "$created_at"},
			"average_transfer_time_ms": bson.M{"$avg": "$transfer_time_ms"},
		}},
		{"$sort": bson.M{
			"date": 1,
		}},
	})
	return aggregation.All(result)
}

func GetOriginList() []string {
	aggregationResult := []bson.M{}
	result := []string{}
	aggregation := getPingCollection().Pipe([]bson.M{
		{"$project": bson.M{"origin": true}},
		{"$group": bson.M{
			"_id": "$origin",
		}},
	})

	_ = aggregation.All(&aggregationResult)

	for _, element := range aggregationResult {
		origin := element["_id"].(string)
		result = append(result, origin)
	}

	return result
}
