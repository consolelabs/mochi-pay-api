package db

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/consolelabs/mochi-pay-api/internal/config"
	"github.com/consolelabs/mochi-pay-api/internal/db/balance"
	"github.com/consolelabs/mochi-pay-api/internal/db/token"
	transferlog "github.com/consolelabs/mochi-pay-api/internal/db/transfer_log"
)

type DB struct {
	Store       IStore
	Balance     balance.Store
	Token       token.Store
	TransferLog transferlog.Store
}

func New(cfg config.View, logger *logrus.Entry) *DB {
	db := NewPostgresStore(cfg, logger)
	return &DB{
		Store:       db,
		Balance:     balance.New(db.DB()),
		Token:       token.New(db.DB()),
		TransferLog: transferlog.New(db.DB()),
	}
}
