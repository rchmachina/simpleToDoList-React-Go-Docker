package routes


import (
	handlers "github.com/rchmachina/soal2/be/handlers"
	repo "github.com/rchmachina/soal2/be/repository"
	"github.com/rchmachina/soal2/be/utils/database"


	"github.com/labstack/echo/v4"
)

func taskRoutes(e *echo.Group) {
	taskRepository := repo.RepositoryTask(database.DB)
	h := handlers.HandlerTask(taskRepository)

	e.GET("/task",h.GetAllTask)
	e.POST("/task",h.CreateTask)
	e.PUT("/task",h.UpdateTask)
	
	e.PUT("/taskComplete",h.UpdateTaskIsComplete)
	e.DELETE("/task/:id",h.DeleteTask)
	//e.DELETE("/room/:roomid", middleware.Auth(h.DeleteRoom))
	// e.get("user")
	//e.POST("/deleteUser/:id", middleware.Auth(h.FindUsersPeer))
}
