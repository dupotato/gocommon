package client

import (
	"context"
	"fmt"
	"go_common/golog"
	pb "go_common/grpc/helloworld_demo/proto"
	"log"
	"strconv"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func StartClient() {

	conn, err := grpc.Dial("127.0.0.1:8090", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建客户端存根对象
	var logger = logging.NewLogger("grpcLogger")
	z := golog.NewZapLogger(logger)
	z.Info("grpc server begin")
	grpclog.SetLoggerV2(z)
	c := pb.NewHelloServiceClient(conn)

	res, err := c.HelloWorld(context.Background(), new(pb.HelloRequest), grpc.EmptyCallOption{})

	fmt.Println(res, err)

	//客户端发起调用，返回一个流
	r, err := c.HelloWorldClientAndServerStream(context.Background(), grpc.EmptyCallOption{})
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	//用流给服务端发送消息
	for i := 0; i < 10; i++ {
		r.Send(&pb.HelloRequest{Request: "my is golang gRpc client " + strconv.Itoa(i)})
	}
	//流关闭
	r.CloseSend()
	for {
		res, err := r.Recv()
		if err != nil && err.Error() == "EOF" {
			break
		}
		if err != nil {
			log.Fatalf("%v", err)
			break
		}
		log.Printf("result:%v", res.Response)
	}
	defer conn.Close()
}
