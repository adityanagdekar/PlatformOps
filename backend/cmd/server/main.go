package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type healthResponse struct {
	Service   string `json:"service"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type notImplementedResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func main() {
	router := setupRouter()

	addr := ":" + envOrDefault("PORT", "8080")
	server := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("platformops backend listening on %s", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/health", healthHandler)
	router.POST("/deploy", deployHandler)
	router.GET("/services", servicesHandler)
	router.GET("/deployments", deploymentsHandler)
	router.POST("/rollback", rollbackHandler)

	return router
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, healthResponse{
		Service:   "platformops-backend",
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

func deployHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, notImplementedResponse{
		Error:   "not_implemented",
		Message: "deployment config loading and Kubernetes apply logic starts next in Phase 3",
	})
}

func servicesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"services": []string{},
	})
}

func deploymentsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"deployments": []string{},
	})
}

func rollbackHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, notImplementedResponse{
		Error:   "not_implemented",
		Message: "rollback metadata and image restore logic starts after deployment metadata exists",
	})
}

func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
