# 支付宝相关接口（简易版）
支付宝相关支付、与小程序登录和生成小程序二维码

## 安装

#### 启用 Go module

```go
go get github.com/Caiqm/alipay
```

```go
import "github.com/Caiqm/alipay"
```

#### 未启用 Go module

```go
go get github.com/Caiqm/alipay
```

```go
import "github.com/Caiqm/alipay"
```

## 如何使用
```go
// 实例化支付宝支付
var client, err = alipay.New(appID, privateKey, isProduction)
```

#### 关于应用私钥 (privateKey)

应用私钥是我们通过工具生成的私钥，调用支付宝接口的时候，我们需要使用该私钥对参数进行签名。

#### 关于 alipay.New() 函数中的最后一个参数 isProduction

支付宝提供了用于开发时测试的 sandbox 环境，对接的时候需要注意相关的 app id 和密钥是 sandbox 环境还是 production 环境的。如果是 sandbox 环境，本参数应该传 false，否则为 true。

### 普通公钥模式

需要注意此处用到的公钥是**支付宝公钥**，不是我们用工具生成的应用公钥。

[如何查看支付宝公钥？](https://opendocs.alipay.com/common/057aqe)

```go
client.LoadAlipayCertPublicKey("aliPublicKey")
```

## 支付宝APP支付
```go
func TestClient_TradeAppPay(t *testing.T) {
	t.Log("========== TradeAppPay ==========")
	var p = TradeAppPay{}
	p.NotifyURL = "http://203.86.24.181:3000/alipay"
	p.Body = "body"
	p.Subject = "商品标题"
	p.OutTradeNo = "01010101"
	p.TotalAmount = "100.00"
	p.ProductCode = "p_1010101"
	param, err := client.TradeAppPay(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(param)
}
```

## 支付宝小程序登录
```go
// 换取授权访问令牌接口
func TestClient_SystemOauthToken(t *testing.T) {
    t.Log("========== SystemOauthToken ==========")
    var p = SystemOauthToken{}
    p.GrantType = "authorization_code"
    p.Code = "647f16afe0b44c49a8eb1cb3c02aXX31"
    rsp, err := client.SystemOauthToken(p)
    if err != nil {
    t.Fatal(err)
    }
    if rsp.IsFailure() {
    t.Fatal(rsp.Msg, rsp.SubMsg)
    }
    t.Logf("%v", rsp)
}
```