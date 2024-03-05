package db

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	log "github.com/sirupsen/logrus"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
	// load mysql
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	// blank import
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationPattern = `^(\d*)_\S+\.(up|down)\.sql$`

var migrationExpr = regexp.MustCompile(migrationPattern)

type MigrationType int

const (
	MigrationUp MigrationType = iota
	MigrationDown
)

type Migration struct {
	Name    string
	Version int
	Type    MigrationType
}

type MySQLMigrate struct {
	Config          *MySQLConfig
	Directory       string
	TotalMigrations uint
	migrate         *migrate.Migrate
}

func NewMySQLMigrate(config *MySQLConfig, migrationDirectory string) (*MySQLMigrate, error) {
	totalMigration, err := validateMigrationFolder(migrationDirectory)
	if err != nil {
		return nil, err
	}

	mySQLMigrate := &MySQLMigrate{
		Config:          config,
		Directory:       migrationDirectory,
		TotalMigrations: totalMigration,
	}

	err = mySQLMigrate.newMigrateObject()
	if err != nil {
		return nil, err
	}

	return mySQLMigrate, nil
}

func (m *MySQLMigrate) Close() {
	sourceErr, databaseErr := m.migrate.Close()
	if sourceErr != nil {
		log.Errorf("%v", sourceErr)
	}
	if databaseErr != nil {
		log.Errorf("%v", databaseErr)
	}
}

func (m *MySQLMigrate) MigrateDB() error {
	err := m.migrate.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	if errors.Is(err, migrate.ErrNoChange) {
		log.Errorf("no migration needed")
	}

	return nil
}

func validateMigrationFolder(path string) (uint, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	migrations := make([]*Migration, 0)
	for _, file := range files {
		migration, ok := parseMigrationFromFileName(file.Name())
		if !ok {
			return 0, fmt.Errorf("%s does not follow predefined format", file.Name())
		}
		migrations = append(migrations, migration)
	}

	if len(migrations) == 0 {
		return 0, fmt.Errorf("migration folder %s is empty", path)
	} else if len(migrations)%2 != 0 {
		return 0, fmt.Errorf("total migrations should be even")
	}

	sort.Slice(migrations, func(i, j int) bool {
		first := migrations[i]
		second := migrations[j]
		if first.Version == second.Version {
			return first.Type < second.Type
		}
		return first.Version < second.Version
	})

	for index, migration := range migrations {
		expectedVersion := index/2 + 1
		if index%2 == 0 {
			if migration.Type != MigrationUp {
				return 0, fmt.Errorf("no migration up for version %d", expectedVersion)
			}
			if migration.Version != expectedVersion {
				return 0, fmt.Errorf("no migration up for version %d", expectedVersion)
			}
		} else {
			if migration.Type != MigrationDown {
				return 0, fmt.Errorf("no migration down for version %d", expectedVersion)
			}
			if migration.Version != expectedVersion {
				return 0, fmt.Errorf("no migration down for version %d", expectedVersion)
			}
		}
	}

	return uint(len(migrations) / 2), nil
}

func parseMigrationFromFileName(name string) (*Migration, bool) {
	res := migrationExpr.FindAllStringSubmatch(name, -1)
	if res == nil {
		return nil, false
	}

	if len(res) != 1 || len(res[0]) != 3 {
		return nil, false
	}

	version, err := strconv.Atoi(res[0][1])
	if err != nil {
		return nil, false
	}

	if version <= 0 {
		return nil, false
	}

	var migrationType MigrationType
	if res[0][2] == "up" {
		migrationType = MigrationUp
	} else {
		migrationType = MigrationDown
	}

	return &Migration{
		Name:    name,
		Version: version,
		Type:    migrationType,
	}, true
}

func (m *MySQLMigrate) newMigrateObject() error {
	connectionString := fmt.Sprintf("mysql://%s", resolveDatabaseConnectionURL(m.Config))
	migrateObj, err := migrate.New(fmt.Sprintf("file://%s", m.Directory), connectionString)
	if err != nil {
		return err
	}
	m.migrate = migrateObj
	return nil
}

func resolveDatabaseConnectionURL(config *MySQLConfig) string {
	format := mysql.Config{
		User:                 config.User,
		Passwd:               config.Password,
		Addr:                 config.Server,
		Net:                  "tcp",
		DBName:               config.Schema,
		ParseTime:            true,
		MultiStatements:      true,
		Loc:                  time.Local,
		AllowNativePasswords: true,
	}
	return format.FormatDSN()
}
