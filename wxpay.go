package wxpay

import (
	"encoding/xml"
	"errors"
)

const (
	WXPAY_UNIFIEDORDER_URL = "https://api.mch.weixin.qq.com/pay/unifiedorder" // 统一下单
	WXPAY_ORDERQUERY_URL   = "https://api.mch.weixin.qq.com/pay/orderquery"   // 查询订单
	WXPAY_CLOSEORDER_URL   = "https://api.mch.weixin.qq.com/pay/closeorder"   // 关闭订单
	WXPAY_REFUNDQUERY_URL  = "https://api.mch.weixin.qq.com/pay/refundquery"  // 退款查询
	// 需要证书
	WXPAY_REFUND_URL = "https://api.mch.weixin.qq.com/secapi/pay/refund" // 申请退款
)

//统一下单，WxPayUnifiedOrder中out_trade_no、body、total_fee、trade_type必填
//appid、mchid、spbill_create_ip、nonce_str不需要填入

func UnifiedOrder(config WxPayConfig, params map[string]string) (resp UnifiedOrderResponse, err error) {
	if params["out_trade_no"] == "" {
		err = errors.New("缺少统一支付接口必填参数out_trade_no！")
		return
	} else if params["body"] == "" {
		err = errors.New("缺少统一支付接口必填参数body！")
		return
	} else if params["total_fee"] == "" {
		err = errors.New("缺少统一支付接口必填参数total_fee！")
		return
	}

	if config.TradeType == "JSAPI" && params["openid"] == "" {
		err = errors.New("统一支付接口中，缺少必填参数openid！trade_type为JSAPI时，openid为必填参数！")
		return
	}
	if config.TradeType == "NATIVE" && params["product_id"] == "" {
		err = errors.New("统一支付接口中，缺少必填参数product_id！trade_type为JSAPI时，product_id为必填参数！")
		return
	}

	// 异步通知url未设置，则使用配置文件中的url
	if params["notify_url"] == "" {
		params["notify_url"] = config.NotifyUrl
	}

	// 支付地址
	if params["spbill_create_ip"] == "" {
		params["spbill_create_ip"] = config.SpbillCreateIp
	}

	params["appid"] = config.AppId
	params["mch_id"] = config.MchId
	params["trade_type"] = config.TradeType
	params["nonce_str"] = getNonceStr(32) //随机字符串
	params["sign"] = makeSign(params, config.AppSecret)

	xmlString := map2Xml(params)
	body, err := sendXmlRequest("POST", WXPAY_UNIFIEDORDER_URL, xmlString, config.TlsConfig, config.Timeout)
	if err != nil {
		return
	}

	resp = UnifiedOrderResponse{}
	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return
	} else if resp.ReturnCode != "SUCCESS" {
		err = errors.New(resp.ReturnMsg)
		return
	}

	// 校验
	xmlMap, err := xml2Map(resp)
	if err != nil {
		return
	}

	sign := makeSign(xmlMap, config.AppSecret)
	if resp.Sign != sign {
		err = errors.New("sign err")
		return
	}
	return
}

//查询订单，WxPayOrderQuery中out_trade_no、transaction_id至少填一个
//appid、mchid、spbill_create_ip、nonce_str不需要填入

func OrderQuery(config WxPayConfig, params map[string]string) (resp OrderQueryResponse, err error) {
	if params["out_trade_no"] == "" && params["transaction_id"] == "" {
		err = errors.New("订单查询接口中，out_trade_no、transaction_id至少填一个！")
		return
	}

	params["appid"] = config.AppId
	params["mch_id"] = config.MchId
	params["nonce_str"] = getNonceStr(32) //随机字符串
	params["sign"] = makeSign(params, config.AppSecret)

	xmlString := map2Xml(params)
	body, err := sendXmlRequest("POST", WXPAY_ORDERQUERY_URL, xmlString, config.TlsConfig, config.Timeout)
	if err != nil {
		return
	}

	resp = OrderQueryResponse{}
	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return
	} else if resp.ReturnCode != "SUCCESS" {
		err = errors.New(resp.ReturnMsg)
		return
	}

	// 不校验Sign
	return
}

//关闭订单，WxPayCloseOrder中out_trade_no必填
//appid、mchid、spbill_create_ip、nonce_str不需要填入

