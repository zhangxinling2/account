package accounts

import (
	"github.com/shopspring/decimal"
	services "github.com/zhangxinling2/account/services"
	"time"
)

type AccountLog struct {
	Id              int64               `db:"id,omitempty"`         //
	LogNo           string              `db:"log_no,uni"`           //流水编号 全局不重复字符或数字，唯一性标识
	TradeNo         string              `db:"trade_no"`             //交易单号 全局不重复字符或数字，唯一性标识
	AccountNo       string              `db:"account_no"`           //账户编号 账户ID
	UserId          string              `db:"user_id"`              //用户编号
	Username        string              `db:"username"`             //用户名称
	TargetAccountNo string              `db:"target_account_no"`    //账户编号 账户ID
	TargetUserId    string              `db:"target_user_id"`       //目标用户编号
	TargetUsername  string              `db:"target_username"`      //目标用户名称
	Amount          decimal.Decimal     `db:"amount"`               //交易金额,该交易涉及的金额
	Balance         decimal.Decimal     `db:"balance"`              //交易后余额,该交易后的余额
	ChangeType      services.ChangeType `db:"change_type"`          //流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义
	ChangeFlag      services.ChangeFlag `db:"change_flag"`          //交易变化标识：-1 出账 1为进账，枚举
	Status          int                 `db:"status"`               //交易状态：
	Decs            string              `db:"decs"`                 //交易描述
	CreatedAt       time.Time           `db:"created_at,omitempty"` //创建时间
}

func (po *AccountLog) ToDTO() *services.AccountLogDTO {
	dto := &services.AccountLogDTO{
		TradeNo:         po.TradeNo,
		LogNo:           po.LogNo,
		AccountNo:       po.AccountNo,
		TargetAccountNo: po.TargetAccountNo,
		UserId:          po.UserId,
		UserName:        po.Username,
		TargetUserId:    po.TargetUserId,
		TargetUserName:  po.TargetUsername,
		Amount:          po.Amount,
		Balance:         po.Balance,
		ChangeType:      po.ChangeType,
		ChangeFlag:      po.ChangeFlag,
		Status:          po.Status,
		Decs:            po.Decs,
		CreatedAt:       po.CreatedAt,
	}
	return dto
}
func (po *AccountLog) FromTransferDTO(dto services.AccountTransferDTO) {
	po.TradeNo = dto.TradeNo
	po.AccountNo = dto.TradeBody.AccountNo
	po.TargetAccountNo = dto.TradeTarget.AccountNo
	po.UserId = dto.TradeBody.UserId
	po.Username = dto.TradeBody.UserName
	po.TargetUserId = dto.TradeTarget.UserId
	po.TargetUsername = dto.TradeTarget.UserName
	po.Amount = dto.Amount
	po.ChangeType = dto.ChangeType
	po.ChangeFlag = dto.ChangeFlag
	po.Decs = dto.Decs
}
