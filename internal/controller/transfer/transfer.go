package transfer

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/consolelabs/mochi-pay-api/internal/apperror/apierror"
	"github.com/consolelabs/mochi-pay-api/internal/appmain"
	"github.com/consolelabs/mochi-pay-api/internal/model"
	"github.com/consolelabs/mochi-pay-api/internal/utils"
)

type transfer struct {
	params *appmain.Params
}

func New(p *appmain.Params) ITransfer {
	return &transfer{
		params: p,
	}
}

func (t *transfer) TransferToken(req *model.TransferRequest) *apierror.ApiError {
	listReceivers := make([]string, 0)
	for _, to := range req.Tos {
		listReceivers = append(listReceivers, to.ProfileGlobalId)
	}

	log := &model.TransferLog{
		SenderProfileId:     req.From.ProfileGlobalId,
		RecipientsProfileId: listReceivers,
		NumberReceiver:      int64(len(listReceivers)),
		Status:              "failed",
		Note:                apierror.ErrTokenNotSupport.Error(),
	}

	// check if token existed
	token, err := t.params.DB().Token.GetById(req.TokenId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			t.params.DB().TransferLog.CreateTransferLog(log)
			return apierror.ErrTokenNotSupport
		}
		t.params.Logger().WithFields(logrus.Fields{"req": req}).Error(err, "[transferController.TransferToken] - failed to get token")
		return apierror.New(err.Error(), 500, apierror.Code500)
	}

	totalTransferAmount := 0.0
	for _, amount := range req.Amount {
		convertedAmountBigFloat := utils.ConvertBigIntString(amount, token)
		convertedAmount, _ := convertedAmountBigFloat.Float64()
		totalTransferAmount += convertedAmount
	}

	log.Amount = totalTransferAmount

	// check sender's balance
	senderBalance, err := t.params.DB().Balance.GetBalanceByTokenID(req.From.ProfileGlobalId, token.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			insufficientLog := log
			insufficientLog.Note = apierror.ErrInsufficientBalance.Error()
			insufficientLog.TokenId = token.Id
			t.params.DB().TransferLog.CreateTransferLog(log)
			return apierror.ErrInsufficientBalance
		}
		t.params.Logger().WithFields(logrus.Fields{"req": req}).Error(err, "[transferController.TransferToken] - failed to get sender balance")
		return apierror.New(err.Error(), 500, apierror.Code500)
	}

	if float64(senderBalance.Amount) < totalTransferAmount {
		insufficientLog := log
		insufficientLog.Note = apierror.ErrInsufficientBalance.Error()
		insufficientLog.TokenId = token.Id
		t.params.DB().TransferLog.CreateTransferLog(log)
		return apierror.ErrInsufficientBalance
	}

	// execute transfer
	batch := []model.Balance{{ProfileId: req.From.ProfileGlobalId, TokenId: token.Id, ChangedAmount: -totalTransferAmount}}
	for idx, r := range req.Tos {
		recipientAmount, _ := utils.ConvertBigIntString(req.Amount[idx], token).Float64()
		batch = append(batch, model.Balance{
			ProfileId:     r.ProfileGlobalId,
			TokenId:       token.Id,
			ChangedAmount: recipientAmount,
			Amount:        recipientAmount,
		})
	}

	tx, fn := t.params.DB().Store.NewTx()
	for _, item := range batch {
		tx.DB().Clauses(
			clause.OnConflict{
				Columns:   []clause.Column{{Name: "token_id"}, {Name: "profile_id"}},
				DoUpdates: clause.Assignments(map[string]interface{}{"amount": gorm.Expr("balances.amount::decimal + ?", item.ChangedAmount)}),
			},
			clause.Returning{Columns: []clause.Column{{Name: "amount"}}},
		).Create(&item)
	}
	err = fn.Commit()
	if err != nil {
		t.params.Logger().WithFields(logrus.Fields{"req": req}).Error(err, "[transferController.TransferToken] - failed to commit tx")
		return apierror.New(err.Error(), 500, apierror.Code500)
	}

	log.Status = "success"
	log.Note = "transfer success"
	err = t.params.DB().TransferLog.CreateTransferLog(log)
	if err != nil {
		t.params.Logger().WithFields(logrus.Fields{"req": req}).Error(err, "[transferController.TransferToken] - failed to create transfer log")
		return apierror.New(err.Error(), 500, apierror.Code500)
	}

	return nil
}
