package controller

import (
	"fmt"
	"github.com/Nainterceptor/MiniProject-Ping/config"
	"github.com/emicklei/go-restful"
	"net/http"
	"os"
	"path"
)

func ServeStatic(req *restful.Request, resp *restful.Response) {
	actual := path.Join(config.StaticPath, req.PathParameter("subpath"))
	if _, err := os.Stat(actual); os.IsNotExist(err) {
		actual = path.Join(config.StaticPath, "index.html")
	}
	fmt.Printf("serving %s ... (from %s)\n", actual, req.PathParameter("subpath"))
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		actual)
}
