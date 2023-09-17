package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func getEnvOrDefault(envVariable string, defaultValue string) string {
	v := os.Getenv(envVariable)
	if v != "" {
		return v
	} else {
		return defaultValue
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Init() {
	err := godotenv.Load()
	check(err)
	if getEnvOrDefault("PRODUCTION", "false") == "true" {
		gin.SetMode("release")
	}

}
