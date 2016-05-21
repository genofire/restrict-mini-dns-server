package lib

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/go-xorm/xorm"

	"github.com/genofire/restrict-mini-dns-server/model"
)

const SessionID = "sessionid"

func GenerateSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

type Session struct {
	DB     *xorm.Engine
	Logins map[string]*model.Login
}

func NewSession(db *xorm.Engine) *Session {
	return &Session{
		DB:     db,
		Logins: make(map[string]*model.Login),
	}
}
func (s *Session) GetLogin(r *http.Request) *model.Login {
	cookie, err := r.Cookie(SessionID)
	if err == nil || cookie == nil || cookie.Value == "" {
		return nil
	}
	if login, _ := s.Logins[cookie.Value]; login != nil {
		return login
	} else {
		return nil
	}
}
