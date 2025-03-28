package config

import (
	"fmt"
	"insuBridge/internal/model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// 환경 변수에서 데이터베이스 설정 가져오기
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// 필수 환경 변수 체크
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("Database configuration is missing in environment variables")
	}

	// DSN 문자열 생성
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 테이블 자동 마이그레이션
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
}
