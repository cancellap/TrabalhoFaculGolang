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

// GetTasks godoc
// @Summary Lista todas as tarefas
// @Description Retorna todas as tarefas cadastradas
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} models.Task
// @Failure 500 {object} map[string]string
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	tasks, err := models.GetTasks(config.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Erro ao buscar tasks no banco"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// UpdateTaskStatus godoc
// @Summary Atualiza o status de uma tarefa
// @Description Atualiza o status de uma tarefa existente
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "ID da Tarefa"
// @Param payload body UpdateTaskStatusRequest true "Status da Tarefa"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/attStatus/{id} [put]
func UpdateTaskStatus(c *gin.Context) {
	var payload UpdateTaskStatusRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	err := models.UpdateTaskStatus(config.DB, id, payload.Completed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar status da task no banco"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status atualizado com sucesso"})
}

type UpdateTaskStatusRequest struct {
	Completed bool `json:"completed" example:"true"`
}
