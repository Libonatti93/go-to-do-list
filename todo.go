package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

// Estrutura da Tarefa
type Task struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

var tasks []Task
var nextID int = 1

// Função para listar todas as tarefas
func getTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

// Função para criar uma nova tarefa
func createTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var task Task
    _ = json.NewDecoder(r.Body).Decode(&task)
    task.ID = nextID
    nextID++
    tasks = append(tasks, task)

    json.NewEncoder(w).Encode(task)
}

// Função para atualizar uma tarefa
func updateTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for index, item := range tasks {
        if item.ID == id {
            tasks = append(tasks[:index], tasks[index+1:]...)

            var updatedTask Task
            _ = json.NewDecoder(r.Body).Decode(&updatedTask)
            updatedTask.ID = id
            tasks = append(tasks, updatedTask)

            json.NewEncoder(w).Encode(updatedTask)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"message": "Task not found"})
}

// Função para deletar uma tarefa
func deleteTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for index, item := range tasks {
        if item.ID == id {
            tasks = append(tasks[:index], tasks[index+1:]...)
            json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted"})
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{"message": "Task not found"})
}

func main() {
    // Inicializa o roteador
    r := mux.NewRouter()

    // Define os endpoints da API
    r.HandleFunc("/tasks", getTasks).Methods("GET")
    r.HandleFunc("/tasks", createTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

    // Inicia o servidor
    fmt.Println("Server running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
