package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Akanibekuly/golang_test_task1.git/internal/config"
	"github.com/Akanibekuly/golang_test_task1.git/internal/delivery"
	"github.com/Akanibekuly/golang_test_task1.git/internal/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	cfg    *config.Config
	router *gin.Engine
	db     *sql.DB
}

func (a *App) Initialize(conf *config.Config) {
	gin.SetMode(conf.App.Mode)

	a.cfg = conf
	a.router = gin.New()

	// log.Println("works with mock db - func Initialize line 30")
	db, err := repository.NewPostgresDB(conf.DB)
	if err != nil {
		log.Println(err)
		return
	}
	a.db = db
	// a.db = nil

	a.router.Use(gin.Logger())
	a.router.Use(gin.Recovery())
	a.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		MaxAge:           30 * time.Second,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{""},
		AllowCredentials: true,
	}))

	a.setComponents()
}

func (a *App) Run(ctx context.Context) {
	srv := http.Server{
		Addr:           a.cfg.App.Port,
		Handler:        a.router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	go func() {
		log.Println("Starting server on port", a.cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	defer a.db.Close()
	log.Println("Server exiting")
}

func (a *App) setComponents() {
	{
		apiV1 := a.router.Group("/api/v1")
		delivery.SetEndpoints(apiV1, a.db)
	}
}
