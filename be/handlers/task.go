package handlers

import (
	"log"
	"strconv"

	//"time"

	"github.com/labstack/echo/v4"
	"github.com/rchmachina/soal2/be/model"
	repositories "github.com/rchmachina/soal2/be/repository"
	"github.com/rchmachina/soal2/be/utils/helper"
)

type taskHandler struct {
	taskRepository repositories.TaskRepository
}

func HandlerTask(repository repositories.TaskRepository) *taskHandler {
	return &taskHandler{repository}
}

func (h *taskHandler) GetAllTask(c echo.Context) error {
	var completeTask []model.Task
	var notCompleteTask []model.Task
	getAllTask, err := h.taskRepository.GetAllTaskRepository()
	if err != nil {
		return helper.JSONResponse(c, 400, err)
	}
	for i := 0; i < len(getAllTask); i++ {
		if (getAllTask[i].IsComplete){
			completeTask = append(completeTask, getAllTask[i])
			continue
		}
		notCompleteTask = append(notCompleteTask, getAllTask[i])
	}


	data := map[string]interface{}{
		"completeTask": completeTask,
		"nonCompleteTask": notCompleteTask,
	}

	return helper.JSONResponse(c, 200, data)

}
func (h *taskHandler) DeleteTask(c echo.Context) error {

	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return err
	}
	err = h.taskRepository.DeleteTaskRepository(intId)
	if err != nil {
		return helper.JSONResponse(c, 400, err)
	}

	return helper.JSONResponse(c, 200, "success")

}
func (h *taskHandler) CreateTask(c echo.Context) error {

	createTask := new(model.Task)
	if err := c.Bind(createTask); err != nil {
		return helper.JSONResponse(c, 400, err)
	}
	if err := helper.Validator(createTask); err != nil {
		return helper.JSONResponse(c, 400, err)
	}

	if(createTask.Deadline.(string) == "") {
		createTask.Deadline = nil
	}

	if err := h.taskRepository.CreateTaskRepository(*createTask); err != nil {
		return helper.JSONResponse(c, 400, err)
	}
	return helper.JSONResponse(c, 200, "success")

}
func (h *taskHandler) UpdateTask(c echo.Context) error {
	log.Println("masuk")
	updateTask := new(model.UpdateTask)
	if err := c.Bind(updateTask); err != nil {
		return helper.JSONResponse(c, 400, err)
	}
	if err := helper.Validator(updateTask); err != nil {
		log.Println(err)
		respErr := helper.RespError(err)
		return helper.JSONResponse(c, 400, respErr)
	}
	if err := h.taskRepository.UpdateTaskRepository(*updateTask); err != nil {
		return helper.JSONResponse(c, 400, err)
	}
	return helper.JSONResponse(c, 200, "success")

}
func (h *taskHandler) UpdateTaskIsComplete(c echo.Context) error {

	updateTask := new(model.UpdateTask)
	if err := c.Bind(updateTask); err != nil {
		return helper.JSONResponse(c, 400, err)
	}

	if err := helper.Validator(updateTask); err != nil {
		log.Println(err)
		respErr := helper.RespError(err)
		return helper.JSONResponse(c, 400, respErr)
	}
	if err := h.taskRepository.UpdateTaskIsCompleteRepository(*updateTask); err != nil {
		return helper.JSONResponse(c, 400, err)
	}
	log.Println("passed pak cik")
	return helper.JSONResponse(c, 200, "success")

}
