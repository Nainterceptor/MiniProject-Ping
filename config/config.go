package config

import (
	"flag"

	"github.com/vharitonsky/iniflags"
	"gopkg.in/mgo.v2"
)

var HttpBinding = flag.String("http_binding", "localhost:1337", "IP/Port to listen HTTP Server")
var mongoCS = flag.String("mongodb_CS", "localhost", "Connection endpoint for mongodb driver")
var mongoName = flag.String("mongodb_DB", "PingProject", "Database to mount")

var MongoDB *mgo.Database

func init() {
	iniflags.Parse()
	session, err := mgo.Dial(*mongoCS)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	MongoDB = session.DB(*mongoName)
}
