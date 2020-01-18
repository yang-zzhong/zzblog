package zzblog

import (
	httprouter "github.com/yang-zzhong/go-httprouter"
)

type Auth struct {
	passport string
}

func NewAuth(passport string) *Auth {
	return &Auth{passport}
}

func (auth *Auth) Before(w *httprouter.ResponseWriter, req *httprouter.Request) bool {
	if req.Request.Header.Get("passport") == auth.passport {
		return true
	}
	w.WithStatusCode(403)
	return false
}

func (auth *Auth) After(_ *httprouter.ResponseWriter, _ *httprouter.Request) bool {
	return true
}
