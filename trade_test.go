package alipay

import (
	"testing"
)

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

func TestClient_TradePagePay(t *testing.T) {
	t.Log("========== TradePagePay ==========")
	var p = TradePagePay{}
	p.NotifyURL = "http://220.112.233.229:3000/alipay"
	p.ReturnURL = "http://220.112.233.229:3000"
	p.Subject = "修正了中文的 Bug"
	p.OutTradeNo = "trade_no_201706230111212"
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	url, err := client.TradePagePay(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}
