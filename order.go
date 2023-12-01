package alipay

// OrderPush 订单数据同步接口 https://opendocs.alipay.com/mini/84f9ee3c_alipay.merchant.order.sync?scene=common&pathHash=103117c9
func (c *Client) OrderPush(param OrderPush) (result *OrderPushRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// MerchantFileUpload 商品文件上传接口 https://opendocs.alipay.com/mini/510d4a72_alipay.merchant.item.file.upload?scene=common&pathHash=c08922b1
func (c *Client) MerchantFileUpload(param MerchantFileUpload) (result *MerchantFileUploadRsp, err error) {
	c.scene = param.Scene
	err = c.postFormData("POST", param, &result)
	return
}
