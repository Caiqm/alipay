package alipay

import (
	"testing"
)

func TestClient_TradeMpPay(t *testing.T) {
	t.Log("========== TradeMpPay ==========")
	var p = TradeCreate{}
	p.NotifyURL = "http://203.86.24.181:3000/alipay"
	p.Body = "body"
	p.Subject = "商品标题"
	p.OutTradeNo = "TSD2024112112171234567892"
	p.TotalAmount = "0.01"
	p.ProductCode = "JSAPI_PAY"
	p.OpAppId = "2021003174668195"
	p.OpBuyerOpenId = "2088502347396560"

	param, err := client.TradeCreate(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(param)
}

func TestClient_TradeAppPay(t *testing.T) {
	t.Log("========== TradeAppPay ==========")
	var p = TradeAppPay{}
	p.NotifyURL = "http://203.86.24.181:3000/alipay"
	p.Body = "body"
	p.Subject = "商品标题"
	p.OutTradeNo = "0101010111111111"
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

func TestClient_TradeWapPay(t *testing.T) {
	t.Log("========== TradeWapPay ==========")
	var p = TradeWapPay{}
	p.NotifyURL = "http://220.112.233.229:3000/alipay"
	p.ReturnURL = "http://220.112.233.229:3000"
	p.Subject = "object"
	p.OutTradeNo = "trade_no_201706230111212"
	p.TotalAmount = "0.01"
	p.ProductCode = "QUICK_WAP_WAY"
	url, err := client.TradeWapPay(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}

func TestClient_TradeQuery(t *testing.T) {
	t.Log("========== TradeQuery ==========")
	var p = TradeQuery{}
	p.OutTradeNo = ""
	//p.TradeNo = ""
	param, err := client.TradeQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(param)
}

func TestClient_TradeRefund(t *testing.T) {
	t.Log("========== TradeRefund ==========")
	var p = TradeRefund{}
	p.OutTradeNo = ""
	p.RefundAmount = ""
	p.OutRequestNo = ""
	param, err := client.TradeRefund(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(param)
}

func TestClient_TradeFastPayRefundQuery(t *testing.T) {
	t.Log("========== TradeFastPayRefundQuery ==========")
	var p = TradeFastPayRefundQuery{}
	p.OutTradeNo = ""
	p.OutRequestNo = ""
	param, err := client.TradeFastPayRefundQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(param)
}
