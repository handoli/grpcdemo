package main

import (
	clientStream "grpcdemo/proto/client_stream"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

type MyClientStreamService struct{
	clientStream.UnimplementedClientStreamServiceServer
}

func (MyClientStreamService) ClientStream(s clientStream.ClientStreamService_ClientStreamServer) error {

	for {
		recv,err :=s.Recv()
		if err != nil {
			return err
		}
		msg := recv.Msg
		if strings.EqualFold(msg,"end") {
			break
		}
 
		if recv.Type == 1 {
			log.Println("门派是",msg)
		}else {
			log.Println("技能是",msg)
		}
	}
	err := s.SendAndClose(&clientStream.ClientStreamRes{Res: "接收完毕！！！"})
	if err != nil {
		return err
	}

	return nil
}


//创建grpc服务端，开始暴露grpc服务
func createServer(){
	//定义grpc服务使用的协议和端口
	listen, _ := net.Listen("tcp", ":9997")
	//通过grpc包创建Server
	s :=grpc.NewServer()
	//将重写了Add方法的结构体绑定到gRpc服务上
	clientStream.RegisterClientStreamServiceServer(s,&MyClientStreamService{})
	//启动服务等待客户端请求
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	createServer()
}