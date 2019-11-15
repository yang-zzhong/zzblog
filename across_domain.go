package zzblog

import (
	httprouter "github.com/yang-zzhong/go-httprouter"
)

type acrossDomain struct{}

func (ad *acrossDomain) Before(_ *httprouter.ResponseWriter, _ *httprouter.Request) bool {
	return true
}

func (ad *acrossDomain) After(w *httprouter.ResponseWriter, _ *httprouter.Request) bool {
	w.WithHeader("Access-Control-Allow-Origin", "*")
	w.WithHeader("Access-Control-Allow-Headers", "id")
	w.WithHeader("Access-Control-Allow-Methods", "*")
	return true
}

var AcrossDomain acrossDomain
