package main

import (
	"fmt"
	"mymachine707/config"
	"mymachine707/docs"
	"mymachine707/handlars"
	storageinterface "mymachine707/storage/StorageInterface"
	"mymachine707/storage/postgres"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	psqlConfigString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.AppVersion
	var err error
	var stg storageinterface.StorageInterface

	stg, err = postgres.Initdb(psqlConfigString)

	if err != nil {
		panic(err)
	}
	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	if cfg.Environment != "development" {
		r.Use(gin.Logger(), gin.Recovery())
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	h := handlars.Handlars{
		Stg: stg,
	}

	// Gruppirovka qilindi
	v1 := r.Group("v1")
	{
		v1.POST("/person", h.RootCreatePerson)
		v1.GET("/person/:id", h.RootGetPersonByID)
		v1.GET("/person", h.RootGetPersonList)
		v1.PUT("/person", h.RootUpdatePerson)
		v1.DELETE("/person/:id", h.RootDeletePerson)

		v1.POST("/lesson", h.RootCreateLesson)
		v1.GET("/lesson/:id", h.RootGetLessonByID)
		v1.GET("/lesson", h.RootGetLessonList)
		v1.PUT("/lesson", h.RootUpdateLesson)
		v1.DELETE("/lesson/:id", h.RootDeleteLesson)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(cfg.HTTPPort)
}
