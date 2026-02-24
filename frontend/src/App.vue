<script setup lang="ts">
import { ref, onMounted } from "vue";
import { createTodo, listTodos, type Todo } from "./client";

const title = ref("");
const todos = ref<Todo[]>([]);
const error = ref("");

async function fetchTodos() {
  try {
    todos.value = await listTodos();
    error.value = "";
  } catch (err) {
    error.value = "Error: " + err;
    console.error(err);
  }
}

async function addTodo() {
  if (!title.value.trim()) return;
  
  try {
    await createTodo(title.value);
    title.value = "";
    await fetchTodos();
  } catch (err) {
    error.value = "Error: " + err;
    console.error(err);
  }
}

onMounted(fetchTodos);
</script>

<template>
  <div style="padding: 2rem; max-width: 600px; margin: 0 auto;">
    <h1>📝 Todo App (Connected to Backend!)</h1>
    
    <div style="margin: 1rem 0;">
      <input 
        v-model="title" 
        @keyup.enter="addTodo"
        placeholder="Enter todo..."
        style="padding: 0.5rem; width: 70%; font-size: 1rem;"
      />
      <button 
        @click="addTodo"
        style="padding: 0.5rem 1rem; margin-left: 0.5rem; font-size: 1rem;"
      >
        Add
      </button>
    </div>

    <div v-if="error" style="color: red; margin: 1rem 0;">
      {{ error }}
    </div>

    <ul style="list-style: none; padding: 0;">
      <li 
        v-for="t in todos" 
        :key="t.id"
        style="padding: 0.75rem; border: 1px solid #ddd; margin: 0.5rem 0; border-radius: 4px;"
      >
        {{ t.title }}
      </li>
    </ul>

    <div v-if="todos.length === 0 && !error" style="color: #999; margin-top: 2rem;">
      No todos yet. Add one above!
    </div>
  </div>
</template>