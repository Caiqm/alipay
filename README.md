# 支付宝相关接口（简易版）
支付宝相关支付、与小程序登录和生成小程序二维码

## 安装

#### 启用 Go module

```go
go get github.com/Caiqm/alipay
```

```go
import github.com/Caiqm/alipay
```

## 支付宝支付使用
```go
// 以支付宝为例，实例化支付宝支付
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