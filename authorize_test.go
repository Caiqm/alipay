package alipay

import (
	"log"
	"testing"
)

var (
	client *Client
	appId  = ""
	priKey = ""
)

func init() {
	var err error
	client, err = New(appId, priKey, true)
	if err != nil {
		log.Fatalln(err)
	}
	client.OnReceivedData(func(method string, data []byte) {
		log.Println(method, string(data))
	})
}

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

// 支付宝会员授权信息查询接口
func TestClient_UserInfoShare(t *testing.T) {
	t.Log("========== UserInfoShare ==========")
	var p = UserInfoShare{}
	p.AuthToken = "authusrB133e40c363934488a9c3e25e17fd9X31"
	rsp, err := client.UserInfoShare(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}
