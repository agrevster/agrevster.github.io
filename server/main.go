package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func main() {
	Init()
	r := gin.Default()

	prod := getEnvOrDefault("PRODUCTION", "false") == "true"
	apiVersion, err := strconv.ParseFloat(getEnvOrDefault("API_VERSION", "Error"), 64)
	check(err)

	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{
			"production":  prod,
			"api-version": apiVersion,
		})
	})

	projects, err := os.ReadFile("./projects.json")
	check(err)

	r.GET("/projects", func(c *gin.Context) {
		if !prod {
			c.Header("Access-Control-Allow-Origin", "*")
		}
		_, err := c.Writer.Write(projects)
		check(err)
	})

	skills, err := os.ReadFile("./skills.json")
	check(err)

	r.GET("/skills", func(c *gin.Context) {
		if !prod {
			c.Header("Access-Control-Allow-Origin", "*")
		}
		_, err := c.Writer.Write(skills)
		check(err)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
