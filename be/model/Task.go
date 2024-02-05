package model



type Task struct{
	TaskID int `json:"taskId"`
	TaskName string `json:"taskName" validate:"required"`
	Deadline interface{} `json:"deadline"`
	IsComplete bool `json:"isComplete" default:"false"`
	SubTask []SubTask `json:"subtasks"`
}

type UpdateTask struct{
	TaskID int `json:"taskId"`
	TaskName string `json:"taskName"`
	Deadline string `json:"deadline"`
	IsComplete bool `json:"isComplete" `
	SubTask []SubTask `json:"subtasks"`
}