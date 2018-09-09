package wxpay

import (
	"crypto/tls"
	"net/http"
	"time"
)

// Config 配置
type Config struct {
	AppID      string        // 公众账号ID
	AppSecret  string        // AppSecret是APPID对应的接口密码
	MchID      string        // 商户号
	TradeType  string        // 交易类型
	NotifyURL  string        // 通知地址
	ServerAddr string        // 服务器地址
	TLSConfig  *tls.Config   // 证书
	Timeout    time.Duration // 请求超时时间
}

// LoadTLSConfig 载入tls证书
func (c *Config) LoadTLSConfig(crtFile, keyFile string) error {
	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		return err
	}

	tlsConfig := new(tls.Config)
	tlsConfig.Certificates = append(tlsConfig.Certificates, cert)
	c.TLSConfig = tlsConfig

	return nil
}

// NewClient http 客户端
func NewClient(cfg *Config) *http.Client {
	client := &http.Client{}

	if cfg.TLSConfig != nil {
		client.Transport = &http.Transport{TLSClientConfig: cfg.TLSConfig}
	}

	if cfg.Timeout != 0 {
		client.Timeout = cfg.Timeout
	}

	return client
}
