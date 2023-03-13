package transferlog

import "github.com/consolelabs/mochi-pay-api/internal/model"

type Store interface {
	CreateTransferLog(al *model.TransferLog) error
}
