package main

import (
	"context"
	//引用.proto生成的包，grpcdemo/proto/single : go.mod(module)/.proto代码所在的目录名
	"grpcdemo/proto/single" 
	"log"
	"net"
	"google.golang.org/grpc"
)

//自己定一个结构体，引用UnimplementedAddServiceServer，来实现.proto中定义的的service
type MySingleService struct{
	/*
		add : 与引用的包不同，引用包使用的.proto代码所在的目录名称，这里的add是.proto代码文件的包名称
		UnimplementedAddServiceServer : 见.proto描述
	*/
	single.UnimplementedSingleServiceServer
}

//重写service接口的暴露方法，注意重写的是实现了service接口的UnimplementedAddServiceServer的Add方法
func (MySingleService) Single(content context.Context, req *single.SingleReq) (*single.SingleRes, error){
	log.Println("参数a的值为：",req.A)
	log.Println("参数b的值为：",req.B)
	return &single.SingleRes{Res : req.A + req.B},nil
}
	
//创建grpc服务端，开始暴露grpc服务
func createServer(){
	//定义grpc服务使用的协议和端口
	lis, _ := net.Listen("tcp", ":9999")
	//通过grpc包创建Server
	s :=grpc.NewServer()
	//将重写了Add方法的结构体绑定到gRpc服务上
	single.RegisterSingleServiceServer(s,&MySingleService{})
	//启动服务等待客户端请求
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


func main() {
	//创建和启动grpc服务端
	createServer()
}