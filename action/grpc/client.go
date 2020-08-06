package grpcserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"hcc/clarinet/action/grpc/rpcflute"
)

const (
	address = "localhost:50051"
)

type server struct {
	rpcflute.UnimplementedFluteServer
}

func (s *server) OnOffNode(ctx context.Context, in *rpcflute.ReqOnOffNode) (*rpcflute.ResOnOffNode, error) {
	switch in.PowerState {
	case rpcflute.ReqOnOffNode_ON:
		for _, node := range in.Nodes {
			args, err := json.Marshal(node)
			if err != nil {
				return &rpcflute.ResOnOffNode{Nodes: nil}, err
			}
			fmt.Print("call dao.NodeON with args: ")
			fmt.Println(node)
			fmt.Println(args)
		}
	case rpcflute.ReqOnOffNode_OFF:
		fmt.Println("call dao.NodeOFF")
	case rpcflute.ReqOnOffNode_RESTART:
		fmt.Println("call dao.NodeRESTART")
	}
	return &rpcflute.ResOnOffNode{Nodes: in.Nodes}, nil
}

func InitGRPC() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("faile listen tcp")
	}
	s := grpc.NewServer()

	rpcflute.RegisterFluteServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return err
}
