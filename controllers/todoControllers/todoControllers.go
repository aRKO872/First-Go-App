package todocontrollers

import (
	"context"
	"errors"
	"log"

	pb "github.com/aRKO872/first-go-app/proto"

	"github.com/aRKO872/first-go-app/controllers/dbcontroller"
	"github.com/google/uuid"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

func (s *TodoServer) FetchAllTodos(ctx context.Context, req *pb.EmptyRequest) (*pb.ArrayOfTodos, error) {
	var todos []*pb.Todos;
	result := dbcontroller.DB.Find(&todos);

	if result.Error != nil {
		return nil, errors.New ("error finding todos") 
	}

	return &pb.ArrayOfTodos{
		TodoArr: todos,
	}, nil
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

	result := dbcontroller.DB.Create(&todo)

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

func (s *TodoServer) EditSingleTodo(ctx context.Context, req *pb.EditOrDeleteRequest) (*pb.Todos, error) {
	todoId := req.Id;
	if todoId == "" {
		return nil, errors.New("Please Provide Todo - ID!")
	}
	todo := &pb.Todos {
		Id: todoId,
	}

	result := dbcontroller.DB.First (&todo);
	if result.Error != nil || todo == nil {
		return nil, errors.New ("Todo Does not Exist!")
	}

	if todo.Done {
		result := dbcontroller.DB.Model (&todo).Where ("id = ?", todoId).Update ("done", false)
		if result.Error != nil {
			return nil, errors.New ("Error Updating Todo!")
		}
	} else {
		result := dbcontroller.DB.Model (&todo).Where ("id = ?", todoId).Update ("done", true)
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

	result := dbcontroller.DB.First (&todo);
	if result.Error != nil || todo == nil {
		return nil, errors.New ("Todo Does not Exist!")
	}

	dbcontroller.DB.Where("id = ?", todoId).Delete(&todo)

	return todo, nil
}