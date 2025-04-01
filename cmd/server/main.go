package main

import (
	"log"
	"time"

	"insuBridge/internal/config"
	"insuBridge/internal/handler"
	"insuBridge/internal/repository"
	"insuBridge/internal/service"

	"github.com/gin-contrib/cors"
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

	// CORS 설정
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",      // 로컬 개발 환경
			"http://localhost:5174",      // 로컬 개발 환경
			"http://3.139.6.169:5173",    // EC2 프론트엔드
			"http://3.139.6.169:8080",    // EC2 백엔드
			"https://insubridge.com",     // 프로덕션 환경
			"https://www.insubridge.com", // 프로덕션 환경
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
