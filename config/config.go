package config

import (
	"flag"

	"github.com/vharitonsky/iniflags"
	"gopkg.in/mgo.v2"
)

var (
	StaticPath  string
	HttpBinding string
	mongoCS     string
	mongoName   string
	MongoDB     *mgo.Database
)

func init() {
	flags()
	iniflags.Parse()
	session, err := mgo.Dial(mongoCS)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	MongoDB = session.DB(mongoName)
}

func flags() {
	flag.StringVar(&StaticPath, "static_path", "static", "Localisation for static files")
	flag.StringVar(&HttpBinding, "http_binding", "localhost:1337", "IP/Port to listen HTTP Server")
	flag.StringVar(&mongoCS, "mongodb_CS", "localhost", "Connection endpoint for mongodb driver")
	flag.StringVar(&mongoName, "mongodb_DB", "PingProject", "Database to mount")
}
