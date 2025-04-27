package handlers

import (
	"net/http"

	"github.com/cancellap/TrabalhoFaculGolang/config"
	"github.com/cancellap/TrabalhoFaculGolang/models"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// CreateTask godoc
// @Summary Cria uma nova tarefa
// @Description Adiciona uma nova tarefa à lista
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Dados da Tarefa"
// @Success 201 {object} models.Task
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [post]
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	err := models.CreateTask(config.DB, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Erro ao criar task no banco"})
		return
	}
	c.JSON(http.StatusCreated, task)
}
