package main

import (
	"log"

	"insuBridge/internal/config"
	"insuBridge/internal/handler"
	"insuBridge/internal/repository"
	"insuBridge/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 환경 변수 로드
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Starting InsuBridge server...")

	// 데이터베이스 초기화
	config.InitDB()

	// 의존성 주입
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Gin 엔진 초기화
	r := gin.Default()

	// CORS 설정 - 모든 요청 허용
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "86400") // 24시간

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 라우트 설정
	api := r.Group("/api")
	{
		api.POST("/users", userHandler.CreateUser)
	}

	// 서버 시작
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
