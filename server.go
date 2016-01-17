package main

import (
	"github.com/Nainterceptor/MiniProject-Ping/config"
	"github.com/Nainterceptor/MiniProject-Ping/controller"
	"github.com/Nainterceptor/MiniProject-Ping/entity"
	"github.com/emicklei/go-restful"
	"net/http"
	"strconv"
)

const (
	VERSION_API = 1
)

func main() {

	router := initRouter()

	server := &http.Server{Addr: config.HttpBinding, Handler: router}
	server.ListenAndServe()
}

func initRouter() *restful.Container {
	routerContainer := restful.NewContainer()
	ping := new(restful.WebService)

	ping.
		Path("/api/" + strconv.Itoa(VERSION_API) + "/pings").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ping.Route(ping.
		POST("").
		To(controller.PingCreate).
		Doc("Create a new ping").
		Operation("PingCreate").
		Returns(http.StatusOK, "Ping has been created", nil).
		Returns(http.StatusBadRequest, "Can't read entity", nil).
		Returns(http.StatusNotAcceptable, "Validation has failed", nil).
		Returns(http.StatusInternalServerError, "Return of MongoDB Insert", nil).
		Reads(entity.Ping{}))

	ping.Route(ping.
		GET("origins").
		To(controller.PingOriginList).
		Doc("List all ping origin").
		Operation("PingOriginList"))

	ping.Route(ping.
		GET("/{origin}/hours").
		To(controller.PingAverageTime).
		Doc("Retrieve the average transfer_time_ms for an origin").
		Operation("PingAverageTime").
		Param(ping.PathParameter("origin", "origin of the ping").DataType("string")).
		Returns(http.StatusOK, "Stats has been returned", nil).
		Returns(http.StatusNotFound, "Origin not found", nil))

	static := new(restful.WebService)

	static.Route(static.GET("/{subpath:*}").To(controller.ServeStatic))

	routerContainer.Add(ping)
	routerContainer.Add(static)

	return routerContainer
}
