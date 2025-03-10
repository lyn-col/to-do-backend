package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Router     *gin.Engine
	HTTPServer *http.Server
}

func NewServer(port string, db *gorm.DB, logger *log.Logger) *Server {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins, restrict it for production
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Cache preflight request for 12 hours
	}))
	SetupRoutes(r, db, logger) // Pass DB to routes
	
	srv := &Server{
		Router: r,
		HTTPServer: &http.Server{
			Addr:    port,
			Handler: r,
		},
	}

	return srv
}

func (s *Server) Start() error {
	return s.HTTPServer.ListenAndServe()
}
