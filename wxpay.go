package wxpay

import (
	"net/http"
)

const (
	WXPAY_UNIFIEDORDER_URL = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	WXPAY_ORDERQUERY_URL   = "https://api.mch.weixin.qq.com/pay/unifiedorder"
)

/**
 *
 * 统一下单，WxPayUnifiedOrder中out_trade_no、body、total_fee、trade_type必填
 * appid、mchid、spbill_create_ip、nonce_str不需要填入
 */
func UnifiedOrder(config WxPayConfig, params map[string]string) (err error) {
	if params["out_trade_no"] == "" {
		err = errors.New("缺少统一支付接口必填参数out_trade_no！")
	} else if params["body"] == "" {
		err = errors.New("缺少统一支付接口必填参数body！")
	} else if params["total_fee"] == "" {
		err = errors.New("缺少统一支付接口必填参数total_fee！")
	} else if params["trade_type"] == "" {
		err = errors.New("缺少统一支付接口必填参数trade_type！")
	}

	if params["trade_type"] == "JSAPI" && params["openid"] == "" {
		err = errors.New("统一支付接口中，缺少必填参数openid！trade_type为JSAPI时，openid为必填参数！")
	}

	if params["trade_type"] == "JSAPI" && params["openid"] == "" {
		err = errors.New("统一支付接口中，缺少必填参数openid！trade_type为JSAPI时，openid为必填参数！")
	}

	//异步通知url未设置，则使用配置文件中的url
	if params["notify_url"] == "" {
		params["notify_url"] = config.NotifyUrl
	}

	params["appid"] = config.AppId
	params["mch_id"] = config.MchId
	params["spbill_create_ip"] = config.SpbillCreateIp
	params["nonce_str"] = getNonceStr() //随机字符串
	params["sign"] = makeSign(params, config.Key)

	xmlString := map2Xml(params)

}
