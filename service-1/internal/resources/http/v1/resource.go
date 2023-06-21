package v1

import (
	"soft-weater/internal/config"
	"soft-weater/internal/managers"
	"soft-weater/internal/resources/http/v1/dto"

	"github.com/gin-gonic/gin"
)

type Server struct {
	conf       config.Configuration
	userMan    managers.User
	accountMan managers.Account
	Router     *gin.Engine
}

func NewServer(conf config.Configuration, userMan managers.User, accountMan managers.Account) *Server {

	server := &Server{
		Router:     gin.Default(),
		conf:       conf,
		userMan:    userMan,
		accountMan: accountMan,
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

	account := v1.Group("/account")
	user := v1.Group("/user")

	user.POST("", srv.CreateUser)
	//user.GET("/:user_id", srv.GetUser)

	account.GET("/:user_id", srv.GetAccount)
	account.PUT("/debt/increase", srv.IncreaseDebt)
	account.PUT("/debt/decrease", srv.DecreaseDebt)

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
