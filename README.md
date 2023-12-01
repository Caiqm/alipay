# 支付宝相关接口（简易版）
支付宝相关支付、小程序登录、小程序二维码、小程序订单同步

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

### 关于应用私钥 (privateKey)

应用私钥是我们通过工具生成的私钥，调用支付宝接口的时候，我们需要使用该私钥对参数进行签名。

### 关于 alipay.New() 函数中的最后一个参数 isProduction

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

## 支付宝小程序订单数据同步

```go
// 订单数据同步接口，注意：需要根据行业来进行调整
func TestClient_OrderPush(t *testing.T) {
	t.Log("========== OrderPush ==========")
	// 商品信息
	var i ItemOrder
	i.ItemName = "开始处理"
	// 素材图片
	materialExtInfo := ExtInfo{
		ExtKey:   "image_material_id",
		ExtValue: "",
	}
	i.ExtInfo = append(i.ExtInfo, materialExtInfo)
	// 额外参数
	var e, e1, e2, e3 ExtInfo
	e.ExtKey = "merchant_order_status"
	e.ExtValue = "START_SERVICE"
	e1.ExtKey = "merchant_order_link_page"
	e1.ExtValue = "/pages/order/orderDetail"
	e2.ExtKey = "tiny_app_id"
	e2.ExtValue = client.appId
	e3.ExtKey = "merchant_biz_type"
	e3.ExtValue = "BUSINESS_REGISTRATION"
	// 主要参数
	var p OrderPush
	p.ItemOrderList = append(p.ItemOrderList, i)
	p.ExtInfo = append(p.ExtInfo, e)
	p.ExtInfo = append(p.ExtInfo, e1)
	p.ExtInfo = append(p.ExtInfo, e2)
	p.ExtInfo = append(p.ExtInfo, e3)
	p.Amount = 1
	p.OrderCreateTime = time.Now().Format(time.DateTime)
	p.OrderModifiedTime = p.OrderCreateTime
	p.SourceApp = "Alipay"
	p.OutBizNo = time.Now().Format("20060102150405") + strconv.Itoa(int(time.Now().Unix()))
	p.OrderType = "SERVICE_ORDER"
	p.BuyerInfo = BuyerInfo{UserId: ""}
	rsp, err := client.OrderPush(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rsp)
}
```