package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type JWTConfig struct {
	Secret string
	Expire int
}

var (
	envLoaded bool
	envOnce   sync.Once
)

func LoadConfig() (*Config, error) {
	// SERVER_PORT=8080
	// SERVER_MODE=debug

	// DB_HOST=localhost
	// DB_PORT=3306
	// DB_USER=root
	// DB_PASSWORD=root
	// DB_NAME=blog

	// JWT_SECRET=blogsystemsecret
	// JWT_EXPIRE=24
	loadEnv()
	serverPort := os.Getenv("SERVER_PORT")
	serverMode := os.Getenv("SERVER_MODE")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExpire, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE"))

	// viper := viper.New()
	// viper.AddConfigPath(`D:\project\goproject\homeWork\blogSystem`)
	// viper.SetConfigType("env")
	// viper.SetConfigFile(".env")
	// viper.AutomaticEnv()
	// if err := viper.ReadInConfig(); err != nil {
	// 	return nil, err
	// }

	config := &Config{
		ServerConfig{
			Port: serverPort,
			Mode: serverMode,
		},
		DatabaseConfig{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
			DBName:   dbName,
		},
		JWTConfig{
			Secret: jwtSecret,
			Expire: jwtExpire,
		},
	}

	// if err := viper.Unmarshal(config); err != nil {
	// 	return nil, err
	// }

	return config, nil
}

func loadEnv() {
	envOnce.Do(func() {
		// Get the directory where this file (helpers.go) is located
		// runtime.Caller(0) gives us the location of this function (loadEnv)
		// which is in helpers.go in the testutil package
		_, currentFile, _, ok := runtime.Caller(0)
		if !ok {
			return
		}

		// Get the examples directory (parent of testutil)
		// currentFile is something like: /path/to/examples/testutil/helpers.go
		testutilDir := filepath.Dir(currentFile)
		examplesDir := filepath.Dir(testutilDir)

		// Load .env file from examples directory
		envPath := filepath.Join(examplesDir, ".env")
		if err := godotenv.Load(envPath); err != nil {
			// .env file is optional, so we don't fail if it doesn't exist
			// Environment variables can still be set directly via system environment
			return
		}
		envLoaded = true
	})
}
