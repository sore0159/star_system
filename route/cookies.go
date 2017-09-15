package route

import (
	"encoding/gob"
	"math/big"
	"net/http"

	"github.com/gorilla/securecookie"
)

func init() {
	gob.Register(&big.Int{})
}

func NewCookieSecurity(hash, block []byte) *securecookie.SecureCookie {
	return securecookie.New(hash, block)
}

func CheckCookieUID(r *http.Request, s *securecookie.SecureCookie) (*big.Int, error) {
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

func SetCookieUID(uid *big.Int, w http.ResponseWriter, s *securecookie.SecureCookie) error {
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
