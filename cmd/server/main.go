package main

import (
	"log"

	"insuBridge/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting InsuBridge server...")

	// Gin 엔진 초기화
	r := gin.Default()

	// 핸들러 초기화
	userHandler := handler.NewUserHandler()

	// 라우트 설정
	r.POST("/api/users", userHandler.CreateUser)

	// 서버 시작
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
