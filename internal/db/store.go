package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/consolelabs/mochi-pay-api/internal/appmain"
	activitylog "github.com/consolelabs/mochi-pay-api/internal/db/activity_log"
	"github.com/consolelabs/mochi-pay-api/internal/db/balance"
	"github.com/consolelabs/mochi-pay-api/internal/db/token"
)

type Store struct {
	DB          *gorm.DB
	Token       token.IToken
	Balance     balance.IBalance
	ActivityLog activitylog.IActivityLog
}

func New(p *appmain.Params) *Store {
	db := connDb(p)
	return &Store{
		DB:          db,
		Token:       token.New(db),
		Balance:     balance.New(db),
		ActivityLog: activitylog.New(db),
	}
}

func connDb(p *appmain.Params) *gorm.DB {
	ds := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.Config().GetString("DB_USER"), p.Config().GetString("DB_PASS"),
		p.Config().GetString("DB_HOST"), p.Config().GetString("DB_PORT"), p.Config().GetString("DB_NAME"),
	)

	conn, err := sql.Open("postgres", ds)
	if err != nil {
		p.Logger().Fatalf("failed to open database connection", err)
	}

	db, err := gorm.Open(postgres.New(
		postgres.Config{Conn: conn}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	if err != nil {
		p.Logger().Fatalf("failed to open database connection", err)
	}

	p.Logger().Info("database connected")

	return db
}
