package main

import (
	"log"
	"net"

	pb "github.com/aRKO872/first-go-app/proto"
	"github.com/aRKO872/first-go-app/controllers/dbcontroller"	
	"google.golang.org/grpc"
	
	"github.com/aRKO872/first-go-app/controllers/todocontrollers"
)

type TodoServer struct {
	*todocontrollers.TodoServer
}

func init() {
	dbcontroller.DatabaseConnection()
}

func main() {
	lis, err := net.Listen("tcp", ":6090")
	if err != nil {
		log.Fatalf("tcp connection failed: %v", err)
	}
	log.Printf("listening at %v", lis.Addr())

	s := grpc.NewServer()

	pb.RegisterTodoServiceServer(s, &TodoServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("grpc server failed: %v", err)
	}
}