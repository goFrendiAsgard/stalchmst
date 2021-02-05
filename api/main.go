package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func getHTTPPort() string {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		return "3000"
	}
	return port
}

func createJSONResponse(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(500, gin.H{
			"status":        "error",
			"error_message": fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":        "ok",
		"error_message": "",
		"data":          data,
	})
}

func main() {
	tm, err := CreateDefaultTemplateModel()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	// default API endpoint
	r.Any("/", func(c *gin.Context) {
		createJSONResponse(c, "API is ready", nil)
	})

	// delay
	r.GET("/delay/:duration", func(c *gin.Context) {
		durationStr := c.Param("duration")
		duration, err := time.ParseDuration(durationStr)
		if err != nil {
			c.JSON(500, gin.H{
				"status":        "error",
				"error_message": fmt.Sprintf("cannot parse %s", durationStr),
			})
		}
		time.Sleep(duration)
		c.JSON(200, gin.H{
			"status":        "ok",
			"error_message": "",
			"data":          fmt.Sprintf("processed in %s", durationStr),
		})
	})

	// get all templates
	r.GET("/templates", func(c *gin.Context) {
		templateRows, err := tm.GetAll()
		createJSONResponse(c, templateRows, err)
	})

	// get single template
	r.GET("/template/:code", func(c *gin.Context) {
		templateRow, err := tm.GetByCode(c.Param("code"))
		createJSONResponse(c, templateRow, err)
	})

	// generate
	r.POST("/generate", func(c *gin.Context) {
		templateStr := c.PostForm("template")
		data := c.PostForm("data")
		generatedData, err := Generate(templateStr, data)
		createJSONResponse(c, generatedData, err)
	})

	r.Run(fmt.Sprintf(":%s", getHTTPPort()))

}
