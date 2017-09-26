package route

import (
	"encoding/gob"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/sore0159/star_system/data"
)

func init() {
	gob.Register(new(data.UID))
}

func NewCookieSecurity(hash, block []byte) *securecookie.SecureCookie {
	return securecookie.New(hash, block)
}

func CheckCookieUID(r *http.Request, s *securecookie.SecureCookie) (data.UID, error) {
	cookie, err := r.Cookie("star-captain")
	if err != nil {
		if err == http.ErrNoCookie {
			return 0, nil
		}
		return 0, err
	}
	var uid data.UID
	err = s.Decode("star-captain", cookie.Value, &uid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

func SetCookieUID(uid data.UID, w http.ResponseWriter, s *securecookie.SecureCookie) error {
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
