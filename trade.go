package alipay

import "net/url"

// TradePagePay 统一收单下单并支付页面接口 https://docs.open.alipay.com/api_1/alipay.trade.page.pay
func (c *Client) TradePagePay(param TradePagePay) (result *url.URL, err error) {
	p, err := c.URLValues(param)
	if err != nil {
		return nil, err
	}
	result, err = url.Parse(c.host + "?" + p.Encode())
	if err != nil {
		return nil, err
	}
	return result, err
}

// TradeAppPay App支付接口 https://docs.open.alipay.com/api_1/alipay.trade.app.pay
func (c *Client) TradeAppPay(param TradeAppPay) (result string, err error) {
	p, err := c.URLValues(param)
	if err != nil {
		return "", err
	}
	return p.Encode(), err
}

// TradeCreate 统一收单交易创建接口 https://docs.open.alipay.com/api_1/alipay.trade.create/
func (c *Client) TradeCreate(param TradeCreate) (result *TradeCreateRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// TradeWapPay 手机网站支付接口 https://docs.open.alipay.com/api_1/alipay.trade.wap.pay/
func (c *Client) TradeWapPay(param TradeWapPay) (result *url.URL, err error) {
	p, err := c.URLValues(param)
	if err != nil {
		return nil, err
	}
	result, err = url.Parse(c.host + "?" + p.Encode())
	if err != nil {
		return nil, err
	}
	return result, err
}

// TradeFastPayRefundQuery 统一收单交易退款查询接口 https://docs.open.alipay.com/api_1/alipay.trade.fastpay.refund.query
func (c *Client) TradeFastPayRefundQuery(param TradeFastPayRefundQuery) (result *TradeFastPayRefundQueryRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// TradeRefund 统一收单交易退款接口 https://docs.open.alipay.com/api_1/alipay.trade.refund/
func (c *Client) TradeRefund(param TradeRefund) (result *TradeRefundRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// TradeQuery 统一收单线下交易查询接口 https://docs.open.alipay.com/api_1/alipay.trade.query/
func (c *Client) TradeQuery(param TradeQuery) (result *TradeQueryRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}
