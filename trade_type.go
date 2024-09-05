package alipay

type Trade struct {
	AuxParam
	NotifyURL    string `json:"-"`
	ReturnURL    string `json:"-"`
	AppAuthToken string `json:"-"` // 可选

	// biz content，这四个参数是必须的
	Body        string `json:"body,omitempty"`
	Subject     string `json:"subject"`      // 订单标题
	OutTradeNo  string `json:"out_trade_no"` // 商户订单号，64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
	TotalAmount string `json:"total_amount"` // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	ProductCode string `json:"product_code"` // 销售产品码，与支付宝签约的产品码名称。 参考官方文档, App 支付时默认值为 QUICK_MSECURITY_PAY
}

// TradePagePay 统一收单下单并支付页面接口请求参数 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay
type TradePagePay struct {
	Trade
	AuthToken   string `json:"auth_token,omitempty"`   // 针对用户授权接口，获取用户相关数据时，用于标识用户授权关系
	QRPayMode   string `json:"qr_pay_mode,omitempty"`  // PC扫码支付的方式，支持前置模式和跳转模式。
	QRCodeWidth string `json:"qrcode_width,omitempty"` // 商户自定义二维码宽度 注：qr_pay_mode=4时该参数生效
}

func (t TradePagePay) APIName() string {
	return "alipay.trade.page.pay"
}

func (t TradePagePay) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = t.AppAuthToken
	m["notify_url"] = t.NotifyURL
	m["return_url"] = t.ReturnURL
	return m
}

type TradeStatus string

const (
	TradeStatusWaitBuyerPay TradeStatus = "WAIT_BUYER_PAY" //（交易创建，等待买家付款）
	TradeStatusClosed       TradeStatus = "TRADE_CLOSED"   //（未付款交易超时关闭，或支付完成后全额退款）
	TradeStatusSuccess      TradeStatus = "TRADE_SUCCESS"  //（交易支付成功）
	TradeStatusFinished     TradeStatus = "TRADE_FINISHED" //（交易结束，不可退款）
)

// TradeQuery 统一收单线下交易查询接口请求参数 https://docs.open.alipay.com/api_1/alipay.trade.query/
type TradeQuery struct {
	AuxParam
	AppAuthToken string   `json:"-"`                       // 可选
	OutTradeNo   string   `json:"out_trade_no,omitempty"`  // 订单支付时传入的商户订单号, 与 TradeNo 二选一
	TradeNo      string   `json:"trade_no,omitempty"`      // 支付宝交易号
	OrgPid       string   `json:"org_pid,omitempty"`       // 可选 银行间联模式下有用，其它场景请不要使用； 双联通过该参数指定需要查询的交易所属收单机构的pid;
	QueryOptions []string `json:"query_options,omitempty"` // 可选 查询选项，商户传入该参数可定制本接口同步响应额外返回的信息字段，数组格式。支持枚举如下：trade_settle_info：返回的交易结算信息，包含分账、补差等信息； fund_bill_list：交易支付使用的资金渠道；voucher_detail_list：交易支付时使用的所有优惠券信息；discount_goods_detail：交易支付所使用的单品券优惠的商品优惠信息；mdiscount_amount：商家优惠金额；
}

func (t TradeQuery) APIName() string {
	return "alipay.trade.query"
}

func (t TradeQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = t.AppAuthToken
	return m
}

