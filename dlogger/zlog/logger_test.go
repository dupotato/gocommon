package zlog

import (
	"testing"
)

func TestLog1(t *testing.T) {
	Infof("111")
	Debugf("2222%s", "2")
}
