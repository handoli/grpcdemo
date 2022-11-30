package main

import (
	serverStream "grpcdemo/proto/server_stream"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

// const (
// 	MP = []string{
// 		"大唐官府","化生寺","方寸山","女儿村",
// 		"狮驼岭","魔王寨","阴曹地府",
// 		"龙宫","普陀山","月宫","花果山",
// 	}

// 	JN = []string {
// 		"横扫千军","推气过功","失心符","一笑倾城",
// 		"鹰击","飞沙走石","尸腐毒",
// 		"龙卷雨击","普度众生","月光摇影","当头一棒",
// 	}
// )


type MyServerStreamService struct{
	serverStream.UnimplementedServerStreamServiceServer
}

func (MyServerStreamService) ServerStream(req *serverStream.ServerStreamReq, s serverStream.ServerStreamService_ServerStreamServer) error {
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

	reqType := req.Type
	number := req.Number
	if reqType == 1 && number > 0{
		//门派
		var res string = "服务端共向客户端发送"+ strconv.Itoa(int(number)) +"个门派"
		s.Send(&serverStream.ServerStreamRes{Res : res})
		for i := 0; i <= int(number); i++{
			res = mp[i];
			s.Send(&serverStream.ServerStreamRes{Res : res})
		}
	}else if reqType == 2 && number > 0{
		//技能
		var res string = "服务端共向客户端发送"+strconv.Itoa(int(number))+"个技能"
		s.Send(&serverStream.ServerStreamRes{Res : res})
		for i := 0; i <= int(number); i++{
			res = jn[i];
			s.Send(&serverStream.ServerStreamRes{Res : res})
		}
	}else {
		//没有找到
		var res string = "服务端没有收到有效数据"
		s.Send(&serverStream.ServerStreamRes{Res : res})
	}
	return nil
}

//创建grpc服务端，开始暴露grpc服务
func createServer(){
	//定义grpc服务使用的协议和端口
	listen, _ := net.Listen("tcp", ":9998")
	//通过grpc包创建Server
	s :=grpc.NewServer()
	//将重写了Add方法的结构体绑定到gRpc服务上
	serverStream.RegisterServerStreamServiceServer(s,&MyServerStreamService{})
	//启动服务等待客户端请求
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	createServer()
}