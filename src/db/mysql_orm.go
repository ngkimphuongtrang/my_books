package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func ConnectORM(config *MySQLConfig) (*gorm.DB, error) {
	option := config.Option
	if len(option) == 0 {
		option = mysqlOption
	}

	connectionLifetimeSeconds := config.ConnectionLifetimeSeconds
	if connectionLifetimeSeconds == 0 {
		connectionLifetimeSeconds = defaultMySQLConnectionLifetimeSeconds
	}

	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", config.User, config.Password, config.Server, config.Schema, option)

	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
	db, err := gorm.Open(mysql.Open(source), cfg)
	if err != nil {
		log.Errorf("cannot connect to database %s", config.Schema)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("cannot obtain sql database object %s", config.Schema)
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(time.Duration(connectionLifetimeSeconds) * time.Second)
	sqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(config.MaxOpenConnections)

	log.Println("connected to database", config.Schema)
	return db, nil
}
