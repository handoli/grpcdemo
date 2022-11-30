package main

import (
	"context"
	"grpcdemo/proto/two_way"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//创建gprc客户端，请求grpc服务端
func createClient(){
	//定义处理gprc请求的服务端地址，可以设置一些连接参数
	conn,err := grpc.Dial("127.0.0.1:9995", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
        log.Fatalf("net.Connect err: %v", err)
    }
	
	defer conn.Close()
	//创建gprc代理并
	c := twoWay.NewTwoWayServiceClient(conn)
	res,err := c.TwoWay(context.Background())
	if err != nil {
        log.Fatalf("Call SayHello err: %v", err)
    }
	a := int64(1)
	b := int64(1)
	count := 0
	out: for{
		res.Send(&twoWay.TwoWayReq{A: a,B: b})
		r,_ := res.Recv()
		for i := 0; i < int(b); i++ {
			if r.Res == 0 {
				log.Println("服务端暂停处理")
				break out
			}else {
				log.Println("接收服务端数据："+r.Msg)
			}
		}
		count++
		a = a+int64(count)
		a = a % 2
	}
}

func main(){
	createClient()
}