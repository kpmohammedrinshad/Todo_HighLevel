export interface Todo {
  id: string;
  title: string;
  completed: boolean;
}

export async function createTodo(title: string): Promise<Todo> {
  const response = await fetch("http://localhost:8080/todo.v1.TodoService/CreateTodo", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ title }),
  });
  const data = await response.json();
  return data.todo;
}

export async function listTodos(): Promise<Todo[]> {
  const response = await fetch("http://localhost:8080/todo.v1.TodoService/ListTodos", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({}),
  });
  const data = await response.json();
  return data.todos || [];
}