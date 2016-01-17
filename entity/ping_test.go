package entity

import (
	"io/ioutil"
	"testing"

	"github.com/Nainterceptor/MiniProject-Ping/config"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/dbtest"
)

var Server dbtest.DBServer
var Session *mgo.Session

func init() {

	tempDir, _ := ioutil.TempDir("", "testing")
	Server.SetPath(tempDir)

	Session = Server.Session()
	config.MongoDB = Session.DB("test_pings")
}

func Wipe() {
	config.MongoDB.DropDatabase()
}

func TestPingNew(t *testing.T) {
	ping := PingNew()
	Convey("Test ping informations", t, func() {
		So(ping.Id, ShouldNotBeEmpty)
		So(ping.CreatedAt, ShouldNotBeEmpty)
	})
}

func TestPingNormalize(t *testing.T) {
	Convey("Test ping normalize", t, func() {
		ping := PingNew()
		ping.Origin = " foo "
		ping.Normalize()
		Convey("ping origin must be trimed", func() {
			So(ping.Origin, ShouldEqual, "foo")
		})

	})
}

func TestPingValidation(t *testing.T) {
	Convey("Test ping validation", t, func() {
		ping := PingNew()

		Convey("Missing origin should back an error", func() {
			So(ping.Validate(), ShouldNotBeNil)
		})
		ping.Origin = "foo"
		Convey("Fullfilled origin should not trigger an error", func() {
			So(ping.Validate(), ShouldBeNil)
		})
	})
}

func TestPingInsert(t *testing.T) {
	Wipe()
	Convey("Test ping insertion", t, func() {
		ping := getFooPing()
		Convey("Ping should be inserted", func() {
			So(ping.Insert(), ShouldBeNil)
		})
	})
}

func getFooPing() *Ping {
	ping := PingNew()
	ping.Origin = "Foo"
	ping.ConnectTimeMs = 100
	ping.NameLookupTimeMs = 100
	ping.Status = 200
	ping.TotalTimeMs = 100
	ping.TransferTimeMs = 100
	return ping
}
