package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"soft-weater/internal/models"
	"soft-weater/internal/resources/http/v1/dto"
)

// CreateUser
// @Tags     				/user
// @Title Method for Create New User
// @Accept 					json
// @Produce 				json
// @Param		 req		body     models.User	true	"user"
// @Success  201
// @Failure  401   {object}  dto.ErrorResponse
// @Failure  400   {object}  dto.ErrorResponse
// @Failure  500   {object}  dto.ErrorResponse
// @Router   /user [post]
func (srv *Server) CreateUser(c *gin.Context) {
	var request models.User

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Description: err.Error(),
		})
		return
	}

	if err := srv.userMan.Create(c.Request.Context(), request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Description: err.Error(),
		})
		return
	}
	c.Status(201)
}
