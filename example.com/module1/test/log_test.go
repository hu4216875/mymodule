package test

import (
	"testing"

	"example.com/module1/log"
)

func init() {
	log.InitLog()
}

func TestDebug(t *testing.T) {
	log.Debugf("hello, world")
}
