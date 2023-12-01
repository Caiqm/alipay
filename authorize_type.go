package alipay

// SystemOauthToken 换取授权访问令牌接口请求参数 https://docs.open.alipay.com/api_9/alipay.system.oauth.token
type SystemOauthToken struct {
	AuxParam
	AppAuthToken string `json:"-"` // 可选
	GrantType    string `json:"-"` // 值为 authorization_code 时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"-"`
	RefreshToken string `json:"-"`
}

func (s SystemOauthToken) APIName() string {
	return "alipay.system.oauth.token"
}

func (s SystemOauthToken) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = s.AppAuthToken
	m["grant_type"] = s.GrantType
	if s.Code != "" {
		m["code"] = s.Code
	}
	if s.RefreshToken != "" {
		m["refresh_token"] = s.RefreshToken
	}
	return m
}

// SystemOauthTokenRsp 换取授权访问令牌接口请求参数
type SystemOauthTokenRsp struct {
	Error
	UserId       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	ReExpiresIn  int64  `json:"re_expires_in"`
	AuthStart    string `json:"auth_start"`
	OpenId       string `json:"open_id"`
}

// UserInfoShare 支付宝会员授权信息查询接口请求参数 https://docs.open.alipay.com/api_2/alipay.user.info.share
type UserInfoShare struct {
	AuxParam
	AppAuthToken string `json:"-"` // 可选
	AuthToken    string `json:"-"` // 是
}

func (u UserInfoShare) APIName() string {
	return "alipay.user.info.share"
}

func (u UserInfoShare) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = u.AppAuthToken
	m["auth_token"] = u.AuthToken
	return m
}

// UserInfoShareRsp 支付宝会员授权信息查询接口响应参数
type UserInfoShareRsp struct {
	Error
	AuthNo             string `json:"auth_no"`
	UserId             string `json:"user_id"`
	Avatar             string `json:"avatar"`
	Province           string `json:"province"`
	City               string `json:"city"`
	NickName           string `json:"nick_name"`
	IsStudentCertified string `json:"is_student_certified"`
	UserType           string `json:"user_type"`
	UserStatus         string `json:"user_status"`
	IsCertified        string `json:"is_certified"`
	Gender             string `json:"gender"`
}
