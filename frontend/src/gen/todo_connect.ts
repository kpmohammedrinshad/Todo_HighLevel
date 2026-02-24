import type { CreateTodoRequest, CreateTodoResponse, ListTodosRequest, ListTodosResponse } from "./todo_pb.js";

const createTodoMethod = {
  name: "CreateTodo",
  kind: "unary" as const,
  I: {
    typeName: "todo.v1.CreateTodoRequest",
  },
  O: {
    typeName: "todo.v1.CreateTodoResponse",
  },
};

const listTodosMethod = {
  name: "ListTodos",
  kind: "unary" as const,
  I: {
    typeName: "todo.v1.ListTodosRequest",
  },
  O: {
    typeName: "todo.v1.ListTodosResponse",
  },
};

export const TodoService = {
  typeName: "todo.v1.TodoService",
  methods: [createTodoMethod, listTodosMethod],
};