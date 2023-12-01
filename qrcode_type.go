package alipay

// CreateQrcode 小程序生成推广二维码接口 https://opendocs.alipay.com/mini/a25c5d8f_alipay.open.app.qrcode.create?scene=common&pathHash=2334bbff
type CreateQrcode struct {
	AuxParam
	AppAuthToken string `json:"-"`           // 可选
	UrlParam     string `json:"url_param"`   // 跳转小程序的页面路径，小程序中能访问到的页面路径
	QueryParam   string `json:"query_param"` // 小程序的启动参数，打开小程序的query ，在小程序 onLaunch的方法中获取
	Describe     string `json:"describe"`    // 对应的二维码描述
	Color        string `json:"color"`       // 圆形二维码颜色（十六进制颜色色值），仅圆形二维码支持颜色设置，方形二维码默认为黑色
	Size         string `json:"size"`        // 合成后图片的大小规格，有s、m、l三档可选，s: 8cm大小 m: 12cm大小 l: 30cm大小
}

func (c CreateQrcode) APIName() string {
	return "alipay.open.app.qrcode.create"
}

func (c CreateQrcode) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = c.AppAuthToken
	return m
}

// CreateQrcodeRsp 小程序生成推广二维码接口响应参数
type CreateQrcodeRsp struct {
	Error
	QrCodeUrl            string `json:"qr_code_url"`              // 方形二维码图片链接地址
	QrCodeUrlCircleWhite string `json:"qr_code_url_circle_white"` // 圆形二维码地址，白色slogan
	QrCodeUrlCircleBlue  string `json:"qr_code_url_circle_blue"`  // 圆形二维码地址，蓝色slogan
}
