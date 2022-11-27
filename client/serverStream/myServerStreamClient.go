package main

import (
	"context"
	serverStream "grpcdemo/proto/server_stream"
	"io"
	"log"

	"google.golang.org/grpc"
)

//创建gprc客户端，请求grpc服务端
func createClient(){
	//定义处理gprc请求的服务端地址，可以设置一些连接参数
	conn,_ := grpc.Dial("127.0.0.1:9998", grpc.WithInsecure())
	
	defer conn.Close()
	//创建gprc代理并
	c := serverStream.NewServerStreamServiceClient(conn)
	//定义封装gprc客户端请求参数
	req := serverStream.ServerStreamReq{
		Type: 5,
		Number: 3,
	}
	//调用远程grpc服务，并取得响应结果
	stream,_ := c.ServerStream(context.Background(), &req)
	//使用响应结果
	for {
		res,err := stream.Recv()
		if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("Conversations get stream err: %v", err)
        }
        // 打印返回值
        log.Println(res.Res)
	}
	
}


func main(){
	createClient()
}