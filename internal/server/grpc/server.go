package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	srv *grpc.Server
}

func NewServer() *Server{
	return &Server{
		srv: grpc.NewServer(),
	}
}

func (s *Server) ListenAndServe(port int) error{
	addr:= fmt.Sprintf(":%d", port)

	lis, err:= net.Listen("tcp", addr)
	if err != nil{
		log.Fatalf("Failed to listen on port %s : %v", addr, err)
		return err
	}
	
	if err:= s.srv.Serve(lis); err != nil{
		log.Fatalf("Failed to serve gRPC server over port %s : %v", addr, err)
	}

	return nil
}