// TradeQueryRsp 统一收单线下交易查询接口响应参数
type TradeQueryRsp struct {
	Error
	TradeNo               string           `json:"trade_no"`                      // 支付宝交易号
	OutTradeNo            string           `json:"out_trade_no"`                  // 商家订单号
	BuyerLogonId          string           `json:"buyer_logon_id"`                // 买家支付宝账号
	TradeStatus           TradeStatus      `json:"trade_status"`                  // 交易状态
	TotalAmount           string           `json:"total_amount"`                  // 交易的订单金额
	TransCurrency         string           `json:"trans_currency"`                // 标价币种
	SettleCurrency        string           `json:"settle_currency"`               // 订单结算币种
	SettleAmount          string           `json:"settle_amount"`                 // 结算币种订单金额
	PayCurrency           string           `json:"pay_currency"`                  // 订单支付币种
	PayAmount             string           `json:"pay_amount"`                    // 支付币种订单金额
	SettleTransRate       string           `json:"settle_trans_rate"`             // 结算币种兑换标价币种汇率
	TransPayRate          string           `json:"trans_pay_rate"`                // 标价币种兑换支付币种汇率
	BuyerPayAmount        string           `json:"buyer_pay_amount"`              // 买家实付金额，单位为元，两位小数。
	PointAmount           string           `json:"point_amount"`                  // 积分支付的金额，单位为元，两位小数。
	InvoiceAmount         string           `json:"invoice_amount"`                // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
	SendPayDate           string           `json:"send_pay_date"`                 // 本次交易打款给卖家的时间
	ReceiptAmount         string           `json:"receipt_amount"`                // 实收金额，单位为元，两位小数
	StoreId               string           `json:"store_id"`                      // 商户门店编号
	TerminalId            string           `json:"terminal_id"`                   // 商户机具终端编号
	FundBillList          []*FundBill      `json:"fund_bill_list,omitempty"`      // 交易支付使用的资金渠道
	StoreName             string           `json:"store_name"`                    // 请求交易支付中的商户店铺的名称
	BuyerUserId           string           `json:"buyer_user_id"`                 // 买家在支付宝的用户id
	BuyerUserName         string           `json:"buyer_user_name"`               // 买家名称；
	IndustrySepcDetailGov string           `json:"industry_sepc_detail_gov"`      // 行业特殊信息-统筹相关
	IndustrySepcDetailAcc string           `json:"industry_sepc_detail_acc"`      // 行业特殊信息-个账相关
	ChargeAmount          string           `json:"charge_amount"`                 // 该笔交易针对收款方的收费金额；
	ChargeFlags           string           `json:"charge_flags"`                  // 费率活动标识，当交易享受活动优惠费率时，返回该活动的标识；
	SettlementId          string           `json:"settlement_id"`                 // 支付清算编号，用于清算对账使用；
	TradeSettleInfo       *TradeSettleInfo `json:"trade_settle_info,omitempty"`   // 返回的交易结算信息，包含分账、补差等信息
	AuthTradePayMode      string           `json:"auth_trade_pay_mode"`           // 预授权支付模式，该参数仅在信用预授权支付场景下返回。信用预授权支付：CREDIT_PREAUTH_PAY
	BuyerUserType         string           `json:"buyer_user_type"`               // 买家用户类型。CORPORATE:企业用户；PRIVATE:个人用户。
	MdiscountAmount       string           `json:"mdiscount_amount"`              // 商家优惠金额
	DiscountAmount        string           `json:"discount_amount"`               // 平台优惠金额
	Subject               string           `json:"subject"`                       // 订单标题；
	Body                  string           `json:"body"`                          // 订单描述;
	AlipaySubMerchantId   string           `json:"alipay_sub_merchant_id"`        // 间连商户在支付宝端的商户编号；
	ExtInfos              string           `json:"ext_infos"`                     // 交易额外信息，特殊场景下与支付宝约定返回。
	PassbackParams        string           `json:"passback_params"`               // 公用回传参数。返回支付时传入的passback_params参数信息
	HBFQPayInfo           *HBFQPayInfo     `json:"hb_fq_pay_info"`                // 若用户使用花呗分期支付，且商家开通返回此通知参数，则会返回花呗分期信息。json格式其它说明详见花呗分期信息说明。 注意：商家需与支付宝约定后才返回本参数。
	CreditPayMode         string           `json:"credit_pay_mode"`               // 信用支付模式。表示订单是采用信用支付方式（支付时买家没有出资，需要后续履约）。"creditAdvanceV2"表示芝麻先用后付模式，用户后续需要履约扣款。 此字段只有信用支付场景才有值，商户需要根据字段值单独处理。此字段以后可能扩展其他值，建议商户使用白名单方式识别，对于未识别的值做失败处理，并联系支付宝技术支持人员。
	CreditBizOrderId      string           `json:"credit_biz_order_id"`           // 信用业务单号。信用支付场景才有值，先用后付产品里是芝麻订单号。
	HYBAmount             string           `json:"hyb_amount"`                    // 惠营宝回票金额
	BKAgentRespInfo       *BKAgentRespInfo `json:"bk_agent_resp_info"`            // 间联交易下，返回给机构的信 息
	ChargeInfoList        []*ChargeInfo    `json:"charge_info_list"`              // 计费信息列表
	DiscountGoodsDetail   string           `json:"discount_goods_detail"`         // 本次交易支付所使用的单品券优惠的商品优惠信息
	VoucherDetailList     []*VoucherDetail `json:"voucher_detail_list,omitempty"` // 本交易支付时使用的所有优惠券信息
}

