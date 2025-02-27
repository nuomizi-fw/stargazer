package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/gofiber/fiber/v2/log"
	"github.com/nuomizi-fw/stargazer/ent"
	"github.com/nuomizi-fw/stargazer/ent/hook"
	"github.com/nuomizi-fw/stargazer/pkg/config"
	"modernc.org/sqlite"
)

var db *ent.Client

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
		return nil, errors.New("failed to enable foreign keys")
	}
	return conn, nil
}

type StargazerDB struct {
	*ent.Client
}

func NewStargazerDB(config config.StargazerConfig) {
	sql.Register(dialect.SQLite, sqliteDriver{Driver: &sqlite.Driver{}})

	dsn := "file:" + config.Database.DBFile + "?cache=shared&_fk=1"
	sqlDB, err := sql.Open(dialect.SQLite, dsn)
	if err != nil {
		log.Error("Failed to open database", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetConnMaxIdleTime(time.Second * 1000)
	sqlDB.SetConnMaxLifetime(time.Second * 1000)

	db = ent.NewClient(ent.Driver(entsql.OpenDB(dialect.SQLite, sqlDB)))

	db.Use(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			start := time.Now()
			defer func() {
				log.Infof("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
			}()
			return next.Mutate(ctx, m)
		})
	})
}

func CloseStargazerDB() {
	err := db.Close()
	if err != nil {
		log.Error("Failed to close database", err)
	}
}

func AutoMigrate() {
	err := db.Schema.Create(context.Background())
	if err != nil {
		log.Error("Failed to migrate database", err)
	}
}
