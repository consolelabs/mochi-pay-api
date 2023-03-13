package activitylog

import "github.com/consolelabs/mochi-pay-api/internal/model"

type Store interface {
	CreateActivityLog(al *model.ActivityLog) error
}
