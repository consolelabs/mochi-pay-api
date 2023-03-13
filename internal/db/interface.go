package db

import (
	"gorm.io/gorm"
)

type IFinalFunc interface {
	Commit() error
	Rollback(err error) error
}

// IStore persistent data interface
type IStore interface {
	DB() *gorm.DB
	NewTx() (IStore, IFinalFunc)
	Shutdown() error
}
