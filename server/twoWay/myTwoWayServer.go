package main

import (
	"grpcdemo/proto/two_way" 
	"log"
	"net"
	"google.golang.org/grpc"
)

type MyTwoWayService struct{
	twoWay.UnimplementedTwoWayServiceServer;
}

func (MyTwoWayService) TwoWay(req twoWay.TwoWayService_TwoWayServer) error {
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
	count := 0
	for{
		r,_ := req.Recv()
		a := r.A
		b := r.B
		if a == 1 {
			for i := 0; i < int(b); i++ {
				req.Send(&twoWay.TwoWayRes{Res : 1,Msg: mp[i]})
			}
		}else {
			for i := 0; i < int(b); i++ {
				req.Send(&twoWay.TwoWayRes{Res : 1,Msg: jn[i]})
			}
		}
		count++
		if count == 5 {
			req.Send(&twoWay.TwoWayRes{Res : 0,Msg: ""})
		}
	}
	return nil
}
	
func createServer(){
	//定义grpc服务使用的协议和端口
	lis, _ := net.Listen("tcp", ":9995")
	//通过grpc包创建Server
	s :=grpc.NewServer()
	//将重写了Add方法的结构体绑定到gRpc服务上
	twoWay.RegisterTwoWayServiceServer(s,&MyTwoWayService{})

	//启动服务等待客户端请求
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


func main() {
	//创建和启动grpc服务端
	createServer()
}