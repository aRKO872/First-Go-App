package main

import (
	"log"
	"net"

	authcontrollers "github.com/aRKO872/first-go-app/controllers/authController"
	"github.com/aRKO872/first-go-app/controllers/dbcontroller"
	pb "github.com/aRKO872/first-go-app/proto"
	"google.golang.org/grpc"

	"github.com/aRKO872/first-go-app/controllers/todocontrollers"
)

type TodoServer struct {
	*todocontrollers.TodoServer
}

type AuthServer struct {
	*authcontrollers.AuthServer
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
	pb.RegisterAuthServiceServer(s, &AuthServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("grpc server failed: %v", err)
	}
}