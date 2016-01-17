package controller

import (
	"github.com/emicklei/go-restful"
	"github.com/Nainterceptor/MiniProject-Ping/entity"
	"net/http"
	"gopkg.in/mgo.v2/bson"
)

func PingCreate(request *restful.Request, response *restful.Response) {
	ping := entity.PingNew()

	if err := request.ReadEntity(&ping); err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}

	if err := ping.Validate(); err != nil {
		response.WriteError(http.StatusNotAcceptable, err)
		return
	}

	if err := ping.Insert(); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	response.WriteEntity(ping)
}

func PingAverageTime(request *restful.Request, response *restful.Response) {
	origin := request.PathParameter("origin")
	result := []bson.M{}
	if err := entity.AggregatePingOrigin(origin, &result); err != nil {
		if err.Error() == "can't convert from BSON type EOO to Date" { //Fix Aggregation "not found" error
			response.WriteErrorString(http.StatusNotFound, "Not Found")
		}
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	response.WriteEntity(result)
}

func PingOriginList(request *restful.Request, response *restful.Response) {
	result := entity.GetOriginList();
	response.WriteAsJson(result);
}