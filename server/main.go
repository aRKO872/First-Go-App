package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"

	pb "github.com/aRKO872/first-go-app/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TodoServer struct {
 	pb.UnimplementedTodoServiceServer
}

func init() {
	DatabaseConnection()
 }

 var DB *gorm.DB
 var err error
 
 
 func DatabaseConnection() {
	log.Println(os.Getenv("DATABASE_URL"))
	// DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	DB, err = gorm.Open(postgres.Open("postgres://postgres@localhost:5432/todo-go-db?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error: ", err)
	}
	log.Println("db connection successful")
 }


func (s *TodoServer) CreateTodo(ctx context.Context, req *pb.NewTodo) (*pb.Todos, error) {
	log.Printf("received: %v", req.GetName())
	
	todo := &pb.Todos{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Id:					 uuid.NewString(),
	}

	result := DB.Create(&todo)

	if result.RowsAffected == 0 {
		return nil, errors.New("error saving todo")
	}

	return &pb.Todos{
		Name: todo.Name,
		Description: todo.Description,
		Id: todo.Id,
		Done: false,
	}, nil
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