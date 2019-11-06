package zzblog

import (
	"testing"
	"log"
)

func TestZzblogCreateIfDataDirExist(t * testing.T) {
	zz := NewFileZzblog("./test/")
	set := zz.Filter(func (*Blog) bool {
		return true
	})
	for _, b := range set.Get() {
		log.Printf("%v\n", b)
	}
}

func TestZzblogCreateIfDataDirNotExist(t * testing.T) {
	NewFileZzblog("./not-exists/")
}