package captain

import (
	"encoding/gob"
	"github.com/gorilla/securecookie"
	"math/big"
	"net/http"
)

func init() {
	gob.Register(&big.Int{})
}

func NewCookieSecurity(hash, block []byte) *securecookie.SecureCookie {
	return securecookie.New(hash, block)
}

func CheckUID(r *http.Request, s *securecookie.SecureCookie) (*big.Int, error) {
	cookie, err := r.Cookie("star-captain")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, nil
		}
		return nil, err
	}
	uid := new(big.Int)
	err = s.Decode("star-captain", cookie.Value, uid)
	if err != nil {
		return nil, err
	}
	return uid, nil
}

func SetUID(uid *big.Int, w http.ResponseWriter, s *securecookie.SecureCookie) error {
	encoded, err := s.Encode("star-captain", uid)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:     "star-captain",
		Value:    encoded,
		Path:     "/",
		MaxAge:   30000000,
		HttpOnly: true,
		// Secure: true,
	}
	http.SetCookie(w, cookie)
	return nil
}

type Logger interface {
	ServerErr(string, ...interface{})
	Record(string, ...interface{})
}

func CaptainRouter(
	l Logger,
	a Academy,
	s *securecookie.SecureCookie,
	found func(w http.ResponseWriter, r *http.Request, c *Captain),
	made func(w http.ResponseWriter, r *http.Request, c *Captain),
	crash func(w http.ResponseWriter, r *http.Request, msg string),
) (handler func(http.ResponseWriter, *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := CheckUID(r, s)
		if err != nil {
			l.ServerErr("uid check failed: %v", err)
			crash(w, r, "UID CHECK FAILURE")
			return
		}
		var c *Captain
		if uid != nil {
			c, err = a.SearchCaptain(uid)
			if err == ERR_CAP404 {
				l.Record("cap404 for UID %s", uid)
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
			SetUID(&c.UID, w, s)
			made(w, r, c)
			return
		}
		found(w, r, c)
	}
}
