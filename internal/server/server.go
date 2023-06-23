package server

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"html/template"
	"l0/internal/postgresql"
	"l0/pkg/inMemory"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
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
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	order := inMemory.Conn().QueryOrder(id)
	if order != nil {
		c.Header("X-Cache-Status", "Hit")
	} else {
		c.Header("X-Cache-Status", "Miss")

		db := postgresql.Conn()
		defer db.Conn.Close()

		var orderId uint64
		orderId, order, err = db.SelectUserOrder(id)
		if err == nil {
			inMemory.Conn().InsertData(orderId, &order)
		}
	}

	var jsonData []byte
	jsonData, err = json.Marshal(&order)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.Writer.Write(jsonData)
}

func Run() {
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
	log.Print("[+] Frontend server started.")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("[*] Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer func() {
		cancel()
	}()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("[!] Server shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("[*] Timeout of 1 second.")
	}
	log.Println("[+] Server exiting")

}
