package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Done    bool   `json:"done`
}

var tasks []Task

func main() {
	r := mux.NewRouter()

	//Define Routes
	r.HandleFunc("/tasks", getTasks).Methods("GET")
	r.HandleFunc("/tasks", getTask).Methods("GET")
	r.HandleFunc("/tasks", createTask).Methods("POST")
	r.HandleFunc("/tasks", updateTask).Methods("PUT")
	r.HandleFunc("/tasks", deleteTask).Methods("DELETE")

	//start server
	log.Fatal(http.ListenAndServe(":4000", r))

}

//get all tasks

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

//get specific tasks

func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) //Get route params
	for _, item := range tasks {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{}) //Return empty task if not found
}

//Create a task(POST /tasks)

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = fmt.Sprintf("%d", len(tasks)+1) //simple ID generation
	tasks = append(tasks, task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// update a task(PUT /tasks/{id})

func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			var updatedTask Task
			_ = json.NewDecoder(r.Body).Decode(&updatedTask)
			updatedTask.ID = params["id"]
			tasks = append(tasks, updatedTask)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTask)
			return

		}
	}
}

// delete a task

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
}
