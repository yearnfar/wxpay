# wxpay
go语言实现微信支付对接

# usage

```golang

// 初始化一个App支付
payConfig := &wxpay.WxPayConfig{
	AppId:         "微信分配的公众账号ID（企业号corpid即为此appId）",
	AppSecret:     "APP密钥",
	MchId:         "微信支付分配的商户号",
	NotifyUrl:     "接收微信支付异步通知回调地址",
	TradeType:     "APP",  // 支持 JSAPI，NATIVE，APP
}

// 统一下单
params := map[string]string{
  "out_trade_no": "1006010215040000",  // 商户订单号
	"body":         "test",  // 商品描述
	"total_fee":    "100",   // 100分 = 1元
}
resp, err := wxpay.UnifiedOrder(payConfig, params)

// 查询订单
params := map[string]string{
	"out_trade_no": "100000000000",
}
resp, err := wxpay.OrderQuery(payConfig, params)

// 关闭订单
params := map[string]string{
	"out_trade_no": "1310001417536",
}
resp, err := wxpay.CloseOrder(payConfig, params)

// 申请退款 -- 需要证书
tlsConfig, err := NewWxPayTlsConfig(
  "app/apiclient_cert.pem",
  "app/apiclient_key.pem",
  "app/rootca.pem",     // 无则不传
)

params := map[string]string{
	"out_trade_no":  "11111111111111",
	"out_refund_no": fmt.Sprintf("%d", time.Now().Unix()),
	"total_fee":     "100",
	"refund_fee":    "100",
	"op_user_id":    "test",
}
resp, err := wxpay.Refund(payConfig, params)

// 查询退款
params := map[string]string{
	"out_trade_no": "222222222222222",
}
resp, err := wxpay.RefundQuery(payConfig, params)

```


