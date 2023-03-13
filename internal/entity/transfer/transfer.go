package nftsaleentity

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/consolelabs/mochi-pay-api/internal/apperror/apierror"
	"github.com/consolelabs/mochi-pay-api/internal/appmain"
	"github.com/consolelabs/mochi-pay-api/internal/db"
	"github.com/consolelabs/mochi-pay-api/internal/model"
)

type transferEntity struct {
	db     *db.Store
	params *appmain.Params
}

func New(p *appmain.Params, db *db.Store) ITransfer {
	return &transferEntity{
		db:     db,
		params: p,
	}
}

func (t *transferEntity) TransferToken(req *model.TransferRequest) error {
	listReceivers := make([]string, 0)
	for _, to := range req.Tos {
		listReceivers = append(listReceivers, to.ProfileGlobalId)
	}

	totalTransferAmount := 0.0
	for _, amount := range req.Amount {
		amount, _ := strconv.ParseFloat(amount, 64)
		totalTransferAmount += amount
	}

	log := &model.ActivityLog{
		ProfileId:      req.From.ProfileGlobalId,
		Receiver:       listReceivers,
		NumberReceiver: int64(len(listReceivers)),
		Amount:         totalTransferAmount,
		Status:         "failed",
		Note:           apierror.ErrTokenNotSupport.Error(),
	}

	// check if token existed
	token, err := t.db.Token.GetBySymbol(req.Token.Symbol)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			t.db.ActivityLog.CreateActivityLog(log)
			return apierror.ErrTokenNotSupport
		}
		t.params.Logger().WithFields(logrus.Fields{"req": req}).Error(err, "[transferEntity.TransferToken] - failed to get token")
		return err
	}

	// check sender's balance
	senderBalance, err := t.db.Balance.GetBalanceByTokenID(req.From.ProfileGlobalId, token.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			insufficientLog := log
			insufficientLog.Note = apierror.ErrInsufficientBalance.Error()
			insufficientLog.TokenId = token.Id
			t.db.ActivityLog.CreateActivityLog(log)
			return apierror.ErrInsufficientBalance
		}
		t.params.Logger().WithFields(logrus.Fields{"req": req}).Error(err, "[transferEntity.TransferToken] - failed to get sender balance")
		return err
	}

	if float64(senderBalance.Amount) < totalTransferAmount {
		insufficientLog := log
		insufficientLog.Note = apierror.ErrInsufficientBalance.Error()
		insufficientLog.TokenId = token.Id
		t.db.ActivityLog.CreateActivityLog(log)
		return apierror.ErrInsufficientBalance
	}

	// execute transfer
	batch := []model.Balance{{ProfileId: req.From.ProfileGlobalId, TokenId: token.Id, ChangedAmount: -totalTransferAmount}}
	for idx, r := range req.Tos {
		recipientAmount, _ := strconv.ParseFloat(req.Amount[idx], 64)
		batch = append(batch, model.Balance{
			ProfileId:     r.ProfileGlobalId,
			TokenId:       token.Id,
			ChangedAmount: recipientAmount,
			Amount:        recipientAmount,
		})
	}

	err = t.db.Balance.UpsertBatch(batch)
	if err != nil {
		t.params.Logger().WithFields(logrus.Fields{"req": req}).Error(err, "[transferEntity.TransferToken] - failed to upsert balance")
		return err
	}
	log.Status = "success"
	log.Note = "transfer success"
	err = t.db.ActivityLog.CreateActivityLog(log)
	if err != nil {
		t.params.Logger().WithFields(logrus.Fields{"req": req}).Error(err, "[transferEntity.TransferToken] - failed to create activity log")
		return err
	}

	return nil
}
