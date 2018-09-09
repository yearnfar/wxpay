package wxpay

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// TradeTypeJSAPI 公众号支付
	TradeTypeJSAPI = "JSAPI"

	// TradeTypeNative 扫码支付
	TradeTypeNative = "NATIVE"

	// TradeTypeApp App支付
	TradeTypeApp = "APP"
)

const (
	// TradeStateSuccess 支付成功
	TradeStateSuccess = "SUCCESS"

	// TradeStateRefund 转入退款
	TradeStateRefund = "REFUND"

	// TradeStateNotPay 未支付
	TradeStateNotPay = "NOTPAY"

	// TradeStateSuccess 已关闭
	TradeStateClosed = "CLOSED"

	// TradeStateRevoked 已撤销（刷卡支付）
	TradeStateRevoked = "REVOKED"

	// TradeStateUserPaying 用户支付中
	TradeStateUserPaying = "USERPAYING"

	// TradeStateUserPayError 支付失败(其他原因，如银行返回失败)
	TradeStateUserPayError = "PAYERROR"
)

const (
	// UnifiedOrderAPI 统一下单接口
	UnifiedOrderAPI = "https://api.mch.weixin.qq.com/pay/unifiedorder"

	// OrderQueryAPI 查询订单接口
	OrderQueryAPI = "https://api.mch.weixin.qq.com/pay/orderquery"
)

// UnifiedOrder 统一下单接口
func UnifiedOrder(cfg *Config, body string, outTradeNo string, totalFee int, openID string, clientIP string) (resp *UnifiedOrderResponse, err error) {
	param := new(UnifiedOrderParam)
	param.Param = NewParam(cfg)
	param.NotifyURL = cfg.NotifyURL
	param.TradeType = cfg.TradeType
	param.Body = body
	param.OutTradeNo = outTradeNo
	param.TotalFee = totalFee
	param.OpenID = openID

	if clientIP != "" {
		param.SpbillCreateIP = clientIP
	} else if cfg.TradeType == TradeTypeNative {
		param.SpbillCreateIP = cfg.ServerAddr
	}

	resp = new(UnifiedOrderResponse)
	err = SendRequest(NewClient(cfg), http.MethodPost, UnifiedOrderAPI, param, resp, cfg.AppSecret)
	return
}

// AppTrade App 支付
func AppTrade(cfg *Config, body string, outTradeNo string, totalFee int, clientIP string) (string, error) {
	if cfg.TradeType != TradeTypeApp {
		return "", fmt.Errorf("支付类型错误，export: %s, got: %s", TradeTypeApp, cfg.TradeType)
	}

	resp, err := UnifiedOrder(cfg, body, outTradeNo, totalFee, "", clientIP)
	if err != nil {
		return "", err
	}

	return resp.PrepayID, nil
}

// JSAPITrade 微信公众号支付
func JSAPITrade(cfg *Config, body string, outTradeNo string, totalFee int, openID string, clientIP string) (string, error) {
	if cfg.TradeType != TradeTypeJSAPI {
		return "", fmt.Errorf("支付类型错误，export: %s, got: %s", TradeTypeJSAPI, cfg.TradeType)
	}

	resp, err := UnifiedOrder(cfg, body, outTradeNo, totalFee, openID, clientIP)
	if err != nil {
		return "", err
	}

	return resp.PrepayID, nil
}

// NativeTrade 扫码支付
func NativeTrade(cfg *Config, body string, outTradeNo string, totalFee int) (string, error) {
	if cfg.TradeType != TradeTypeNative {
		return "", fmt.Errorf("支付类型错误，export: %s, got: %s", TradeTypeNative, cfg.TradeType)
	}

	resp, err := UnifiedOrder(cfg, body, outTradeNo, totalFee, "", "")
	if err != nil {
		return "", err
	}

	return resp.CodeURL, nil
}

// Notify 异步回调
func Notify(cfg *Config, req *http.Request) (resp *NotifyResponse, err error) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	// 写回 body 内容
	req.Body = ioutil.NopCloser(bytes.NewReader(data))

	// 解析获取的xml数据
	resp = new(NotifyResponse)
	err = parseResponse(data, resp, cfg.AppSecret)

	return
}

// OrderQuery 查询订单
func OrderQuery(cfg *Config, transactionID string, outTradeNo string) (resp *OrderQueryResponse, err error) {
	param := new(OrderQueryParam)
	param.Param = NewParam(cfg)
	param.TransactionID = transactionID
	param.OutTradeNo = outTradeNo

	resp = new(OrderQueryResponse)
	err = SendRequest(NewClient(cfg), http.MethodPost, OrderQueryAPI, param, resp, cfg.AppSecret)
	return
}

// SendRequest 发送请求
func SendRequest(client *http.Client, method string, urlStr string, param Parameter, resp Responser, key string) error {
	req, err := NewRequest(method, urlStr, param, key)
	if err != nil {
		return err
	}

	response, err := client.Do(req)
	if err != nil {
		return err
	}

	return ParseResponse(response, resp, key)
}
