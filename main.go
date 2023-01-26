package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"bmwgroup.net/qt56157/excel2json/_service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//lets create some tables if they DONT EXIST!!!
	// allowedHost := os.Getenv("ALLOWED")
	appAddr := "0.0.0.0:" + os.Getenv("PORT")
	// appAddr := "0.0.0.0:4000"

	// gin.SetMode(gin.ReleaseMode)
	// routes
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Origin", "Accept", "X-Requested-With", "Content-Type", "Authorization", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.POST("/api/v1/parse-excel", _service.ParseExcelUpload)
	r.POST("/api/v1/parse-csv", _service.ParseCSVUpload)
	r.GET("/api/v1/health", _service.HealthCheck)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	//server config

	srv := &http.Server{
		Addr:    appAddr,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
