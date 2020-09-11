package zzblog

import (
	hr "github.com/yang-zzhong/go-httprouter"
)

type acrossDomain struct{}

func (ad *acrossDomain) Before(_ *hr.Response, _ *hr.Request) bool {
	return true
}

func (ad *acrossDomain) After(w *hr.Response, _ *hr.Request) bool {
	w.WithHeader("Access-Control-Allow-Origin", "*")
	w.WithHeader("Access-Control-Allow-Headers", "id")
	w.WithHeader("Access-Control-Allow-Methods", "*")
	return true
}

var AcrossDomain acrossDomain
