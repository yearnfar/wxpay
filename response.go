package wxpay

import (
	"encoding/xml"
)

// 基础返回数据
type BaseResponse struct {
	XmlName    xml.Name `xml:"xml"`
	ReturnCode string   `xml:"return_code"`
	ReturnMsg  string   `xml:"return_msg"`
	ResultCode string   `xml:"result_code"`  // 业务结果
	ErrCode    string   `xml:"err_code"`     // 错误码
	ErrCodeDes string   `xml:"err_code_des"` // 错误描述
	AppId      string   `xml:"appid"`        // 公众账号ID
	MchId      string   `xml:"mch_id"`       // 商户号
	DeviceInfo string   `xml:"device_info"`  // 设备号
	NonceStr   string   `xml:"nonce_str"`    // 随机字符串
	Sign       string   `xml:"sign"`         // 签名
}

// 统一下单
type UnifiedOrderResponse struct {
	BaseResponse `xml:",innerXml"`

	TradeType string `xml:"trade_type"` // 交易类型
	PrepayId  string `xml:"prepay_id"`  // 预支付交易会话标识
	CodeUrl   string `xml:"code_url"`   // 二维码链接
}

// 统一下单
type RefundResponse struct {
	BaseResponse `xml:",innerXml"`

	TransactionId     string `xml:"transaction_id"`      // 微信订单号
	OutTradeNo        string `xml:"out_trade_no"`        // 商户订单号
	OutRefundNo       string `xml:"out_refund_no"`       // 商户退款单号
	RefundId          string `xml:"refund_id"`           // 微信退款单号
	RefundChannel     string `xml:"refund_channel"`      // 退款渠道
	RefundFee         string `xml:"refund_fee"`          // 退款金额
	TotalFee          string `xml:"total_fee"`           // 订单总金额
	FeeType           string `xml:"fee_type"`            // 订单金额货币种类
	CashFee           string `xml:"cash_fee"`            // 现金支付金额
	CashRefundFee     string `xml:"cash_refund_fee"`     // 现金退款金额
	CouponRefundFee   string `xml:"coupon_refund_fee"`   // 代金券或立减优惠退款金额
	CouponRefundCount string `xml:"coupon_refund_count"` // 代金券或立减优惠使用数量
	CouponRefundId    string `xml:"coupon_refund_id"`    // 代金券或立减优惠ID
}

// 查询订单
type OrderQueryResponse struct {
	BaseResponse `xml:",innerXml"`

	OpenId         string `xml:"openid"`           // 用户标识
	IsSubscribe    string `xml:"is_subscribe"`     // 是否关注公众账号
	TradeType      string `xml:"trade_type"`       // 交易类型
	TradeState     string `xml:"trade_state"`      // 交易状态
	BankType       string `xml:"bank_type"`        // 付款银行
	TotalFee       string `xml:"total_fee"`        // 总金额
	FeeType        string `xml:"fee_type"`         // 货币种类
	CashFee        string `xml:"cash_fee"`         // 现金支付金额
	CashFeeType    string `xml:"cash_fee_type"`    // 现金支付货币类型
	CouponFee      string `xml:"coupon_fee"`       // 代金券或立减优惠金额
	CouponCount    string `xml:"coupon_count"`     // 代金券或立减优惠使用数量
	TransactionId  string `xml:"transaction_id"`   // 微信支付订单号
	OutTradeNo     string `xml:"out_trade_no"`     // 商户订单号
	Attach         string `xml:"attach"`           // 附加数据
	TimeEnd        string `xml:"time_end"`         // 支付完成时间
	TradeStateDesc string `xml:"trade_state_desc"` // 交易状态描述
}

// 关闭订单
type CloseOrderResponse struct {
	BaseResponse `xml:",innerXml"`
}

// 查询退款
type RefundQueryResponse struct {
	BaseResponse `xml:",innerXml"`

	TransactionId string `xml:"transaction_id"` // 微信订单号
	OutTradeNo    string `xml:"out_trade_no"`   // 商户订单号
	TotalFee      string `xml:"total_fee"`      // 订单总金额
	FeeType       string `xml:"fee_type"`       // 订单金额货币种类
	CashFee       string `xml:"cash_fee"`       // 现金支付金额
	RefundCount   string `xml:"refund_count"`   // 退款笔数
}

// 回调
type NotifyResponse struct {
	BaseResponse `xml:",innerXml"`

	OpenId        string `xml:"openid"`         // 用户标识
	IsSubscribe   string `xml:"is_subscribe"`   // 是否关注公众账号
	TradeType     string `xml:"trade_type"`     // 交易类型
	BankType      string `xml:"bank_type"`      // 付款银行
	TotalFee      string `xml:"total_fee"`      // 总金额
	FeeType       string `xml:"fee_type"`       // 货币种类
	CashFee       string `xml:"cash_fee"`       // 现金支付金额
	CashFeeType   string `xml:"cash_fee_type"`  // 现金支付货币类型
	CouponFee     string `xml:"coupon_fee"`     // 代金券或立减优惠金额
	CouponCount   string `xml:"coupon_count"`   // 代金券或立减优惠使用数量
	TransactionId string `xml:"transaction_id"` // 微信支付订单号
	OutTradeNo    string `xml:"out_trade_no"`   // 商户订单号
	Attach        string `xml:"attach"`         // 商家数据包，原样返回
	TimeEnd       string `xml:"time_end"`       // 支付完成时间
}
