package zzblog

import (
	"log"
	"testing"
)

func TestTraversingDir(t *testing.T) {
	log.Printf("testing traversing\n")
	TraversingDir("./test/", func(pathfile string) {

	})
}
