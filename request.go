package wxpay

import (
	"bytes"
	"encoding/xml"
	"errors"
	"net/http"
)

// Parameter 接口
type Parameter interface {
	Valid() error
	SetSign(sign string)
}

// Param 参数
type Param struct {
	XMLName xml.Name `xml:"xml"`

	// AppID 微信支付分配的公众账号ID（企业号corpid即为此appId）
	AppID string `xml:"appid"`

	// MchID 微信支付分配的商户号
	MchID string `xml:"mch_id"`

	// DeviceInfo 自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"
	DeviceInfo string `xml:"device_info"`

	// NonceStr 随机字符串，长度要求在32位以内。推荐随机数生成算法
	NonceStr string `xml:"nonce_str"`

	// Sign 通过签名算法计算得出的签名值，详见签名生成算法
	Sign string `xml:"sign"`

	// SignType 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
	SignType string `xml:"sign_type,omitempty"`
}

// Valid 参数校验
func (p *Param) Valid() error {
	if p.AppID == "" {
		return errors.New("公众账号ID不能为空")
	}

	if p.MchID == "" {
		return errors.New("商户号不能为空")
	}

	if p.NonceStr == "" {
		return errors.New("随机字符串不能为空")
	}
	return nil
}

// SetSign 设置Sign
func (p *Param) SetSign(sign string) {
	p.Sign = sign
}

// NewParam 参数
func NewParam(cfg *Config) Param {
	return Param{
		AppID:    cfg.AppID,
		MchID:    cfg.MchID,
		NonceStr: makeNonceStr(20),
	}
}

// NewRequest 实例化一个request
func NewRequest(method string, urlStr string, param Parameter, key string) (*http.Request, error) {
	data, err := toXML(param, key)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	return req, nil
}
