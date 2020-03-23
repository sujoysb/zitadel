package http

import (
	"net/http"

	"github.com/gorilla/securecookie"

	"github.com/caos/zitadel/internal/errors"
)

type CookieHandler struct {
	securecookie *securecookie.SecureCookie
	secureOnly   bool
	sameSite     http.SameSite
	path         string
	maxAge       int
	domain       string
}

func NewCookieHandler(opts ...CookieHandlerOpt) *CookieHandler {
	c := &CookieHandler{
		secureOnly: true,
		sameSite:   http.SameSiteLaxMode,
		path:       "/",
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

type CookieHandlerOpt func(*CookieHandler)

func WithEncryption(hashKey, encryptKey []byte) CookieHandlerOpt {
	return func(c *CookieHandler) {
		c.securecookie = securecookie.New(hashKey, encryptKey)
	}
}

func WithUnsecure() CookieHandlerOpt {
	return func(c *CookieHandler) {
		c.secureOnly = false
	}
}

func WithSameSite(sameSite http.SameSite) CookieHandlerOpt {
	return func(c *CookieHandler) {
		c.sameSite = sameSite
	}
}

func WithPath(path string) CookieHandlerOpt {
	return func(c *CookieHandler) {
		c.path = path
	}
}

func WithMaxAge(maxAge int) CookieHandlerOpt {
	return func(c *CookieHandler) {
		c.maxAge = maxAge
		c.securecookie.MaxAge(maxAge)
	}
}

func WithDomain(domain string) CookieHandlerOpt {
	return func(c *CookieHandler) {
		c.domain = domain
	}
}

func (c *CookieHandler) GetCookieValue(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func (c *CookieHandler) GetEncryptedCookieValue(r *http.Request, name string, value interface{}) error {
	cookie, err := r.Cookie(name)
	if err != nil {
		return err
	}
	if c.securecookie == nil {
		return errors.ThrowInternal(nil, "HTTP-X6XpnL", "securecookie not configured")
	}
	if err := c.securecookie.Decode(name, cookie.Value, value); err != nil {
		return err
	}
	return nil
}

func (c *CookieHandler) SetCookie(w http.ResponseWriter, name string, value string) {
	c.httpSet(w, name, value, c.maxAge)
}

func (c *CookieHandler) SetEncryptedCookie(w http.ResponseWriter, name string, value interface{}) error {
	if c.securecookie == nil {
		return errors.ThrowInternal(nil, "HTTP-s2HUtx", "securecookie not configured")
	}
	encoded, err := c.securecookie.Encode(name, value)
	if err != nil {
		return err
	}
	c.httpSet(w, name, encoded, c.maxAge)
	return nil
}

func (c *CookieHandler) DeleteCookie(w http.ResponseWriter, name string) {
	c.httpSet(w, name, "", -1)
}

func (c *CookieHandler) httpSet(w http.ResponseWriter, name, value string, maxage int) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Domain:   c.domain,
		Path:     c.path,
		MaxAge:   maxage,
		HttpOnly: true,
		Secure:   c.secureOnly,
		SameSite: c.sameSite,
	})
}
