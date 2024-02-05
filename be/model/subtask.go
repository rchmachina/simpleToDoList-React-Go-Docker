package model

type SubTask struct {
	TaskID      int    `json:"taskId" validate:"required"`
	SubTaskName string `json:"subTaskName" validate:"required"`
	SubtaskId   int    `json:"subtaskId"`
	IsComplete  bool   `json:"isComplete" default:"false"`
}

type UpdateSubTask struct {
	TaskID      int    `json:"taskId" `
	SubTaskName string `json:"subTaskName"`
	SubtaskId   int    `json:"subTaskId" validate:"required"`
	IsComplete  bool   `json:"isComplete"`
}