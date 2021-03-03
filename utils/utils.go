package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "..")
}

func GoDotEnvVariable(key string) string {
	fmt.Println(RootDir())
	envFileName := filepath.Join(RootDir(), "conf", ".env")
	err := godotenv.Load(envFileName)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
