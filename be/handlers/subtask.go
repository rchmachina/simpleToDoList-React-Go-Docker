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

type subtaskHandler struct {
	subtaskRepository repositories.SubtaskRepository
}

func HandlerSubTask(repository repositories.SubtaskRepository) *subtaskHandler {
	return &subtaskHandler{repository}
}

func (h *subtaskHandler) DeleteSubTask(c echo.Context) error {

	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return err
	}
	err = h.subtaskRepository.DeleteSubTaskRepository(intId)
	if err != nil {
		return helper.JSONResponse(c, 400, err)
	}

	return helper.JSONResponse(c, 200, "success")

}
func (h *subtaskHandler) CreateSubTask(c echo.Context) error {

	createTask := new(model.SubTask)
	if err := c.Bind(createTask); err != nil {
		return helper.JSONResponse(c, 400, err)
	}
	if err := helper.Validator(createTask); err != nil {
		log.Println(err)
		respErr := helper.RespError(err)
		return helper.JSONResponse(c, 400, respErr)
	}

	if err := h.subtaskRepository.CreateSubTaskRepository(*createTask); err != nil {
		return helper.JSONResponse(c, 400, err)
	}
	return helper.JSONResponse(c, 200, "success")

}
func (h *subtaskHandler) UpdateSubTask(c echo.Context) error {

	updateTask := new(model.UpdateSubTask)
	if err := c.Bind(updateTask); err != nil {
		return helper.JSONResponse(c, 400, err)
	}

	if err := helper.Validator(updateTask); err != nil {
		log.Println(err)
		respErr := helper.RespError(err)
		return helper.JSONResponse(c, 400, respErr)
	}
	if err := h.subtaskRepository.UpdateSubTaskRepository(*updateTask); err != nil {
		return helper.JSONResponse(c, 400, err)
	}
	return helper.JSONResponse(c, 200, "success")

}
func (h *subtaskHandler) UpdateSubTaskIsComplete(c echo.Context) error {
	log.Println("masuk pak ")
	updateTask := new(model.UpdateSubTask)
	if err := c.Bind(updateTask); err != nil {
		return helper.JSONResponse(c, 400, err)
	}

	if err := helper.Validator(updateTask); err != nil {
		log.Println(err)
		respErr := helper.RespError(err)
		return helper.JSONResponse(c, 400, respErr)
	}
	log.Println(updateTask)
	if err := h.subtaskRepository.UpdateSubTaskIsCompleteRepository(*updateTask); err != nil {
		return helper.JSONResponse(c, 400, err)
	}
	log.Println("passed pak cik")
	return helper.JSONResponse(c, 200, "success")

}
