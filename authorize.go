package alipay

// SystemOauthToken 换取授权访问令牌接口 https://docs.open.alipay.com/api_9/alipay.system.oauth.token
func (c *Client) SystemOauthToken(param SystemOauthToken) (result *SystemOauthTokenRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// UserInfoShare 支付宝会员授权信息查询接口 https://docs.open.alipay.com/api_2/alipay.user.info.share
func (c *Client) UserInfoShare(param UserInfoShare) (result *UserInfoShareRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// DecodePhoneNumber 小程序获取会员手机号  https://opendocs.alipay.com/mini/api/getphonenumber
// 本方法用于解码小程序端 my.getPhoneNumber 获取的数据
func (c *Client) DecodePhoneNumber(data []byte) (result *MobileNumber, err error) {
	if err = c.decode(data, "response", true, &result); err != nil {
		return nil, err
	}
	return result, nil
}
