package main

import (
	"context"
	"grpcdemo/proto/one_way"
	"google.golang.org/grpc"
)

//创建gprc客户端，请求grpc服务端
func createClient(){
	//定义处理gprc请求的服务端地址，可以设置一些连接参数
	conn,_ := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
	//创建gprc代理并
	c := oneWay.NewOneWayServiceClient(conn)
	//定义封装gprc客户端请求参数
	req := oneWay.OneWayReq{
		A: 19,
		B: 11,
	}
	//调用远程grpc服务，并取得响应结果
	reply,_ := c.OneWay(context.Background(), &req)
	//使用响应结果
	println(reply.Res)
}


func main(){
	createClient()
}