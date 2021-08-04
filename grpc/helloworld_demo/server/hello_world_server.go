package server

import (
	"fmt"
	golog "gocommon/dlogger/grpclog"
	"gocommon/grpc/helloworld_demo/impl"
	pb "gocommon/grpc/helloworld_demo/proto"
	"log"
	"net"
	"net/http"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// http 127.0.0.1:50051/debug/request
func startTrace() {

	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {

		return true, true

	}

	go http.ListenAndServe(":50051", nil)

	grpclog.Info("Trace listen on 50051")

}

func StartServer() {
	fmt.Println("begin")
	grpc.EnableTracing = true
	go startTrace()
	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var logger = logging.NewLogger("grpcLogger")
	z := golog.NewZapLogger(logger)
	z.Info("grpc server begin")
	grpclog.SetLoggerV2(z)
	//创建一个grpc服务器对象
	gRpcServer := grpc.NewServer()

	pb.RegisterHelloServiceServer(gRpcServer, &impl.HelloServiceServer{})
	//开启服务端
	gRpcServer.Serve(lis)
}
