package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-test-task/internal/register_layers"
	repositoryStorage "github.com/golang-test-task/internal/repository/storage"
	"github.com/golang-test-task/internal/server"
	"github.com/golang-test-task/internal/server/route"
	"github.com/golang-test-task/pkg/common/core/clock"
	"github.com/golang-test-task/pkg/common/logging"
	"github.com/golang-test-task/pkg/core/identity"
	"github.com/golang-test-task/pkg/databases/postgresql"
	"github.com/golang-test-task/pkg/service/email"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"strconv"
	"time"
)

type App struct {
	ctx    context.Context
	server *gin.Engine
}

func NewApp(ctx context.Context) *App {
	return &App{
		ctx: ctx,
	}
}

func (a *App) Start() error {

	db, err := a.initDB()
	if err != nil {
		logging.L(a.ctx).Fatal(errors.Errorf("failed to connect to postgres: %s", err).Error())
		return err
	}

	defer db.Close()

	gd := a.initGlobalDelivery(db)
	r := a.initServer(gd)

	logging.L(a.ctx).Info("starting server")

	return r.Run(os.Getenv("SERVER_URL"))
}

func (a *App) initGlobalDelivery(db *pgxpool.Pool) *register_layers.GlobalDelivery {

	storage := repositoryStorage.NewStorage(db)

	gr := register_layers.NewGlobalRepository(storage)

	generator := identity.NewGenerator()
	c := clock.New()

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	emailSender := email.NewEmailService(os.Getenv("SMTP_SENDER"), os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))

	logging.L(a.ctx).Info(fmt.Sprintf("SMTP emailSender: %v", emailSender))
	gs := register_layers.NewGlobalService(gr, generator, c, emailSender)
	gd := register_layers.NewGlobalDelivery(gs)

	return gd
}

func (a *App) initDB() (pool *pgxpool.Pool, err error) {
	pgConfig := postgresql.PgConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	db, err := postgresql.NewClient(a.ctx, 10, 5*time.Second, pgConfig, true)
	if err != nil {
		logging.L(a.ctx).Fatal(errors.Errorf("failed to connect to postgres: %s", err).Error())
	}

	return db, nil
}

func (a *App) initServer(gd *register_layers.GlobalDelivery) *gin.Engine {
	// Create a new gin router.
	r := server.New()
	// Set delimiters for templates.
	r.Delims("{[{", "}]}")

	// Create public API routes.
	public := r.Group("/api/v1")
	// Define routes for authentication, users, friends, posts, and commentaries.
	route.Students(public, *gd.StudentDelivery)
	route.Tasks(public, *gd.TaskDelivery)
	route.Job(public, *gd.JobDelivery)

	// Serve Swagger UI.
	public.GET("/swagger", server.Redirect)
	public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
