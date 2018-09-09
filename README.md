# wxpay

微信支付Golang sdk v2版

# 使用方法：

## 支付

### 统一下单 UnifiedOrder

```
cfg := &wxpay.Config{
		AppID:      "wx11111111111",       // 微信分配的公众账号ID（企业号corpid即为此appId）
		MchID:      "138888888",           // 微信支付分配的商户号
		TradeType:  wxpay.TradeTypeNative, // 支付方式
		AppSecret:  "66666666666666666",   // APP密钥
		NotifyURL:  "www.example.com",     // 接收微信支付异步通知回调地址
		ServerAddr: "192.168.1.101",       // 当前服务器ip地址
	}

resp, err := wxpay.UnifiedOrder(cfg, "统一下单测试", "110000000001", 1, "", "")
```

### APP支付 AppTrade

```
cfg := &wxpay.Config{
		AppID:      "wx11111111111",     // 微信分配的公众账号ID（企业号corpid即为此appId）
		MchID:      "138888888",         // 微信支付分配的商户号
		TradeType:  wxpay.TradeTypeApp,  // 支付方式
		AppSecret:  "66666666666666666", // APP密钥
		NotifyURL:  "www.example.com",   // 接收微信支付异步通知回调地址
		ServerAddr: "192.168.1.101",     // 当前服务器ip地址
	}

prepayID, err := wxpay.App(cfg, "APP支付测试", "100000000001", 1, "114.114.114.114")
```

### 公众号支付 JSAPITrade

```
cfg := &wxpay.Config{
		AppID:      "wx11111111111",      // 微信分配的公众账号ID（企业号corpid即为此appId）
		MchID:      "138888888",          // 微信支付分配的商户号
		TradeType:  wxpay.TradeTypeJSAPI, // 支付方式
		AppSecret:  "66666666666666666",  // APP密钥
		NotifyURL:  "www.example.com",    // 接收微信支付异步通知回调地址
		ServerAddr: "192.168.1.101",      // 当前服务器ip地址
	}

prepayID, err := wxpay.JSAPITrade(cfg, "公众号支付测试", "100000000001", 1, "xxxxxxxxxx", "114.114.114.114")
```


### 扫码支付 NativeTrade

```
cfg := &wxpay.Config{
		AppID:      "wx11111111111",     // 微信分配的公众账号ID（企业号corpid即为此appId）
		MchID:      "138888888",         // 微信支付分配的商户号
		TradeType:  wxpay.TradeTypeNative, // 支付方式
		AppSecret:  "66666666666666666", // APP密钥
		NotifyURL:  "www.example.com",   // 接收微信支付异步通知回调地址
		ServerAddr: "127.0.0.1",         // 当前服务器ip地址
	}

codeURL, err := wxpay.NativeTrade(cfg, "扫码支付测试", "100000000001")
```
