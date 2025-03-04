package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"CookingMaster_Backend/app/recipe/api/internal/config"
	"CookingMaster_Backend/app/recipe/api/internal/handler"
	"CookingMaster_Backend/app/recipe/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/recipe-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	domains := strings.Split(c.FrontendDomains, ",")
	server := rest.MustNewServer(c.RestConf, rest.WithCors(domains...), rest.WithCustomCors(func(header http.Header) {
		header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
		header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		header.Add("Access-Control-Allow-Headers", "Content-Length, Content-Type")
	}, func(writer http.ResponseWriter) {
		writer.WriteHeader(http.StatusForbidden)
	}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
