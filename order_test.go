package alipay

import (
	"strconv"
	"testing"
	"time"
)

// 订单同步
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

// 素材上传
func TestClient_MerchantFileUpload(t *testing.T) {
	t.Log("========== MerchantFileUpload ==========")
	// 设置图片信息，fileType=true则是网络图片，fileType=false则是本地图片，需要绝对路径
	client.LoadOptionFunc(WithFileInformation("test.png", "https://test_detail.png", true))
	// 加载公钥
	_ = client.LoadAlipayCertPublicKey("")
	var p MerchantFileUpload
	p.Scene = "SYNC_ORDER"
	rsp, err := client.MerchantFileUpload(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", rsp)
}
