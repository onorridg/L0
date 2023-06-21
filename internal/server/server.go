package server

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"l0/internal/models"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type PageData struct {
	JSON string
}

func handleIndex(c *gin.Context) {
	tmpl, err := template.ParseFiles("frontend/index.html")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	data := PageData{
		JSON: "{}",
	}

	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func handleGetJSON(c *gin.Context) {
	id := c.Query("id")
	log.Println("ID:", id)

	jsonFile, err := os.Open("cmd/sender/model.json")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer jsonFile.Close()
	byteJson, err := io.ReadAll(jsonFile)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var order models.Order
	err = json.Unmarshal(byteJson, &order)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	jsonData, err := json.Marshal(&order)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.Writer.Write(jsonData)
}

func Run() {
	//http.HandleFunc("/", handleIndex)
	//http.HandleFunc("/get-json", handleGetJSON)
	//
	//log.Fatal(http.ListenAndServe(":8080", nil))

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	basePath := r.Group("/")
	{
		basePath.GET("", handleIndex)
		basePath.GET("/get-json", handleGetJSON)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer func() {
		cancel()
	}()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 1 second.")
	}
	log.Println("Server exiting")

}
