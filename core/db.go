package core

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/nuomizi-fw/stargazer/ent"
	"github.com/nuomizi-fw/stargazer/ent/hook"
	"github.com/pkg/errors"
	"modernc.org/sqlite"
)

type sqliteDriver struct {
	*sqlite.Driver
}

func (d sqliteDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})
	if _, err := c.Exec("PRAGMA foreign_keys = ON;", nil); err != nil {
		conn.Close()
		return nil, errors.Wrap(err, "failed to enable enable foreign keys")
	}
	return conn, nil
}

func init() {
	sql.Register("sqlite3", sqliteDriver{Driver: &sqlite.Driver{}})
}

type StargazerDB struct {
	*ent.Client
}

func NewStargazerDB(config StargazerConfig, logger StargazerLogger) StargazerDB {
	dsn := "file:" + config.Database.DBFile + "?cache=shared&_fk=1&_pragma=foreign_keys(1)"
	db, err := sql.Open(dialect.SQLite, dsn)
	if err != nil {
		logger.Error("Failed to open database", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(1)
	db.SetConnMaxIdleTime(time.Second * 1000)

	client := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.SQLite, db)))

	client.Use(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			start := time.Now()
			defer func() {
				logger.Infof("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
			}()
			return next.Mutate(ctx, m)
		})
	})

	return StargazerDB{client}
}
