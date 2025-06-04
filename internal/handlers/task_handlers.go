package handlers

import (
	"fmt"
   "net/http"
    "context"
    "github.com/gin-gonic/gin"
	taskdomain "TrabalhoFaculGolang/internal/domain/task"
    taskservice "TrabalhoFaculGolang/internal/service/task"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type TaskHandler struct {
	service *taskservice.Service
}

func NewTaskHandler(service *taskservice.Service) *TaskHandler {
	return &TaskHandler{service: service}
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
func (h *TaskHandler) Create(c *gin.Context) {
	var t taskdomain.Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Printf("Erro ao listar: %v\n", err)
		return
	}
	if err := h.service.CreateTask(c.Request.Context(), &t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar task"})
		return
	}
	c.JSON(http.StatusCreated, t)
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
func (h *TaskHandler) List(c *gin.Context) {
	tasks, err := h.service.ListTasks(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar tasks"})
		fmt.Printf("Erro ao listar: %v\n", err)
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
func (h *TaskHandler) UpdateStatus(c *gin.Context) {
	var payload struct {
		Completed bool `json:"completed"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	if err := h.service.UpdateTaskStatus(c.Request.Context(), id, payload.Completed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar status"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Status atualizado"})
}