type HBFQPayInfo struct {
	UserInstallNum string `json:"user_install_num"` // 用户使用花呗分期支付的分期数
}

type BKAgentRespInfo struct {
	BindtrxId        string `json:"bindtrx_id"`
	BindclrissrId    string `json:"bindclrissr_id"`
	BindpyeracctbkId string `json:"bindpyeracctbk_id"`
	BkpyeruserCode   string `json:"bkpyeruser_code"`
	EstterLocation   string `json:"estter_location"`
}

type ChargeInfo struct {
	ChargeFee               string          `json:"charge_fee"`
	OriginalChargeFee       string          `json:"original_charge_fee"`
	SwitchFeeRate           string          `json:"switch_fee_rate"`
	IsRatingOnTradeReceiver string          `json:"is_rating_on_trade_receiver"`
	IsRatingOnSwitch        string          `json:"is_rating_on_switch"`
	ChargeType              string          `json:"charge_type"`
	SubFeeDetailList        []*SubFeeDetail `json:"sub_fee_detail_list"`
}

type SubFeeDetail struct {
	ChargeFee         string `json:"charge_fee"`
	OriginalChargeFee string `json:"original_charge_fee"`
	SwitchFeeRate     string `json:"switch_fee_rate"`
}

type FundBill struct {
	FundChannel string  `json:"fund_channel"`       // 交易使用的资金渠道，详见 支付渠道列表
	BankCode    string  `json:"bank_code"`          // 银行卡支付时的银行代码
	Amount      string  `json:"amount"`             // 该支付工具类型所使用的金额
	RealAmount  float64 `json:"real_amount,string"` // 渠道实际付款金额
}

type VoucherDetail struct {
	Id                 string `json:"id"`                  // 券id
	Name               string `json:"name"`                // 券名称
	Type               string `json:"type"`                // 当前有三种类型： ALIPAY_FIX_VOUCHER - 全场代金券, ALIPAY_DISCOUNT_VOUCHER - 折扣券, ALIPAY_ITEM_VOUCHER - 单品优惠
	Amount             string `json:"amount"`              // 优惠券面额，它应该会等于商家出资加上其他出资方出资
	MerchantContribute string `json:"merchant_contribute"` // 商家出资（特指发起交易的商家出资金额）
	OtherContribute    string `json:"other_contribute"`    // 其他出资方出资金额，可能是支付宝，可能是品牌商，或者其他方，也可能是他们的一起出资
	Memo               string `json:"memo"`                // 优惠券备注信息
}

type TradeSettleInfo struct {
	TradeSettleDetailList []*TradeSettleDetail `json:"trade_settle_detail_list"`
}

type TradeSettleDetail struct {
	OperationType     string `json:"operation_type"`
	OperationSerialNo string `json:"operation_serial_no"`
	OperationDate     string `json:"operation_dt"`
	TransOut          string `json:"trans_out"`
	TransIn           string `json:"trans_in"`
	Amount            string `json:"amount"`
}

// TradeRefund 统一收单交易退款接口请求参数 https://docs.open.alipay.com/api_1/alipay.trade.refund/
type TradeRefund struct {
	AuxParam
	AppAuthToken            string                    `json:"-"`                                   // 可选
	OutTradeNo              string                    `json:"out_trade_no,omitempty"`              // 与 TradeNo 二选一
	TradeNo                 string                    `json:"trade_no,omitempty"`                  // 与 OutTradeNo 二选一
	RefundAmount            string                    `json:"refund_amount"`                       // 必须 需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数
	RefundReason            string                    `json:"refund_reason"`                       // 可选 退款的原因说明
	OutRequestNo            string                    `json:"out_request_no"`                      // 必须 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。
	RefundRoyaltyParameters []*RefundRoyaltyParameter `json:"refund_royalty_parameters,omitempty"` // 可选 退分账明细信息。 注： 1.当面付且非直付通模式无需传入退分账明细，系统自动按退款金额与订单金额的比率，从收款方和分账收入方退款，不支持指定退款金额与退款方。2.直付通模式，电脑网站支付，手机 APP 支付，手机网站支付产品，须在退款请求中明确是否退分账，从哪个分账收入方退，退多少分账金额；如不明确，默认从收款方退款，收款方余额不足退款失败。不支持系统按比率退款。
	QueryOptions            []string                  `json:"query_options,omitempty"`             // 可选 查询选项。 商户通过上送该参数来定制同步需要额外返回的信息字段，数组格式。支持：refund_detail_item_list：退款使用的资金渠道；deposit_back_info：触发银行卡冲退信息通知；
	// OperatorId   string `json:"operator_id"`            // 可选 商户的操作员编号
	// StoreId    string `json:"store_id"`    // 可选 商户的门店编号
	// TerminalId string `json:"terminal_id"` // 可选 商户的终端编号
}

