package routes

import (
    "TrabalhoFaculGolang/internal/handlers"
    taskrepo "TrabalhoFaculGolang/internal/repository/task"
    taskservice "TrabalhoFaculGolang/internal/service/task"
    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5/pgxpool" // ou o driver que você estiver usando
)

func SetupRoutes(router *gin.Engine, db *pgxpool.Pool) {
    repo := taskrepo.NewRepository(db)
    service := taskservice.NewService(repo)
    handler := handlers.NewTaskHandler(service)

    tasks := router.Group("/tasks")
    {
        tasks.POST("", handler.Create)
        tasks.GET("", handler.List)
        tasks.PUT("/attStatus/:id", handler.UpdateStatus)
    }
}
