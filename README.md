wxpay

微信支付Golang sdk v2版

使用方法：

统一下单 UnifiedOrder

cfg := &Config{
		AppID:      "wx11111111111",     // 微信分配的公众账号ID（企业号corpid即为此appId）
		MchID:      "138888888",         // 微信支付分配的商户号
		TradeType:  "NATIVE",            // 支付方式
		AppSecret:  "66666666666666666", // APP密钥
		NotifyURL:  "www.example.com",   // 接收微信支付异步通知回调地址
		ServerAddr: "127.0.0.1",         // 当前服务器ip地址
	}

resp, err := wxpay.UnifiedOrder(cfg, "统一下单测试", "110000000001", 1, "", "")


公众号支付 JSAPITrade

cfg := &Config{
		AppID:      "wx11111111111",     // 微信分配的公众账号ID（企业号corpid即为此appId）
		MchID:      "138888888",         // 微信支付分配的商户号
		TradeType:  "JSAPI",            // 支付方式
		AppSecret:  "66666666666666666", // APP密钥
		NotifyURL:  "www.example.com",   // 接收微信支付异步通知回调地址
		ServerAddr: "127.0.0.1",         // 当前服务器ip地址
	}

resp, err := wxpay.JSAPITrade(cfg, "支付测试", "100000000001", 1, "", "")

扫码支付 NativeTrade

支付方式Native

cfg := &Config{
		AppID:      "wx11111111111",     // 微信分配的公众账号ID（企业号corpid即为此appId）
		MchID:      "138888888",         // 微信支付分配的商户号
		TradeType:  "NATIVE",            // 支付方式
		AppSecret:  "66666666666666666", // APP密钥
		NotifyURL:  "www.example.com",   // 接收微信支付异步通知回调地址
		ServerAddr: "127.0.0.1",         // 当前服务器ip地址
	}

codeURL, err := NativeTrade(cfg, "支付测试", "100000000001")


