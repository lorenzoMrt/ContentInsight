package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	cr "github.com/lorenzoMrt/ContentInsight/internal"
	"github.com/lorenzoMrt/ContentInsight/internal/creating"
	"github.com/lorenzoMrt/ContentInsight/internal/increasing"
	"github.com/lorenzoMrt/ContentInsight/internal/platform/bus/inmemory"
	"github.com/lorenzoMrt/ContentInsight/internal/platform/server"
	"github.com/lorenzoMrt/ContentInsight/internal/platform/storage/mysql"
)

const (
	host            = "localhost"
	port            = 8080
	shutdownTimeout = 10 * time.Second

	dbUser = "contentInsight"
	dbPass = "contentInsight123"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "contents"
)

func Run() error {
	var cfg config
	err := envconfig.Process("CR", &cfg)
	if err != nil {
		return err
	}
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	contentRepository := mysql.NewContentRepository(db)

	createContentService := creating.NewContentService(contentRepository, eventBus)
	increasingContentCounterService := increasing.NewContentCounterService()

	createContentCommandHandler := creating.NewContentCommandHandler(createContentService)
	commandBus.Register(creating.ContentCommandType, createContentCommandHandler)

	eventBus.Subscribe(
		cr.ContentCreatedEventType,
		creating.NewIncreaseContentsCounterOnContentCreated(increasingContentCounterService),
	)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}

type config struct {
	Host            string        `default: "localhost"`
	Port            uint          `default: "8080"`
	ShutdownTimeout time.Duration `default: "10s"`
	DbUser          string        `default: "contentInsight"`
	DbPass          string        `default: "contentInsight123"`
	DbHost          string        `default: "localhost"`
	DbPort          string        `default: "3306"`
	DbName          string        `default: "contents"`
	DbTimeout       time.Duration `default: "5s"`
}
