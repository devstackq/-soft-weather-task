package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"soft-weater/internal/models"
	"soft-weater/internal/resources/http/v1/dto"
)

func (srv *Server) IncreaseDebt(c *gin.Context) {
	var req models.Account

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Description: err.Error()})
		return
	}

	if err := srv.accountMan.IncreaseDebt(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Description: err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
func (srv *Server) DecreaseDebt(c *gin.Context) {

	var req models.Account

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Description: err.Error()})
		return
	}

	if err := srv.accountMan.DecreaseDebt(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Description: err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (srv *Server) GetAccount(c *gin.Context) {
	uid := c.Param("user_id")
	res, err := srv.accountMan.Get(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Description: err.Error()})
		return
	}
	res.UserID = uid

	c.JSON(http.StatusOK, res)
}
