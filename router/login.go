package router

import (
	"github.com/julienschmidt/httprouter"

	"github.com/genofire/restrict-mini-dns-server/controller"
	"github.com/genofire/restrict-mini-dns-server/lib"
)

func SetLoginRoutes(r *httprouter.Router, session *lib.Session) {
	prefix := "/login"
	con := controller.NewLoginController(session)

	r.POST(prefix+"/signup", con.Signup)
	r.POST(prefix+"/login", con.Login)
	r.GET(prefix+"/current", con.Current)
}
