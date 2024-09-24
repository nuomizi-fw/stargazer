package core

import (
	"context"
	"database/sql"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nuomizi-fw/stargazer/ent"
	"github.com/nuomizi-fw/stargazer/ent/hook"
)

type StargazerDB struct {
	*ent.Client
}

func NewStargazerDB(config StargazerConfig, logger StargazerLogger) StargazerDB {
	dsn := "file:" + config.Database.DBFile + "?cache=shared&_fk=1"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		logger.Error("Failed to open database", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(1)
	db.SetConnMaxIdleTime(time.Second * 1000)

	client := ent.NewClient(ent.Driver(entsql.OpenDB("sqlite3", db)))

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
