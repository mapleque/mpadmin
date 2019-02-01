package mpadmin

import (
	"github.com/mapleque/kelp/http"
	"strconv"
)

func (this *Server) AuthToken(c *http.Context) *http.Status {
	if token := c.Request.Header.Get("Authoration"); token != this.token {
		return http.STATUS_FORBIDDEN
	}

	c.Next()
	return nil
}

func (this *Server) AuthUser(c *http.Context) *http.Status {
	idstr := c.Request.Header.Get("User-Id")
	if idstr == "" {
		return http.STATUS_UNAUTHORIZED
	}
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return http.STATUS_UNAUTHORIZED
	}
	c.Set("user_id", id)
	c.Next()
	return nil
}
