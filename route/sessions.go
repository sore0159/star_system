package route

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/sore0159/star_system/data"
)

type Logger interface {
	ServerErr(string, ...interface{})
	Record(string, ...interface{})
}

func CaptainRouter(
	l Logger,
	a data.Academy,
	s *securecookie.SecureCookie,
	found func(w http.ResponseWriter, r *http.Request, c *data.Captain),
	made func(w http.ResponseWriter, r *http.Request, c *data.Captain),
	crash func(w http.ResponseWriter, r *http.Request, msg string),
) (handler func(http.ResponseWriter, *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := CheckCookieUID(r, s)
		if err != nil {
			l.ServerErr("uid check failed: %v", err)
			crash(w, r, "UID CHECK FAILURE")
			return
		}

		var c *data.Captain
		if uid != 0 {
			l.Record("FOUND UID %v", uid)
			c, err = a.SearchCaptain(uid)
			if err == data.ERR_CAP404 {
				l.Record("cap404 for UID %v", uid)
				c, err = nil, nil
			} else if err != nil {
				l.ServerErr("captain search failed: %v", err)
				crash(w, r, "CAPTAIN SEARCH FAILURE")
				return
			}
		}
		if c == nil {
			c, err = a.NewCaptain()
			if err != nil {
				l.ServerErr("captain creation failed: %v", err)
				crash(w, r, "CAPTAIN CREATION FAILURE")
				return
			}
			l.Record("SET UID %v", c.UID)
			SetCookieUID(c.UID, w, s)
			made(w, r, c)
			return
		}
		found(w, r, c)
	}
}
