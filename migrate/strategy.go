package migrate

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
)

type SQLMigrator struct {
	migrator *migrate.Migrate
}

func (s *SQLMigrator) Up() error {
	return s.migrator.Up()
}

func (s *SQLMigrator) Down() error {
	return s.migrator.Down()
}

func NewSqlMigrator(dsn string) (*SQLMigrator, error) {
	mt, err := migrate.New("file://migrate/sql", dsn)

	if err != nil {
		return nil, err
	}

	return &SQLMigrator{
		migrator: mt,
	}, nil
}

type MongoMigrator struct {
	migrator *migrate.Migrate
}

func (s *MongoMigrator) Up() error {
	return s.migrator.Up()
}

func (s *MongoMigrator) Down() error {
	return s.migrator.Down()
}

func NewMongoMigrator(dsn string) (*SQLMigrator, error) {
	mt, err := migrate.New("file://migrate/mongo", dsn)

	if err != nil {
		return nil, err
	}

	return &SQLMigrator{
		migrator: mt,
	}, nil
}

func NewMigrator(dbType string, dsn string) (Migrator, error) {
	switch dbType {
	case "postgres", "sqlite3":
		return NewSqlMigrator(dsn)
	case "mongodb":
		return NewMongoMigrator(dsn)
	default:
		return nil, fmt.Errorf("unsupported driver")
	}
}
