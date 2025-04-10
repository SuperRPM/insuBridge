package main

import (
	"fmt"
	"log"
	"time"

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
		fmt.Println(err)
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
	r.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "http://localhost:5173" || origin == "http://3.139.6.169:5173" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 라우트 설정
	api := r.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
				"time":   time.Now().Format(time.RFC3339),
			})
		})
		api.POST("/users", userHandler.CreateUser)
	}

	// 서버 시작
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
