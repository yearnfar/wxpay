package wxpay

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

// Responser 验证
type Responser interface {
	CheckReturn() error
	CheckSign(data []byte, key string) error
}

// Response 业务返回
type Response struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode string   `xml:"return_code"`
	ReturnMsg  string   `xml:"return_msg"`
	AppId      string   `xml:"appid"`        // 公众账号ID
	MchId      string   `xml:"mch_id"`       // 商户号
	DeviceInfo string   `xml:"device_info"`  // 设备号
	NonceStr   string   `xml:"nonce_str"`    // 随机字符串
	Sign       string   `xml:"sign"`         // 签名
	SignType   string   `xml:"sign_type"`    // 签名类型
	ResultCode string   `xml:"result_code"`  // 业务结果
	ErrCode    string   `xml:"err_code"`     // 错误码
	ErrCodeDes string   `xml:"err_code_des"` // 错误描述
}

// CheckReturn 此字段是通信标识，非交易标识，交易是否成功需要查看trade_state来判断
func (r Response) CheckReturn() error {
	if r.ReturnCode != "SUCCESS" {
		return errors.New(r.ReturnCode + ", " + r.ReturnMsg)
	}

	return nil
}

// CheckSign 验证sign
func (r Response) CheckSign(data []byte, key string) error {
	if r.Sign == "" {
		return errors.New("sign为空")
	}

	sign, err := makeSign(data, key)
	if err != nil {
		return err
	}

	if r.Sign != sign {
		return errors.New("sign错误")
	}

	return nil
}

// IsResultOK 返回成
func (r Response) IsResultOK() bool {
	return r.ResultCode == "SUCCESS"
}

// ParseResponse 解析返回数据
func ParseResponse(response *http.Response, resp Responser, key string) error {
	defer response.Body.Close()
	xmlBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(xmlBody, &resp)
	if err != nil {
		return err
	}

	err = resp.CheckReturn()
	if err != nil {
		return err
	}

	err = resp.CheckSign(xmlBody, key)
	if err != nil {
		return err
	}

	return nil
}
