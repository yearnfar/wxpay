package wxpay

import (
	"errors"
)

// UnifiedOrderParam 参数
type UnifiedOrderParam struct {
	Param `xml:",innerXml"`

	// Body 商品简单描述，该字段请按照规范传递，具体请见参数规定
	Body string `xml:"body"`

	// Detail 商品详细描述，对于使用单品优惠的商户，改字段必须按照规范上传，详见“单品优惠参数说明”
	Detail string `xml:"detail"`

	// Attach 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	Attach string `xml:"attach"`

	// OutTradeNo 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|* 且在同一个商户号下唯一。详见商户订单号
	OutTradeNo string `xml:"out_trade_no"`

	// FeeType 符合ISO 4217标准的三位字母代码，默认人民币：CNY，详细列表请参见货币类型
	FeeType string `xml:"fee_type"`

	// TotalFee 订单总金额，单位为分，详见支付金额
	TotalFee int `xml:"total_fee"` // 订单总金额，单位为分，详见支付金额

	// SpbillCreateIP APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。
	SpbillCreateIP string `xml:"spbill_create_ip"`

	// TimeStart 订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	TimeStart string `xml:"time_start"`

	// TimeExpire 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。
	// 订单失效时间是针对订单号而言的，由于在请求支付的时候有一个必传参数prepay_id只有两小时的有效期，所以在重入时间超过2小时的时候需要重新请求下单接口获取新的prepay_id。
	// 其他详见时间规则 建议：最短失效时间间隔大于1分钟
	TimeExpire string `xml:"time_expire"`

	// GoodsTag 订单优惠标记，使用代金券或立减优惠功能时需要的参数，说明详见代金券或立减优惠
	GoodsTag string `xml:"goods_tag"`

	// NotifyURL 异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。
	NotifyURL string `xml:"notify_url"`

	// TradeType 说明详见参数规定
	// JSAPI 公众号支付 NATIVE 扫码支付 APP APP支付
	TradeType string `xml:"trade_type"`

	// ProductID trade_type=NATIVE时（即扫码支付），此参数必传。此参数为二维码中包含的商品ID，商户自行定义。
	ProductID string `xml:"product_id"`

	// LimitPay 上传此参数no_credit--可限制用户不能使用信用卡支付
	LimitPay string `xml:"limit_pay"`

	// OpenID trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识。
	// openid如何获取，可参考【获取openid】。 企业号请使用【企业号OAuth2.0接口】获取企业号内成员userid，再调用【企业号userid转openid接口】进行转换
	OpenID string `xml:"openid"`

	// SceneInfo 该字段用于上报场景信息，目前支持上报实际门店信息。
	// 该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }} ，字段详细说明请点击行前的展开
	SceneInfo string `xml:"scene_info"`
}

// SceneInfo 场景信息
type SceneInfo struct {
	// ID 门店唯一标识
	ID string `json:"id"`

	// Name 门店名称
	Name string `json:"name"`

	// AreaCode 门店所在地行政区划码，详细见《最新县及县以上行政区划代码》
	AreaCode string `json:"areaCode"`

	// Address 门店详细地址
	Address string `json:"address"`
}

// Valid 验证
func (p *UnifiedOrderParam) Valid() error {
	err := p.Param.Valid()
	if err != nil {
		return err
	}

	if p.Body == "" {
		return errors.New("商品描述不能为空")
	}

	if p.OutTradeNo == "" {
		return errors.New("商户订单号不能为空")
	}

	if p.TotalFee == 0 {
		return errors.New("标价金额不能为空")
	}

	if p.SpbillCreateIP == "" {
		return errors.New("终端IP不能为空")
	}

	if p.NotifyURL == "" {
		return errors.New("通知地址不能为空")
	}

	if p.TradeType == "" {
		return errors.New("交易类型不能为空")
	}

	if p.TradeType == TradeTypeJSAPI && p.OpenID == "" {
		return errors.New("openId不能为空")
	}

	return nil
}

// UnifiedOrderResponse 统一下单返回
type UnifiedOrderResponse struct {
	Response `xml:",innerXml"`

	TradeType string `xml:"trade_type"` // 交易类型
	PrepayID  string `xml:"prepay_id"`  // 预支付交易会话标识
	CodeURL   string `xml:"code_url"`   // 二维码链接
}

// NotifyResponse 回调数据
type NotifyResponse struct {
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

// IsPaid 是否支付成功
func (r *NotifyResponse) IsPaid() bool {
	return r.TradeState == TradeStateSuccess
}
