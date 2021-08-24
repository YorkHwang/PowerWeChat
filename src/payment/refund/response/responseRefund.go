package response

import (
	"github.com/ArtisanCloud/go-libs/object"
)

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_9.shtml

type ResponseRefund struct {
	RefundID            string          `json:"refund_id"`
	OutRefundNO         string          `json:"out_refund_no"`
	TransactionID       string          `json:"transaction_id"`
	OutTradeNO          string          `json:"out_trade_no"`
	Channel             string          `json:"channel"`
	UserReceivedAccount string          `json:"user_received_account"`
	SuccessTime         string          `json:"success_time"`
	CreateTime          string          `json:"create_time"`
	Status              string          `json:"status"`
	FundsAccount        string          `json:"funds_account"`
	Amount              *object.HashMap `json:"amount"`
	PromotionDetail     *object.HashMap `json:"promotion_detail"`
}