package router

import (
	"github.com/julienschmidt/httprouter"

	"github.com/genofire/restrict-mini-dns-server/controller"
	"github.com/genofire/restrict-mini-dns-server/lib"
)

func SetDNSRoutes(_ *httprouter.Router, session *lib.Session) {
	// prefix := "/dns"
	controller.NewDNSController(session)
}
