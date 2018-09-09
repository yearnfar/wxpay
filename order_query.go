package wxpay

import "errors"

// OrderQueryParam 查询订单参数
type OrderQueryParam struct {
	Param `xml:",innerXml"`

	// TransactionID 微信的订单号，建议优先使用
	TransactionID string `xml:"transaction_id"`

	// OutTradeNo 商户订单号
	OutTradeNo string `xml:"out_trade_no"`
}

// Valid 参数校验
func (p *OrderQueryParam) Valid() error {
	if p.TransactionID == "" && p.OutTradeNo == "" {
		return errors.New("微信订单号、商户订单号不能同时为空！")
	}

	return nil
}

// OrderQueryParam 查询订单返回数据结构
type OrderQueryResponse struct {
	Response `xml:",innerXml"`

	// OpenID 用户在商户appid下的唯一标识
	OpenID string `xml:"openid"`

	// IsSubscribe 用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	IsSubscribe string `xml:"is_subscribe,omitempty"`

	// TradeType 调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，MICROPAY，详细说明见参数规定
	TradeType string `xml:"trade_type"`

	// TradeState 交易状态 支付状态机请见下单API页面
	TradeState string `xml:"trade_state"`

	// BankType 银行类型，采用字符串类型的银行标识
	BankType string `xml:"bank_type"`

	// TotalFee 标价金额
	TotalFee int `xml:"total_fee"`

	// SettlementTotalFee 当订单使用了免充值型优惠券后返回该参数，应结订单金额=订单金额-免充值优惠券金额。
	SettlementTotalFee int `xml:"settlement_total_fee,omitempty"`

	// FeeType 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	FeeType string `xml:"fee_type,omitempty"`

	// CashFee 现金支付金额订单现金支付金额，详见支付金额
	CashFee int `xml:"cash_fee"`

	// CashFeeType 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFeeType string `xml:"cash_fee_type,omitempty"`

	// CouponFee “代金券”金额<=订单金额，订单金额-“代金券”金额=现金支付金额，详见支付金额
	CouponFee int `xml:"coupon_fee"`

	// CouponCount 代金券使用数量
	CouponCount int `xml:"coupon_count"`

	// TransactionID 微信支付订单号
	TransactionID string `xml:"transaction_id"`

	// OutTradeNo 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	OutTradeNO string `xml:"out_trade_no"`

	// Attach 附加数据，原样返回
	Attach string `xml:"attach"`

	// TimeEnd 订单支付时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	TimeEnd string `xml:"time_end"`

	// TradeStateDesc 对当前查询订单状态的描述和下一步操作的指引
	TradeStateDesc string `xml:"trade_state_desc"`
}
