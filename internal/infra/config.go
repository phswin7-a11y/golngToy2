package infra

import (
	"os"
)

// Config 애플리케이션 설정 구조체
type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
	AppPort string
}

// LoadConfig 환경 변수에서 설정을 로드합니다.
func LoadConfig() *Config {
	return &Config{
		DBUser: getEnv("DB_USER", "root"),
		DBPass: getEnv("DB_PASS", ""),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "3306"),
		DBName: getEnv("DB_NAME", "toy_db"),
		AppPort: getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
