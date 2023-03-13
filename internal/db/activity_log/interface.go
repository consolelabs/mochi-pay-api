package activitylog

import "github.com/consolelabs/mochi-pay-api/internal/model"

type IActivityLog interface {
	CreateActivityLog(al *model.ActivityLog) error
}
