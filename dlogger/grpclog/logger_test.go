package grpclog

import (
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
)

func TestLoggerTest(t *testing.T) {
	var logger = logging.NewLogger("grpcLogger")
	zlog := NewZapLogger(logger)
	zlog.Info("111")
}