func (t TradeRefund) APIName() string {
	return "alipay.trade.refund"
}

func (t TradeRefund) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = t.AppAuthToken
	return m
}

type RefundRoyaltyParameter struct {
	RoyaltyType  string `json:"royalty_type,omitempty"`   // 可选 分账类型. 普通分账为：transfer;补差为：replenish;为空默认为分账transfer;
	TransOut     string `json:"trans_out,omitempty"`      // 可选 支出方账户。如果支出方账户类型为userId，本参数为支出方的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字；如果支出方类型为loginName，本参数为支出方的支付宝登录号。 泛金融类商户分账时，该字段不要上送。
	TransOutType string `json:"trans_out_type,omitempty"` // 可选 支出方账户类型。userId表示是支付宝账号对应的支付宝唯一用户号;loginName表示是支付宝登录号； 泛金融类商户分账时，该字段不要上送。
	TransInType  string `json:"trans_in_type,omitempty"`  // 可选 收入方账户类型。userId表示是支付宝账号对应的支付宝唯一用户号;cardAliasNo表示是卡编号;loginName表示是支付宝登录号；
	TransIn      string `json:"trans_in"`                 // 必选 收入方账户。如果收入方账户类型为userId，本参数为收入方的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字；如果收入方类型为cardAliasNo，本参数为收入方在支付宝绑定的卡编号；如果收入方类型为loginName，本参数为收入方的支付宝登录号；
	Amount       string `json:"amount,omitempty"`         // 可选 分账的金额，单位为元
	Desc         string `json:"desc,omitempty"`           // 可选 分账描述
	RoyaltyScene string `json:"royalty_scene,omitempty"`  // 可选 可选值：达人佣金、平台服务费、技术服务费、其他
	TransInName  string `json:"trans_in_name,omitempty"`  // 可选 分账收款方姓名，上送则进行姓名与支付宝账号的一致性校验，校验不一致则分账失败。不上送则不进行姓名校验
}

// TradeRefundRsp 统一收单交易退款接口响应参数
type TradeRefundRsp struct {
	Error
	TradeNo              string              `json:"trade_no"`                          // 支付宝交易号
	OutTradeNo           string              `json:"out_trade_no"`                      // 商户订单号
	BuyerLogonId         string              `json:"buyer_logon_id"`                    // 用户的登录id
	BuyerUserId          string              `json:"buyer_user_id"`                     // 买家在支付宝的用户id
	FundChange           string              `json:"fund_change"`                       // 本次退款是否发生了资金变化
	RefundFee            string              `json:"refund_fee"`                        // 退款总金额
	StoreName            string              `json:"store_name"`                        // 交易在支付时候的门店名称
	RefundDetailItemList []*TradeFundBill    `json:"refund_detail_item_list,omitempty"` // 退款使用的资金渠道
	SendBackFee          string              `json:"send_back_fee"`                     // 本次商户实际退回金额。 说明：如需获取该值，需在入参query_options中传入 refund_detail_item_list。
	RefundHYBAmount      string              `json:"refund_hyb_amount"`                 // 本次请求退惠营宝金额
	RefundChargeInfoList []*RefundChargeInfo `json:"refund_charge_info_list,omitempty"` // 退费信息
}

type TradeFundBill struct {
	FundChannel string `json:"fund_channel"` // 交易使用的资金渠道，详见 支付渠道列表
	Amount      string `json:"amount"`       // 该支付工具类型所使用的金额
	RealAmount  string `json:"real_amount"`  // 渠道实际付款金额
	FundType    string `json:"fund_type"`    // 渠道所使用的资金类型
}

