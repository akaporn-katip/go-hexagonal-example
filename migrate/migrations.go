package migrate

import (
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator interface {
	Up() error
	Down() error
}
