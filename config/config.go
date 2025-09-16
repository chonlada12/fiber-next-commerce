package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppPort string
	AppURL  string
	DBHost  string
	DBPort  string
	DBName  string
	DBUser  string
	DBPass  string
	DBSSL   string
}

// loadConfig โหลดค่าจากไฟล์ .env และตั้งค่าเริ่มต้น
func LoadConfig() *Config {

	// โหลดไฟล์ .env
	err := godotenv.Load()

	// ถ้ามี error ให้ log ข้อความและออกจากโปรแกรม
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	return &Config{
		AppEnv:  os.Getenv("APP_ENV"),
		AppPort: os.Getenv("APP_PORT"),
		AppURL:  os.Getenv("APP_URL"),
		DBHost:  os.Getenv("DB_HOST"),
		DBPort:  os.Getenv("DB_PORT"),
		DBName:  os.Getenv("DB_NAME"),
		DBUser:  os.Getenv("DB_USER"),
		DBPass:  os.Getenv("DB_PASS"),
		DBSSL:   os.Getenv("DB_SSL"),
	}

}
