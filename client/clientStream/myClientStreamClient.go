package main

import (
	"context"
	clientStream "grpcdemo/proto/client_stream"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//创建gprc客户端，请求grpc服务端
func createClient(){
	//定义处理gprc请求的服务端地址，可以设置一些连接参数
	conn,err := grpc.Dial("127.0.0.1:9997", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
        log.Fatalf("net.Connect err: %v", err)
    }
	
	defer conn.Close()
	//创建gprc代理并
	c := clientStream.NewClientStreamServiceClient(conn)
	res,err := c.ClientStream(context.Background())
	if err != nil {
        log.Fatalf("Call SayHello err: %v", err)
    }

	mp := []string{
		"大唐官府","化生寺","方寸山","女儿村",
		"狮驼岭","魔王寨","阴曹地府",
		"龙宫","普陀山","月宫","花果山",
	}

	jn := []string {
		"横扫千军","推气过功","失心符","一笑倾城",
		"鹰击","飞沙走石","尸腐毒",
		"龙卷雨击","普度众生","月光摇影","当头一棒",
	}

	for i := 0; i < 10; i++ {
		if i < 5 {
			_ = res.Send(&clientStream.ClientStreamReq{Type: 1,Msg: mp[i]})
		}else {
			_ = res.Send(&clientStream.ClientStreamReq{Type: 2,Msg: jn[i]})
		}
	}

	res.Send(&clientStream.ClientStreamReq{Type: 1,Msg: "end"})

	recv,_ := res.CloseAndRecv()
	
	log.Println("服务端接收状态:"+recv.Res)
}


func main(){
	createClient()
}