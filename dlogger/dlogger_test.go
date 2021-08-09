package dlogger

import "testing"

func TestLog(t *testing.T) {
	Init("grpc")
	Debugf("%s", "122")
}
