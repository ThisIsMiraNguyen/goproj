package views

import (
	"go-todo-app/model"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("views/*.html"))

func RenderTaskListView(w http.ResponseWriter, tasks []model.Task) {
    data := struct {
        Tasks []model.Task
    }{
        Tasks: tasks,
    }

    err := templates.ExecuteTemplate(w, "task-list.html", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func RenderTaskFormView(w http.ResponseWriter, task model.Task) {
    data := struct {
        Task model.Task
    }{
        Task: task,
    }

    err := templates.ExecuteTemplate(w, "task-form.html", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
