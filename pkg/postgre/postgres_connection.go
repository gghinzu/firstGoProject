package postgre

import (
	"database/sql"
	"firstGoProject/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Connection method returns a postgres connection using GORM
func Connection() *gorm.DB {
	configuration, errC := config.LoadConfig()
	if errC != nil {
		log.Fatal("cannot load config:", errC)
	}

	dsn := "host=" + configuration.PostgresHost + " user=" + configuration.PostgresUser + " password=" + configuration.PostgresPassword + " dbname=" + configuration.PostgresDB + " port=" + configuration.PostgresPort + " sslmode=disable"

	conn, err := sql.Open(configuration.DBDriver, dsn)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	gormDB, errG := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{})
	if errG != nil {
		panic(errG)
	}

	return gormDB
}
