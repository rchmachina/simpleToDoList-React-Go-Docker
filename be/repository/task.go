package repositories

import (
	"encoding/json"

	"fmt"
	"log"

	"github.com/rchmachina/soal2/be/model"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

// kontrak
type TaskRepository interface {
	CreateTaskRepository(task model.Task) error
	DeleteTaskRepository(int) error
	UpdateTaskRepository(task model.UpdateTask) error
	GetAllTaskRepository() ([]model.Task, error)
	UpdateTaskIsCompleteRepository(task model.UpdateTask) error
}

func RepositoryTask(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTaskRepository(task model.Task) error {
	var result string

	
	paramsJSON, err := json.Marshal(task)
	if err != nil {
		return err
	}

	if err := r.db.Raw("SELECT * FROM to_do_list.create_task($1::jsonb)", string(paramsJSON)).Scan(&result).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteTaskRepository(id int) error {

	var result string

	paramsJSON, err := json.Marshal(map[string]interface{}{"taskId": id})
	if err != nil {
		return err
	}
	if err := r.db.Raw("SELECT * FROM to_do_list.delete_task($1::jsonb)", string(paramsJSON)).Scan(&result).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateTaskRepository(task model.UpdateTask ) error{

	var result string

	paramsJSON, err := json.Marshal(task)
	if err != nil {
		return err
	}

	if err := r.db.Raw("SELECT * from to_do_list.update_task($1::jsonb)", string(paramsJSON)).Scan(&result).Error; err != nil {
		return err
	}

	return nil
}
func (r *repository) UpdateTaskIsCompleteRepository(task model.UpdateTask) error {

	var result string

	paramsJSON, err := json.Marshal(task)
	if err != nil {
		return err
	}

	if err := r.db.Raw("SELECT * from to_do_list.update_task_is_complete($1::jsonb)", string(paramsJSON)).Scan(&result).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllTaskRepository() ([]model.Task, error) {

	var getAlltask []model.Task
	paramsJSON, err := json.Marshal("")

	if err != nil {
		fmt.Println("ko ade error ke?")
		return getAlltask, err
	}

	var result string
	if err := r.db.Raw("select * from to_do_list.get_tasks_with_subtasks($1::jsonb)", string(paramsJSON)).Scan(&result).Error; err != nil {
		return getAlltask, err
	}

	err = json.Unmarshal([]byte(result), &getAlltask)
	if err != nil {
		log.Println(err)
		return getAlltask, err
	}

	return getAlltask, err
}
