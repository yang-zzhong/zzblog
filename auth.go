package zzblog

import (
	hr "github.com/yang-zzhong/go-httprouter"
)

type Auth struct {
	passport string
}

func NewAuth(passport string) *Auth {
	return &Auth{passport}
}

func (auth *Auth) Before(w *hr.Response, req *hr.Request) bool {
	if req.Request.Header.Get("passport") == auth.passport {
		return true
	}
	w.WithStatus(403)
	return false
}

func (auth *Auth) After(_ *hr.Response, _ *hr.Request) bool {
	return true
}