type RefundChargeInfo struct {
	RefundChargeFee        string                `json:"refund_charge_fee"`                    // 实退费用
	SwitchFeeRate          string                `json:"switch_fee_rate"`                      // 签约费率
	ChargeType             string                `json:"charge_type"`                          // 收单手续费trade，花呗分期手续hbfq，其他手续费charge
	RefundSubFeeDetailList []*RefundSubFeeDetail `json:"refund_sub_fee_detail_list,omitempty"` // 组合支付退费明细
}

type RefundSubFeeDetail struct {
	RefundChargeFee string `json:"refund_charge_fee"` // 实退费用
	SwitchFeeRate   string `json:"switch_fee_rate"`   // 签约费率
}

// TradeFastPayRefundQuery 统一收单交易退款查询接口请求参数 https://docs.open.alipay.com/api_1/alipay.trade.fastpay.refund.query
type TradeFastPayRefundQuery struct {
	AuxParam
	AppAuthToken string   `json:"-"`                       // 可选
	OutTradeNo   string   `json:"out_trade_no,omitempty"`  // 与 TradeNo 二选一
	TradeNo      string   `json:"trade_no,omitempty"`      // 与 OutTradeNo 二选一
	OutRequestNo string   `json:"out_request_no"`          // 必须 请求退款接口时，传入的退款请求号，如果在退款请求时未传入，则该值为创建交易时的外部交易号
	QueryOptions []string `json:"query_options,omitempty"` // 可选 查询选项，商户通过上送该参数来定制同步需要额外返回的信息字段，数组格式。 refund_detail_item_list
}

func (t TradeFastPayRefundQuery) APIName() string {
	return "alipay.trade.fastpay.refund.query"
}

func (t TradeFastPayRefundQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = t.AppAuthToken
	return m
}

// TradeFastPayRefundQueryRsp 统一收单交易退款查询接口响应参数
type TradeFastPayRefundQueryRsp struct {
	Error
	TradeNo              string              `json:"trade_no"`                          // 支付宝交易号
	OutTradeNo           string              `json:"out_trade_no"`                      // 创建交易传入的商户订单号
	OutRequestNo         string              `json:"out_request_no"`                    // 本笔退款对应的退款请求号
	TotalAmount          string              `json:"total_amount"`                      // 发该笔退款所对应的交易的订单金额
	RefundAmount         string              `json:"refund_amount"`                     // 本次退款请求，对应的退款金额
	RefundStatus         string              `json:"refund_status"`                     // 退款状态。枚举值： REFUND_SUCCESS 退款处理成功； 未返回该字段表示退款请求未收到或者退款失败；
	RefundRoyaltys       []*RefundRoyalty    `json:"refund_royaltys"`                   // 退分账明细信息
	GMTRefundPay         string              `json:"gmt_refund_pay"`                    // 退款时间。
	RefundDetailItemList []*TradeFundBill    `json:"refund_detail_item_list"`           // 本次退款使用的资金渠道； 默认不返回该信息，需要在入参的query_options中指定"refund_detail_item_list"值时才返回该字段信息。
	SendBackFee          string              `json:"send_back_fee"`                     // 本次商户实际退回金额；
	DepositBackInfo      []*DepositBackInfo  `json:"deposit_back_info"`                 // 银行卡冲退信息
	RefundHYBAmount      string              `json:"refund_hyb_amount"`                 // 本次请求退惠营宝金额
	RefundChargeInfoList []*RefundChargeInfo `json:"refund_charge_info_list,omitempty"` // 退费信息
}

type RefundRoyalty struct {
	RefundAmount  string `json:"refund_amount"`
	RoyaltyType   string `json:"royalty_type"`
	ResultCode    string `json:"result_code"`
	TransOut      string `json:"trans_out"`
	TransOutEmail string `json:"trans_out_email"`
	TransIn       string `json:"trans_in"`
	TransInEmail  string `json:"trans_in_email"`
}

type DepositBackInfo struct {
	HasDepositBack     string `json:"has_deposit_back"`
	DBackStatus        string `json:"dback_status"`
	DBackAmount        string `json:"dback_amount"`
	BankAckTime        string `json:"bank_ack_time"`
	ESTBankReceiptTime string `json:"est_bank_receipt_time"`
}

