package wxpay

type WxPayConfig struct {
	AppId     string
	MchId     string
	AppSecret string
	createIp  string
	SSL       WxPaySSLConfig
}

type WxPaySSLConfig struct {
	Cert xx
	Key  xx
}

func NewWxPayConfig(appId, mchId, appSecret, createIp string, ssl *WxPaySSLConfig) (config WxPayConfig) {

}

func NewWxPaySSLConfig(xxPath, xxPath, xxPath) (ssl *WxPaySSLConfig) {

}
