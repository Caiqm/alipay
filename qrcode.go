package alipay

// CreateQrcode 小程序生成推广二维码接口 https://opendocs.alipay.com/mini/a25c5d8f_alipay.open.app.qrcode.create?scene=common&pathHash=2334bbff
func (c *Client) CreateQrcode(param CreateQrcode) (result *CreateQrcodeRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}
