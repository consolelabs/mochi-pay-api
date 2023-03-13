package db

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/consolelabs/mochi-pay-api/internal/config"
	activitylog "github.com/consolelabs/mochi-pay-api/internal/db/activity_log"
	"github.com/consolelabs/mochi-pay-api/internal/db/balance"
	"github.com/consolelabs/mochi-pay-api/internal/db/token"
)

type DB struct {
	Store       IStore
	Balance     balance.Store
	Token       token.Store
	ActivityLog activitylog.Store
}

func New(cfg config.View, logger *logrus.Entry) *DB {
	db := NewPostgresStore(cfg, logger)
	return &DB{
		Store:       db,
		Balance:     balance.New(db.DB()),
		Token:       token.New(db.DB()),
		ActivityLog: activitylog.New(db.DB()),
	}
}
