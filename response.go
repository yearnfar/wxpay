package wxpay

import (
	"encoding/xml"
)

type IResponse interface {
	IsSuccess() // 是否成功
	Valid()     //校验
}

type BaseResponse struct {
	XmlName    xml.Name `xml:"xml"`
	ReturnCode string   `xml:"return_code"`
	ReturnMsg  string   `xml:"return_msg"`
}

type UnifiedOrderResponse struct {
	BaseResponse `xml:"-,innerXml"`
	AppId        string `xml:"appId"`
}




func (resp *UnifiedOrderResponse) Valid() {
xml.Unmarshal()
}
