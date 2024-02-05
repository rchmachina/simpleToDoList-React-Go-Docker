package routes

import (
	handlers "github.com/rchmachina/soal2/be/handlers"
	repo "github.com/rchmachina/soal2/be/repository"
	"github.com/rchmachina/soal2/be/utils/database"

	"github.com/labstack/echo/v4"
)

func subTaskRoutes(e *echo.Group) {
	subTaskRepository := repo.RepositorySubtask(database.DB)
	h := handlers.HandlerSubTask(subTaskRepository)

	e.POST("/subtask", h.CreateSubTask)
	e.PUT("/subtask", h.UpdateSubTask)
	e.PUT("/subtaskComplete", h.UpdateSubTaskIsComplete)
	e.DELETE("/subtask/:id", h.DeleteSubTask)
	//e.DELETE("/room/:roomid", middleware.Auth(h.DeleteRoom))
	// e.get("user")
	//e.POST("/deleteUser/:id", middleware.Auth(h.FindUsersPeer))
}
