package repository

import (
	"database/sql"
	"go-todo-app/model"
)

type TaskRepository struct{
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) ITaskRepository {
	return &TaskRepository{db}
}
func (r *TaskRepository) CreateTask(task model.Task) error {
_, err := r.db.Query("INSERT INTO todos (title, completed) VALUES(?,?)", task.Title, task.Completed)
return err;
}

func (r *TaskRepository) GetTasks()([]model.Task, error){
rows, err := r.db.Query("SELECT id, title, completed FROM todos")
if err != nil {
	return nil, err
}
defer rows.Close()
var tasks []model.Task
for rows.Next(){
	var task model.Task
	err = rows.Scan(&task.Id,&task.Title, &task.Completed)
	if(err != nil){
		return nil, err
	}
	tasks = append(tasks,task)
}
return tasks, nil
}

func (r *TaskRepository) GetTaskById(id int)(model.Task, error){
var task model.Task
err:= r.db.QueryRow("SELECT id, title, completed FROM todos WHERE id = ?", id).Scan(&task.Id, &task.Title, &task.Completed)
return task, err
}

func (r *TaskRepository)UpdateTask(task model.Task)error{
	_, err:= r.db.Exec("UPDATE todos SET title = ?, completed=? WHERE id = ?", task.Title, task.Completed,task.Id)
	return err
}

func (r *TaskRepository) DeleteTask(id int) error{
	_, err := r.db.Exec("DELETE FROM todos WHERE id =?", id)
	return err
}
