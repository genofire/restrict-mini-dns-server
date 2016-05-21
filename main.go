package main

import (
	"flag"

	"log"
	"net"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/go-xorm/xorm"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"

	"github.com/genofire/restrict-mini-dns-server/lib"
	"github.com/genofire/restrict-mini-dns-server/model"
	"github.com/genofire/restrict-mini-dns-server/router"
)

var (
	configFile string
	config     *model.Config
)

func main() {
	flag.StringVar(&configFile, "config", "config.yml", "path of configuration file (default:config.yaml)")
	flag.Parse()

	config = model.ReadConfigFile(configFile)

	log.Print("connecting database")
	engine, _ := xorm.NewEngine("sqlite3", config.Database.Path)
	engine.Sync(&model.DNSEntry{}, &model.Login{})
	log.Print("connected  database")

	r := httprouter.New()
	if config.Webserver.Api {
		log.Print("enable api server")
		session := lib.NewSession(engine)

		router.SetLoginRoutes(r, session)
		router.SetDNSRoutes(r, session)
	}
	if config.Webserver.Enable {
		log.Print("enable web server")
		r.NotFound = gziphandler.GzipHandler(http.FileServer(http.Dir(config.Webserver.Webroot)))
	}

	if config.Webserver.Api || config.Webserver.Enable {
		address := net.JoinHostPort(config.Webserver.Address, config.Webserver.Port)
		log.Print("listen server on ", address)
		log.Fatal(http.ListenAndServe(address, r))
	}
}
