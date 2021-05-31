package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tuanden0/simple_api/internal/models"
	"github.com/tuanden0/simple_api/internal/repository"
	"github.com/tuanden0/simple_api/internal/services"
)

const ADDR string = ":8000"

func main() {

	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()

	// Logging to a file.
	var f *os.File
	if _, err := os.Stat("gin.log"); os.IsNotExist(err) {
		f, _ = os.Create("gin.log")
	} else {
		f, _ = os.OpenFile("gin.log", os.O_APPEND, 0666)
	}

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.New()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("[%s] - %s\t\"%s\t%s\t%s\t%d\t%s\t%s\"\n",
			param.TimeStamp.Format(time.RFC1123),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	// Connect and Migrate Database
	db := models.ConnectDatabase()

	// Create Student repo
	studentRepo := repository.NewStudentRepo(db)

	// Create Student Service
	studentSrv := services.NewStudentService(studentRepo)

	// Grouping route to versionning API
	studentV1 := router.Group("/v1")
	{
		studentV1.GET("/students/", studentSrv.List)
		studentV1.GET("/student/:id", studentSrv.Retrieve)
		studentV1.POST("/student", studentSrv.Create)
		studentV1.PATCH("/student/:id", studentSrv.Update)
		studentV1.DELETE("/student/:id", studentSrv.Delete)
	}

	router.Run(ADDR)

}