func CloseOrder(config WxPayConfig, params map[string]string) (resp CloseOrderResponse, err error) {
	if params["out_trade_no"] == "" {
		err = errors.New("订单关闭接口中，out_trade_no必填！")
		return
	}

	params["appid"] = config.AppId
	params["mch_id"] = config.MchId
	params["nonce_str"] = getNonceStr(32) //随机字符串
	params["sign"] = makeSign(params, config.AppSecret)

	xmlString := map2Xml(params)
	body, err := sendXmlRequest("POST", WXPAY_CLOSEORDER_URL, xmlString, config.TlsConfig, config.Timeout)
	if err != nil {
		return
	}

	resp = CloseOrderResponse{}
	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return
	} else if resp.ReturnCode != "SUCCESS" {
		err = errors.New(resp.ReturnMsg)
		return
	}

	// 校验
	xmlMap, err := xml2Map(resp)
	if err != nil {
		return
	}

	sign := makeSign(xmlMap, config.AppSecret)
	if resp.Sign != sign {
		err = errors.New("sign err")
		return
	}
	return
}

//申请退款，WxPayRefund中out_trade_no、transaction_id至少填一个且
//out_refund_no、total_fee、refund_fee、op_user_id为必填参数

func Refund(config WxPayConfig, params map[string]string) (resp RefundResponse, err error) {
	if params["out_trade_no"] == "" && params["transaction_id"] == "" {
		err = errors.New("退款申请接口中，out_trade_no、transaction_id至少填一个！")
		return
	} else if params["out_refund_no"] == "" {
		err = errors.New("退款申请接口中，缺少必填参数out_refund_no！")
		return
	} else if params["total_fee"] == "" {
		err = errors.New("款申请接口中，缺少必填参数total_fee！")
		return
	} else if params["refund_fee"] == "" {
		err = errors.New("退款申请接口中，缺少必填参数refund_fee！")
		return
	} else if params["op_user_id"] == "" {
		err = errors.New("退款申请接口中，缺少必填参数op_user_id！")
		return
	}

	params["appid"] = config.AppId
	params["mch_id"] = config.MchId
	params["nonce_str"] = getNonceStr(32) //随机字符串
	params["sign"] = makeSign(params, config.AppSecret)

	xmlString := map2Xml(params)
	body, err := sendXmlRequest("POST", WXPAY_REFUND_URL, xmlString, config.TlsConfig, config.Timeout)
	if err != nil {
		return
	}

	resp = RefundResponse{}
	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return
	} else if resp.ReturnCode != "SUCCESS" {
		err = errors.New(resp.ReturnMsg)
		return
	}

	// 校验
	xmlMap, err := xml2Map(resp)
	if err != nil {
		return
	}

	sign := makeSign(xmlMap, config.AppSecret)
	if resp.Sign != sign {
		err = errors.New("sign err")
		return
	}
	return
}

// 查询退款
// 提交退款申请后，通过调用该接口查询退款状态。退款有一定延时，
// 用零钱支付的退款20分钟内到账，银行卡支付的退款3个工作日后重新查询退款状态。
// WxPayRefundQuery中out_refund_no、out_trade_no、transaction_id、refund_id四个参数必填一个
// appid、mchid、spbill_create_ip、nonce_str不需要填入

func RefundQuery(config WxPayConfig, params map[string]string) (resp RefundQueryResponse, err error) {
	if params["out_refund_no"] == "" &&
		params["out_trade_no"] == "" &&
		params["transaction_id"] == "" &&
		params["refund_id"] == "" {

		err = errors.New("退款查询接口中，out_refund_no、out_trade_no、transaction_id、refund_id四个参数必填一个！")
		return
	}

	params["appid"] = config.AppId
	params["mch_id"] = config.MchId
	params["nonce_str"] = getNonceStr(32) //随机字符串
	params["sign"] = makeSign(params, config.AppSecret)

	xmlString := map2Xml(params)
	body, err := sendXmlRequest("POST", WXPAY_REFUNDQUERY_URL, xmlString, config.TlsConfig, config.Timeout)
	if err != nil {
		return
	}

	resp = RefundQueryResponse{}
	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return
	} else if resp.ReturnCode != "SUCCESS" {
		err = errors.New(resp.ReturnMsg)
		return
	}
	// 不校验Sign
	return
}

//
//
////撤销订单API接口，WxPayReverse中参数out_trade_no和transaction_id必须填写一个
////appid、mchid、spbill_create_ip、nonce_str不需要填入
//
//func Reverse(config WxPayConfig, params map[string]string) (resp ReverseResponse, err error) {
//	return
//}
