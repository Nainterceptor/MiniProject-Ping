package controller

import (
	"github.com/emicklei/go-restful"
	"github.com/Nainterceptor/MiniProject-Ping/entity"
	"net/http"
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
}
