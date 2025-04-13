package db

import (
	"fmt"
	"log"
	"sync"

	"github.com/MatheusTimmers/backend-test/internal/infra/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnOnce sync.Once
var dbConn *gorm.DB

func Db() *gorm.DB {
	if dbConn == nil {
		dbConnOnce.Do(
			func() {
				dbConn = openDbConnection()
			},
		)
	}

	return dbConn
}

func openDbConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Config().Database.Host,
		config.Config().Database.User,
		config.Config().Database.Password,
		config.Config().Database.Name,
		config.Config().Database.Port,
		config.Config().Database.SslMode,
	)

	dialect := postgres.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{TranslateError: true})
	if err != nil {
		log.Panic("Erro: Failed to connect to database:", err)
	}

	return db
}
