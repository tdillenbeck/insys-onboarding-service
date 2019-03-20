package wiggum

/*
	methods for setting/deleteing token as cookie
*/

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SetCookie(token string, w http.ResponseWriter, r *http.Request) error {
	return Default.SetCookie(token, w, r)
}

func (k *KeySet) SetCookie(token string, w http.ResponseWriter, r *http.Request) error {
	cookieName := k.CookieName()

	tokenCookie, err := r.Cookie(cookieName)
	if err != nil && err != http.ErrNoCookie {
		return fmt.Errorf("unable to get cookie: %s", err)
	}

	if err == http.ErrNoCookie {
		tokenCookie = &http.Cookie{Name: cookieName}
	}

	tokenCookie.Value = token
	tokenCookie.Path = "/"
	tokenCookie.Expires = time.Now().Add(time.Hour * 4) // set cookie to expire in 4 hours

	http.SetCookie(w, tokenCookie)
	return nil
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	Default.DeleteCookie(w, r)
}

func (k *KeySet) DeleteCookie(w http.ResponseWriter, r *http.Request) {
	cookieName := k.CookieName()

	tokenCookie, err := r.Cookie(cookieName)
	if err != nil {
		return
	}
	tokenCookie.Path = "/"
	tokenCookie.Value = ""
	tokenCookie.Expires = time.Now()
	http.SetCookie(w, tokenCookie)
	return
}

/* --- Private --- */
func getCookie(r *http.Request, cookieName string) (string, bool) {
	tokenCookie, err := r.Cookie(cookieName)
	switch {
	case err == http.ErrNoCookie:
		return "No cookie", false
	case err != nil:
		return "Error parsing cookie", false
	}
	val := tokenCookie.Value
	ok := present(val)
	return val, ok
}

func getHeader(r *http.Request) (string, bool) {
	authHeader, ok := r.Header["Authorization"]
	if ok == false || authHeader == nil {
		return "No Authorization header", false
	}
	token := ""
	for _, header := range authHeader {
		headerArr := strings.Split(header, " ")
		if headerArr[0] == "Bearer" {
			if len(headerArr) > 1 {
				token = headerArr[1]
				break
			}
		}
	}

	if token == "" {
		return "No Bearer token", false
	}

	return token, true
}

func getQuery(f url.Values) (string, bool) {
	tokenParam := f.Get("token")
	if tokenParam == "" {
		return "No token query param", false
	}

	return tokenParam, true
}

func present(val string) bool {
	if val == "" {
		return false
	} else {
		return true
	}
}
