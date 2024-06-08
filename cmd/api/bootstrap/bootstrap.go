package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lorenzoMrt/ContentInsight/internal/creating"
	"github.com/lorenzoMrt/ContentInsight/internal/platform/server"
	"github.com/lorenzoMrt/ContentInsight/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "contentInsight"
	dbPass = "contentInsight123"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "contents"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	contentRepository := mysql.NewContentRepository(db)

	createContentService := creating.NewContentService(contentRepository)

	srv := server.New(host, port, createContentService)
	return srv.Run()
}
