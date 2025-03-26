package main

import (
	"log"

	"insuBridge/internal/config"
	"insuBridge/internal/handler"
	"insuBridge/internal/repository"
	"insuBridge/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting InsuBridge server...")

	// 데이터베이스 초기화
	config.InitDB()

	// 의존성 주입
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Gin 엔진 초기화
	r := gin.Default()

	// 라우트 설정
	r.POST("/api/users", userHandler.CreateUser)

	// 서버 시작
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
