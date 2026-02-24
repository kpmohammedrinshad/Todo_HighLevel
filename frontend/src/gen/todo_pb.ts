export interface Todo {
  id: string;
  title: string;
  completed: boolean;
}

export interface CreateTodoRequest {
  title: string;
}

export interface CreateTodoResponse {
  todo?: Todo;
}

export interface ListTodosRequest {}

export interface ListTodosResponse {
  todos: Todo[];
}