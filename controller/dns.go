package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/genofire/restrict-mini-dns-server/lib"
	"github.com/genofire/restrict-mini-dns-server/model"
)

type DNSController struct {
	session *lib.Session
}

func NewDNSController(session *lib.Session) *DNSController {
	return &DNSController{session: session}
}

func (c *DNSController) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var data []model.DNSEntry
	c.session.DB.Find(&data)
	lib.JSONOutput(w, r, data)
}

func (c *DNSController) GetOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dns := &model.DNSEntry{}
	id, err := strconv.ParseUint(ps.ByName("id"), 10, 64)
	if err != nil {
		dns.Subdomain = ps.ByName("id")
	} else {
		dns.Id = id
	}
	c.session.DB.Get(dns)
	lib.JSONOutput(w, r, dns)
}
func (c *DNSController) AddOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dns := &model.DNSEntry{Subdomain: ps.ByName("name")}
	// var dns model.DNSEntry
	// err := json.NewDecoder(r.Body).Decode(&dns)
	c.session.DB.Insert(dns)
	lib.JSONOutput(w, r, dns)
}
func (c *DNSController) SetOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.ParseUint(ps.ByName("id"), 10, 64)
	var dns model.DNSEntry
	c.session.DB.Id(id).Get(&dns)
	json.NewDecoder(r.Body).Decode(&dns)
	c.session.DB.Id(id).Update(&dns)
	lib.JSONOutput(w, r, dns)
}
