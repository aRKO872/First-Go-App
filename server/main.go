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
	
	if req.GetName() == "" || req.GetDescription() == "" {
		return nil, errors.New("Please provide Name & Description!");
	}

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

func (s *TodoServer) FetchAllTodos(ctx context.Context, req *pb.EmptyRequest) (*pb.ArrayOfTodos, error) {
	var todos []*pb.Todos;
	result := DB.Find(&todos);

	if result.Error != nil {
		return nil, errors.New ("error finding todos") 
	}

	return &pb.ArrayOfTodos{
		TodoArr: todos,
	}, nil
} 

func (s *TodoServer) EditSingleTodo(ctx context.Context, req *pb.EditOrDeleteRequest) (*pb.Todos, error) {
	todoId := req.Id;
	if todoId == "" {
		return nil, errors.New("Please Provide Todo - ID!")
	}
	todo := &pb.Todos {
		Id: todoId,
	}

	result := DB.First (&todo);
	if result.Error != nil || todo == nil {
		return nil, errors.New ("Todo Does not Exist!")
	}

	if todo.Done {
		result := DB.Model (&todo).Where ("id = ?", todoId).Update ("done", false)
		if result.Error != nil {
			return nil, errors.New ("Error Updating Todo!")
		}
	} else {
		result := DB.Model (&todo).Where ("id = ?", todoId).Update ("done", true)
		if result.Error != nil {
			return nil, errors.New ("Error Updating Todo!")
		}
	}

	return todo, nil
}

func (s *TodoServer) DeleteSingleTodo(ctx context.Context, req *pb.EditOrDeleteRequest) (*pb.Todos, error) {
	todoId := req.GetId();
	if todoId == "" {
		return nil, errors.New("Please Provide Todo - ID!")
	}
	todo := &pb.Todos {
		Id: todoId,
	}

	result := DB.First (&todo);
	if result.Error != nil || todo == nil {
		return nil, errors.New ("Todo Does not Exist!")
	}

	DB.Where("id = ?", todoId).Delete(&todo)

	return todo, nil
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