// TradeCreate 统一收单交易创建接口请求参数 https://docs.open.alipay.com/api_1/alipay.trade.create/
type TradeCreate struct {
	Trade
	DiscountableAmount string             `json:"discountable_amount"`        // 可打折金额. 参与优惠计算的金额，单位为元，精确到小数点后两位
	BuyerId            string             `json:"buyer_id"`                   // 买家支付宝用户ID。 2088开头的16位纯数字，小程序场景下获取用户ID请参考：用户授权; 其它场景下获取用户ID请参考：网页授权获取用户信息; 注：交易的买家与卖家不能相同。
	BuyerOpenId        string             `json:"buyer_open_id"`              // 新版接口无法获取user_id, 这里只能传递openid值
	OpAppId            string             `json:"op_app_id,omitempty"`        // 小程序支付中，商户实际经营主体的小程序应用的appid, 注意事项:商户需要先在产品管理中心绑定该小程序appid，否则下单会失败
	OpBuyerOpenId      string             `json:"op_buyer_open_id,omitempty"` // 买家支付宝用户唯一标识（商户实际经营主体的小程序应用关联的买家open_id）
	GoodsDetail        []*GoodsDetailItem `json:"goods_detail,omitempty"`
	OperatorId         string             `json:"operator_id"`
	TerminalId         string             `json:"terminal_id"`
	SellerId           string             `json:"seller_id,omitempty"`       // 卖家支付宝用户ID
	TimeExpire         string             `json:"time_expire,omitempty"`     // 订单绝对超时时间。格式为yyyy-MM-dd HH:mm:ss。【示例值】2021-12-31 10:05:00
	TimeoutExpress     string             `json:"timeout_express,omitempty"` // 订单相对超时时间。从交易创建时间开始计算。该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。当面付场景默认值为3h。time_expire和timeout_express两者只需传入一个或者都不传，如果两者都传，优先使用time_expire。
	PassbackParams     string             `json:"passback_params,omitempty"` // 公用回传参数。 如果请求时传递了该参数，支付宝会在异步通知时将该参数原样返回。【示例值】merchantBizType%3d3C%26merchantBizNo%3d2016010101111
}

func (t TradeCreate) APIName() string {
	return "alipay.trade.create"
}

func (t TradeCreate) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = t.AppAuthToken
	m["notify_url"] = t.NotifyURL
	return m
}

// TradeCreateRsp 统一收单交易创建接口响应参数
type TradeCreateRsp struct {
	Error
	TradeNo    string `json:"trade_no"` // 支付宝交易号
	OutTradeNo string `json:"out_trade_no"`
}

type GoodsDetailItem struct {
	GoodsId       string `json:"goods_id"`
	AliPayGoodsId string `json:"alipay_goods_id"`
	GoodsName     string `json:"goods_name"`
	Quantity      string `json:"quantity"`
	Price         string `json:"price"`
	GoodsCategory string `json:"goods_category"`
	Body          string `json:"body"`
	ShowUrl       string `json:"show_url"`
}

// TradeAppPay App支付接口请求参数 https://docs.open.alipay.com/api_1/alipay.trade.app.pay/
type TradeAppPay struct {
	Trade
}

func (t TradeAppPay) APIName() string {
	return "alipay.trade.app.pay"
}

func (t TradeAppPay) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = t.AppAuthToken
	m["notify_url"] = t.NotifyURL
	return m
}

func (t TradeAppPay) NeedEncrypt() bool {
	return false
}

// TradeWapPay 手机网站支付接口请求参数 https://docs.open.alipay.com/api_1/alipay.trade.wap.pay/
type TradeWapPay struct {
	Trade
	QuitURL    string `json:"quit_url,omitempty"`
	AuthToken  string `json:"auth_token,omitempty"`  // 针对用户授权接口，获取用户相关数据时，用于标识用户授权关系
	TimeExpire string `json:"time_expire,omitempty"` // 绝对超时时间，格式为yyyy-MM-dd HH:mm。
}

func (t TradeWapPay) APIName() string {
	return "alipay.trade.wap.pay"
}

func (t TradeWapPay) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = t.AppAuthToken
	m["notify_url"] = t.NotifyURL
	m["return_url"] = t.ReturnURL
	return m
}
