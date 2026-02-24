package service

import (
	"context"
	"sync"

	todov1 "todo-backend/gen/todo/v1"

	"connectrpc.com/connect"
	"github.com/google/uuid"
)

type TodoService struct {
	mu    sync.RWMutex
	todos map[string]*todov1.Todo
}

func NewTodoService() *TodoService {
	return &TodoService{
		todos: make(map[string]*todov1.Todo),
	}
}

func (s *TodoService) CreateTodo(
	ctx context.Context,
	req *connect.Request[todov1.CreateTodoRequest],
) (*connect.Response[todov1.CreateTodoResponse], error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo := &todov1.Todo{
		Id:        uuid.NewString(),
		Title:     req.Msg.Title,
		Completed: false,
	}

	s.todos[todo.Id] = todo

	return connect.NewResponse(&todov1.CreateTodoResponse{
		Todo: todo,
	}), nil
}

func (s *TodoService) ListTodos(
	ctx context.Context,
	req *connect.Request[todov1.ListTodosRequest],
) (*connect.Response[todov1.ListTodosResponse], error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todos := make([]*todov1.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}

	return connect.NewResponse(&todov1.ListTodosResponse{
		Todos: todos,
	}), nil
}

func (s *TodoService) DeleteTodo(
	ctx context.Context,
	req *connect.Request[todov1.DeleteTodoRequest],
) (*connect.Response[todov1.DeleteTodoResponse], error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.todos[req.Msg.Id]
	if exists {
		delete(s.todos, req.Msg.Id)
	}

	return connect.NewResponse(&todov1.DeleteTodoResponse{
		Success: exists,
	}), nil
}
