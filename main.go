package main

import (
	"log"
	"net/http"
	"go-todo-app/config"
	"go-todo-app/controller"
	"go-todo-app/repository"
	"go-todo-app/service"

	"github.com/gorilla/mux"
)

func main() {
	// init db
	config.Connect();
	// init repo, service, controller
	taskRepository := repository.NewTaskRepository(config.DB);
	taskService := service.NewTaskService(taskRepository);
	taskController := controller.NewTaskController(taskService);

	// setup Routes
	r:= mux.NewRouter();
	r.HandleFunc("/tasks", taskController.GetTaskList).Methods("GET")
	r.HandleFunc("/tasks/new", taskController.NewTaskForm).Methods("GET")
	r.HandleFunc("/tasks", taskController.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/edit/{id}", taskController.EditTaskForm).Methods("GET")
    r.HandleFunc("/tasks/{id}", taskController.UpdateTask).Methods("POST")
    r.HandleFunc("/tasks/delete/{id}", taskController.DeleteTask).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", r))

}