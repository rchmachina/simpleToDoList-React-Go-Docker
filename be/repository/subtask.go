package repositories

import (
	"encoding/json"



	"github.com/rchmachina/soal2/be/model"
	"github.com/rchmachina/soal2/be/utils/helper"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

// kontrak
type SubtaskRepository interface {
	CreateSubTaskRepository(task model.SubTask) error
	DeleteSubTaskRepository(int) error
	UpdateSubTaskRepository(task model.UpdateSubTask) error
	UpdateSubTaskIsCompleteRepository(task model.UpdateSubTask) error
}

func RepositorySubtask(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateSubTaskRepository(task model.SubTask) error {
	var result string

	
	paramsJSON, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = r.db.Raw("SELECT * FROM to_do_list.create_subtask($1::jsonb)", string(paramsJSON)).Scan(&result).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteSubTaskRepository(id int) error {

	var result string

	paramsJSON, err := json.Marshal(map[string]interface{}{"subTaskId": id})
	if err != nil {
		return err
	}

	err = r.db.Raw("SELECT * FROM to_do_list.delete_subtask($1::jsonb)", string(paramsJSON)).Scan(&result).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateSubTaskRepository(task model.UpdateSubTask) error {

	var result string

	params := helper.ToJSON(task)
	err := r.db.Raw("SELECT * from to_do_list.update_subtask($1::jsonb)", params).Scan(&result).Error
	if err != nil {
		return err
	}

	return nil
}
func (r *repository) UpdateSubTaskIsCompleteRepository(task model.UpdateSubTask) error {

	var result string

	paramsJSON, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = r.db.Raw("SELECT * from to_do_list.update_subtask_is_complete($1::jsonb)", string(paramsJSON)).Scan(&result).Error
	if err != nil {
		return err
	}

	return nil
}
