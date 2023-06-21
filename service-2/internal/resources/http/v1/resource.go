package v1

import (
	"soft-weater/internal/config"
	"soft-weater/internal/managers"
	"soft-weater/internal/resources/http/v1/dto"

	"github.com/gin-gonic/gin"
)

type Server struct {
	conf      config.Configuration
	taskMan   managers.Task
	solverMan managers.Solver
	Router    *gin.Engine
}

func NewServer(conf config.Configuration,
	taskMan managers.Task,
	solveMane managers.Solver) *Server {

	server := &Server{
		Router:    gin.Default(),
		conf:      conf,
		taskMan:   taskMan,
		solverMan: solveMane,
	}

	server.Register(server.Router)

	server.Router.RedirectTrailingSlash = true
	server.Router.RedirectFixedPath = true
	server.Router.HandleMethodNotAllowed = true

	return server
}

func (srv *Server) Register(engine *gin.Engine) {

	v1 := engine.Group("/v1")
	v1.Use(CORS)

	task := v1.Group("/task")

	task.POST("/:task_id/solve", srv.SolveTask)
	task.PUT("/:task_id/price", srv.UpdateTaskPrice)
	task.GET("", srv.GetTasks)
	task.GET("/history/:user_id", srv.GetHistory)

	engine.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(405, []dto.ErrorResponse{{
			Description: "Unsupported method",
		},
		})
	})

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, []dto.ErrorResponse{{
			Description: "Route not found",
		},
		})
	})
}
