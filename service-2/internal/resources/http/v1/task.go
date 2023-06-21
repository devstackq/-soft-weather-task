package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"soft-weater/internal/models"
)

func (srv *Server) GetTasks(c *gin.Context) {

	tasks, err := srv.taskMan.GetList(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (srv *Server) UpdateTaskPrice(c *gin.Context) {
	var req models.Task
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = c.Param("task_id")

	if err := srv.taskMan.Update(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (srv *Server) SolveTask(c *gin.Context) {
	var req models.Solve

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.TaskID = c.Param("task_id")
	req.UserID = c.Query("user_id")

	res, err := srv.solverMan.Do(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (srv *Server) GetHistory(c *gin.Context) {

}
