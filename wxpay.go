package wxpay

import (
	"fmt"
	"net/http"
)

const (
	TradeTypeJSAPI  = "JSAPI"
	TradeTypeNative = "NATIVE"
	TradeTypeApp    = "APP"
)

const (
	// UnifiedOrderAPI 统一下单接口
	UnifiedOrderAPI = "https://api.mch.weixin.qq.com/pay/unifiedorder"
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
	err = Send(NewClient(cfg), http.MethodPost, UnifiedOrderAPI, param, resp, cfg.AppSecret)
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
