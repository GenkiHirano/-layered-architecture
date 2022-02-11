package main

import (
	"github.com/GenkiHirano/layered-architecture/config"
	"github.com/GenkiHirano/layered-architecture/infra"
	"github.com/GenkiHirano/layered-architecture/interface/handler"
	"github.com/GenkiHirano/layered-architecture/service"
	"github.com/gin-gonic/gin"
)

func main() {
	taskRepository := infra.NewTaskRepository(config.NewDB())
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	r := gin.New()
	handler.InitRouting(r, taskHandler)
	r.Run()
}
