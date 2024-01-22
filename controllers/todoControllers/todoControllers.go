package todocontrollers

import (
	"context"
	"errors"

	verify_auth "github.com/aRKO872/first-go-app/middlewares"
	pb "github.com/aRKO872/first-go-app/proto"

	"github.com/aRKO872/first-go-app/controllers/dbcontroller"
	"github.com/google/uuid"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

func (s *TodoServer) FetchAllTodos(ctx context.Context, req *pb.EmptyRequest) (*pb.ArrayOfTodos, error) {
	tokenPayload, tokenErr := verify_auth.VerifyAuthToken(ctx);

	if tokenErr != nil {
		return nil, tokenErr
	}

	logged_in_user_id := tokenPayload.UserId;

	var todos []*pb.Todos;
	result := dbcontroller.DB.Where("user_id = ?", logged_in_user_id).Find(&todos);

	if result.Error != nil {
		return nil, errors.New ("error finding todos") 
	}

	return &pb.ArrayOfTodos{
		TodoArr: todos,
	}, nil
} 

func (s *TodoServer) CreateTodo(ctx context.Context, req *pb.NewTodo) (*pb.Todos, error) {
	tokenPayload, tokenErr := verify_auth.VerifyAuthToken(ctx);

	if tokenErr != nil {
		return nil, errors.New(tokenErr.Error())
	}

	logged_in_user_id := tokenPayload.UserId
	
	if req.GetName() == "" || req.GetDescription() == "" {
		return nil, errors.New("please provide name & description");
	} 

	todo := &pb.Todos{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		UserId: 		 logged_in_user_id,
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
		UserId: todo.UserId,
		Done: false,
	}, nil
}

func (s *TodoServer) EditSingleTodo(ctx context.Context, req *pb.EditOrDeleteRequest) (*pb.Todos, error) {
	tokenPayload, tokenErr := verify_auth.VerifyAuthToken(ctx);

	if tokenErr != nil {
		return nil, tokenErr
	}

	logged_in_user_id := tokenPayload.UserId;

	todoId := req.Id;
	if todoId == "" {
		return nil, errors.New("please provide todo-id")
	}
	todo := &pb.Todos {
		Id: todoId,
	}

	result := dbcontroller.DB.First (&todo);
	if result.Error != nil || todo == nil {
		return nil, errors.New ("todo does not exist")
	}

	if todo.UserId != logged_in_user_id {
		return nil, errors.New("unauthorized")
	}

	if todo.Done {
		result := dbcontroller.DB.Model (&todo).Where ("id = ?", todoId).Update ("done", false)
		if result.Error != nil {
			return nil, errors.New ("error updating todo")
		}
	} else {
		result := dbcontroller.DB.Model (&todo).Where ("id = ?", todoId).Update ("done", true)
		if result.Error != nil {
			return nil, errors.New ("error updating todo")
		}
	}

	return todo, nil
}

func (s *TodoServer) DeleteSingleTodo(ctx context.Context, req *pb.EditOrDeleteRequest) (*pb.Todos, error) {
	tokenPayload, tokenErr := verify_auth.VerifyAuthToken(ctx);

	if tokenErr != nil {
		return nil, tokenErr
	}

	logged_in_user_id := tokenPayload.UserId;

	todoId := req.GetId();
	if todoId == "" {
		return nil, errors.New("please provide todo - id")
	}
	todo := &pb.Todos {
		Id: todoId,
	}

	result := dbcontroller.DB.First (&todo);
	if result.Error != nil || todo == nil {
		return nil, errors.New ("todo does not exist")
	}

	if todo.UserId != logged_in_user_id {
		return nil, errors.New("unauthorized")
	}

	dbcontroller.DB.Where("id = ?", todoId).Delete(&todo)

	return todo, nil
}