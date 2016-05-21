package controller

import (
	"encoding/json"
	"net/http"
	"net/url"

	password "dev.sum7.de/sum7/sum7_api/lib_password"
	"github.com/julienschmidt/httprouter"

	"github.com/genofire/restrict-mini-dns-server/lib"
	"github.com/genofire/restrict-mini-dns-server/model"
)

const (
	maxlifetime = 3600
)

type LoginController struct {
	session *lib.Session
}

func NewLoginController(session *lib.Session) *LoginController {
	return &LoginController{session: session}
}

func (c *LoginController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var login model.Login
	json.NewDecoder(r.Body).Decode(&login)
	cookie, err := r.Cookie(lib.SessionID)
	var sid string
	if err != nil || cookie.Value == "" {
		sid = lib.GenerateSessionID()
		c.session.Logins[sid] = &model.Login{
			Username: login.Username,
		}
		cookie := http.Cookie{Name: lib.SessionID, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ = url.QueryUnescape(cookie.Value)
	}
	lib.JSONOutput(w, r, nil)
}
func (c *LoginController) Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var login model.Login
	json.NewDecoder(r.Body).Decode(&login)
	login.Password = password.NewHesh(login.Password)
	// c.session.DB.Get()
	lib.JSONOutput(w, r, nil)
}
func (c *LoginController) Current(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	login := c.session.GetLogin(r)
	lib.JSONOutput(w, r, login)
}
