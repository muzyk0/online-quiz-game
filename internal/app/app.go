package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/muzyk0/online-quiz-game/internal/app/config"
	"github.com/muzyk0/online-quiz-game/internal/app/database"
	"github.com/muzyk0/online-quiz-game/internal/app/server"

	authhttp "github.com/muzyk0/online-quiz-game/internal/domain/auth/delivery/http"
	gamehttp "github.com/muzyk0/online-quiz-game/internal/domain/game/delivery/http"
	testinghttp "github.com/muzyk0/online-quiz-game/internal/domain/testing/delivery/http"
	gamerepo "github.com/muzyk0/online-quiz-game/internal/domain/game/repository"
	gameservice "github.com/muzyk0/online-quiz-game/internal/domain/game/service"
	healthhttp "github.com/muzyk0/online-quiz-game/internal/domain/health/delivery/http"
	questionhttp "github.com/muzyk0/online-quiz-game/internal/domain/question/delivery/http"
	questionrepo "github.com/muzyk0/online-quiz-game/internal/domain/question/repository"
	questionservice "github.com/muzyk0/online-quiz-game/internal/domain/question/service"
	userhttp "github.com/muzyk0/online-quiz-game/internal/domain/user/delivery/http"
	userrepo "github.com/muzyk0/online-quiz-game/internal/domain/user/repository"
	userservice "github.com/muzyk0/online-quiz-game/internal/domain/user/service"

	"github.com/muzyk0/online-quiz-game/internal/pkg/auth"
	"github.com/muzyk0/online-quiz-game/internal/pkg/logger"
	"github.com/muzyk0/online-quiz-game/internal/pkg/validation"

)

// App is the main application struct.
type App struct {
	cfg    *config.Config
	db     *database.DB
	server *server.Server

	tokenManager *auth.TokenManager

	healthHandler   *healthhttp.Handler
	userHandler     *userhttp.Handler
	authHandler     *authhttp.Handler
	questionHandler *questionhttp.Handler
	gameHandler     *gamehttp.Handler
	saUserHandler   *userhttp.SAHandler
	testingHandler  *testinghttp.TestingHandler
}

func New(cfg *config.Config) (*App, error) {
	logger.Initialize(cfg.ServerEnv)
	logger.Info("initializing application", "env", cfg.ServerEnv)

	a := &App{cfg: cfg}
	if err := a.initInfrastructure(); err != nil {
		return nil, fmt.Errorf("infrastructure init: %w", err)
	}
	a.initDomains()
	a.initServer()
	return a, nil
}

func (a *App) initInfrastructure() error {
	dbCtx, dbCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer dbCancel()
	db, err := database.New(dbCtx, a.cfg.DatabaseURL,
		a.cfg.DBMaxOpenConns, a.cfg.DBMaxIdleConns, a.cfg.DBConnMaxLifetimeMin)
	if err != nil {
		return fmt.Errorf("database connection: %w", err)
	}
	a.db = db

	accessExpiry := time.Duration(a.cfg.JWTAccessExpiryMin) * time.Minute
	refreshExpiry := time.Duration(a.cfg.JWTRefreshExpiryDays) * 24 * time.Hour
	a.tokenManager = auth.NewTokenManager(a.cfg.JWTSecret, accessExpiry, refreshExpiry)

	return nil
}

func (a *App) initDomains() {
	userRepo := userrepo.NewUserRepository(a.db)
	userSvc := userservice.NewUserService(userRepo)

	a.healthHandler = healthhttp.NewHandler(a.db)
	a.userHandler = userhttp.NewHandler(userSvc, a.tokenManager)
	a.saUserHandler = userhttp.NewSAHandler(userSvc)
	a.authHandler = authhttp.NewHandler(userSvc, a.tokenManager)

	questionRepository := questionrepo.NewQuestionRepository(a.db)
	questionSvc := questionservice.NewQuestionService(questionRepository)
	a.questionHandler = questionhttp.NewHandler(questionSvc)

	gameRepository := gamerepo.NewGameRepository(a.db)
	userLookup := gameservice.NewUserLookupAdapter(userRepo)
	gameSvc := gameservice.NewGameService(gameRepository, questionRepository, userLookup)
	a.gameHandler = gamehttp.NewHandler(gameSvc)

	a.testingHandler = testinghttp.NewHandler(a.db)
}

func (a *App) initServer() {
	a.server = server.New(a.cfg, validation.NewValidator())
	e := a.server.Echo

	authMiddleware := auth.JWTMiddleware(a.tokenManager)
	saMiddleware := auth.BasicAuthMiddleware(a.cfg.SAAdminLogin, a.cfg.SAAdminPassword)

	healthhttp.RegisterRoutes(e, a.healthHandler)
	userhttp.RegisterRoutes(e, a.userHandler, authMiddleware)
	authhttp.RegisterRoutes(e, a.authHandler, authMiddleware)

	userhttp.RegisterSARoutes(e, a.saUserHandler, saMiddleware)
	questionhttp.RegisterRoutes(e, a.questionHandler, saMiddleware)
	gamehttp.RegisterRoutes(e, a.gameHandler, authMiddleware)
	testinghttp.RegisterRoutes(e, a.testingHandler)
}

func (a *App) Run() error {
	port := fmt.Sprintf(":%d", a.cfg.ServerPort)
	log.Printf("Server is starting on port %s", port)

	serverErrors := make(chan error, 1)
	go func() {
		if err := a.server.Echo.Start(port); err != nil {
			serverErrors <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server failed to start: %w", err)
	case sig := <-stop:
		log.Printf("Received signal (%v), starting graceful shutdown...", sig)
		return a.Shutdown(context.Background())
	}
}

func (a *App) Shutdown(ctx context.Context) error {
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 10*time.Second)
	defer shutdownCancel()

	if err := a.server.Echo.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
		_ = a.server.Echo.Close()
	}

	if err := a.db.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}

	log.Println("Server stopped gracefully")
	return nil
}
