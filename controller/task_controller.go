package controller

import (
	"go-todo-app/model"
	"go-todo-app/service"
	"go-todo-app/views"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type TaskController struct {
	service service.ITaskService
}

func NewTaskController(service service.ITaskService) *TaskController {
	return &TaskController{service}
}
func (c *TaskController) GetTaskList ( w http.ResponseWriter, r *http.Request){
 tasks, err := c.service.GetTasks()
 if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
 }
 views.RenderTaskListView(w, tasks)
}

func (c *TaskController) NewTaskForm(w http.ResponseWriter, r *http.Request){
	views.RenderTaskFormView(w, model.Task{})
}
func (c *TaskController) EditTaskForm(w http.ResponseWriter, r *http.Request){
	
	param:= mux.Vars(r);
	id, _:= strconv.Atoi(param["id"]);
	task, err := c.service.GetTaskById(id);
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound);
		return
	}
	views.RenderTaskFormView(w, task)
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
    completed := r.FormValue("completed") == "on"

    task := model.Task{
        Title:     title,
        Completed: completed,
    }

    if err := c.service.CreateTask(task); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}
func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request){
	param := mux.Vars(r);
	id, _ := strconv.Atoi(param["id"]);
	title := r.FormValue("title")
	completed := r.FormValue("completed") =="on"
	task:= model.Task{
		Id: id,
		Title: title,
		Completed: completed,
	}
	if err:= c.service.UpdateTask(task); err!= nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}
func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request){

	param := mux.Vars(r);
	id, _ := strconv.Atoi(param["id"]);
	if err:= c.service.DeleteTask(id); err!= nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)

}