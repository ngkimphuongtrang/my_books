package store

import (
	"gorm.io/gorm"
	"path"

	"github.com/ngkimphuongtrang/runkit/db"
	log "github.com/sirupsen/logrus"
)

const (
	dbMigrateDirectory = "/schema/migration/"
)

type SQLStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *SQLStore {
	return &SQLStore{db: db}
}

type DBStores struct {
	*SQLStore
	config *db.MySQLConfig

	BookStore *BookStore
	ReadStore *ReadStore
}

func NewDBStores(cfg *db.MySQLConfig) (*DBStores, error) {
	gormDB, err := db.ConnectORM(cfg)
	if err != nil {
		return nil, err
	}

	return &DBStores{
		config:    cfg,
		SQLStore:  NewSQLStore(gormDB),
		BookStore: NewBookStore(gormDB),
		ReadStore: NewReadStore(gormDB),
	}, nil
}

func (s *DBStores) Migrate(baseDir string) error {
	migrationPath := baseDir + dbMigrateDirectory
	migrate, err := db.NewMySQLMigrate(s.config, migrationPath)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}

	defer migrate.Close()

	if err := migrate.MigrateDB(); err != nil {
		log.Errorf("%v", err)
		return err
	}

	return nil
}

func (s *DBStores) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		log.Errorf("%v", err)
		return err
	}

	if err := sqlDB.Close(); err != nil {
		log.Errorf("%v", err)
		return err
	}

	return nil
}

func (s *DBStores) Reset(baseDir string) error {
	migrationPath := path.Join(baseDir, dbMigrateDirectory)
	migrate, err := db.NewMySQLMigrate(s.config, migrationPath)
	if err != nil {
		return err
	}
	defer migrate.Close()

	err = migrate.ResetDB()
	if err != nil {
		return err
	}

	return s.Migrate(baseDir)
}
