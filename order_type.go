package alipay

// OrderPush 订单数据同步接口 https://opendocs.alipay.com/mini/84f9ee3c_alipay.merchant.order.sync?scene=common&pathHash=103117c9
type OrderPush struct {
	AuxParam
	AppAuthToken      string          `json:"-"`                             // 可选
	OutBizNo          string          `json:"out_biz_no"`                    // 外部订单号 out_biz_no唯一对应一笔订单，相同的订单需传入相同的out_biz_no
	BuyerInfo         BuyerInfo       `json:"buyer_info"`                    // 买家信息
	ServiceCode       string          `json:"service_code,omitempty"`        // 服务code：传入小程序后台提报的服务id，将订单与服务关联，有利于提高服务曝光机会；入参服务id的类目须与订单类型相符，若不相符将会报错；如订单类型为“外卖”，则入参的服务ID所对应的服务类目也必须得是”外卖“；service_code 通过 alipay.open.app.service.apply，(服务提报申请)接口创建服务后获取。
	Amount            float64         `json:"amount"`                        // 订单总金额：某笔交易订单优惠前的总金额，单位为【元】SERVICE_ORDER且不涉及金额可不传入该字段，其他场景必传
	PayAmount         float64         `json:"pay_amount"`                    // 用户应付金额 ：用户最终结算时需要支付金额（不包含选择支付宝付款时，支付宝给予的优惠减免金额），单位为【元】SERVICE_ORDER且不涉及金额可不传入该字段，其他场景必传
	OrderType         string          `json:"order_type"`                    // 订单类型 服务订单: SERVICE_ORDER
	CategoryId        string          `json:"category_id,omitempty"`         // 标准服务类目
	DiscountAmount    float64         `json:"discount_amount,omitempty"`     // 商户总计优惠金额：代表商户侧给予用户的总计优惠金额 （不包含选择支付宝付款时，支付宝给予的优惠减免金额），单位为【元】。
	TradeNo           string          `json:"trade_no"`                      // 订单所对应的支付宝交易号
	TradeType         string          `json:"trade_type"`                    // 交易号类型 交易: TRADE 受托: ENTRUST 转账: TRANSFER
	PayTimeoutExpress string          `json:"pay_timeout_express,omitempty"` // 支付超时时间，超过时间支付宝自行关闭订单，例如：15h
	ItemOrderList     []ItemOrder     `json:"item_order_list,omitempty"`     // 商品信息列表，可选
	LogisticsInfoList []LogisticsInfo `json:"logistics_info_list,omitempty"` // "物流信息 列表最多支持物流信息个数，请参考产品文档 注：若该值不为空，且物流信息同步至我的快递，则在查询订单时可返回具体物流信息"
	ShopInfo          ShopInfo        `json:"shop_info,omitempty"`           // 门店信息，扫码点餐获取返佣时必填。
	OrderCreateTime   string          `json:"order_create_time,omitempty"`   // 订单创建时间 注意事项: 当order_type为SERVICE_ORDER时必传
	OrderModifiedTime string          `json:"order_modified_time"`           // 订单修改时间 注意事项: 用于订单状态或数据变化较快的顺序控制，SERVICE_ORDER按照行业标准化接入场景必须传入该字段控制乱序。order_modified_time较晚的同步会被最终存储，order_modified_time相同的两次同步会被幂等处理
	OrderPayTime      string          `json:"order_pay_time,omitempty"`      // 订单支付时间 注意事项: 当pay_channel为非ALIPAY时，且订单状态已流转到“支付”或支付后时，需要将支付时间传入
	SendMsg           string          `json:"send_msg"`                      // 是否需要小程序订单代理发送模版消息。不传默认不发送，需要发送: Y 不需要发送: N
	DiscountInfoList  []DiscountInfo  `json:"discount_info_list,omitempty"`  // 订单优惠信息
	SyncContent       string          `json:"sync_content"`                  // 同步内容 全部(默认): ALL 仅行程信息: JOURNEY_ONLY
	JourneyOrderList  []JourneyOrder  `json:"journey_order_list,omitempty"`  // 行程信息
	SourceApp         string          `json:"source_app"`                    // 用于区分用户下单的订单来源 支付宝端内: Alipay 钉钉小程序: DingTalk
	ExtInfo           []ExtInfo       `json:"ext_info"`                      // 扩展信息
}

func (o OrderPush) APIName() string {
	return "alipay.merchant.order.sync"
}

func (o OrderPush) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = o.AppAuthToken
	return m
}

// 买家信息
type BuyerInfo struct {
	Name     string    `json:"name,omitempty"`      // 姓名
	Mobile   string    `json:"mobile,omitempty"`    // 手机号
	CertType string    `json:"cert_type,omitempty"` // 证件类型。身份证: IDENTITY_CARD 户口本: HOKOU 护照: PASSPORT 军官证: OFFICER_CARD 士兵证: SOLDIER_CARD
	CertNo   string    `json:"cert_no,omitempty"`   // 证件号
	UserId   string    `json:"user_id,omitempty"`   // 支付宝uid
	ExtInfo  []ExtInfo `json:"ext_info,omitempty"`  // 扩展信息
}

// 额外信息
type ExtInfo struct {
	ExtKey   string `json:"ext_key"`   // 键
	ExtValue string `json:"ext_value"` // 值
}

// 商品信息
type ItemOrder struct {
	SkuId      string    `json:"sku_id"`             // 商品 sku id
	ItemId     string    `json:"item_id"`            // 商品ID 支付宝平台侧商品ID
	ItemName   string    `json:"item_name"`          // 商品名称
	UnitPrice  float64   `json:"unit_price"`         // 商品单价（单位：元）。 小程序订单助手业务中，为必传；其他业务场景参见对应的产品文档。
	Quantity   int64     `json:"quantity"`           // 商品数量（单位：自拟）小程序订单助手业务中，为必传；其他业务场景参见对应的产品文档。
	Unit       string    `json:"unit"`               // 商品规格，例如：斤
	Status     string    `json:"status"`             // 商品状态枚举 初始状态: INIT 已退款: REFUNDED 已关闭: CLOSED 处理中: PROCESSING 待支付: WAIT_PAY 业务完结: BIZ_FINISHED 已完结: SUCCESS_FINISHED 已支付: PAID 已提交: SUBMITTED 业务关闭: BIZ_CLOSED (注意事项: 默认无需传入，如需使用请联系业务负责人)
	StatusDesc string    `json:"status_desc"`        // 商品状态描述，默认无需传入，如需使用请联系业务负责人
	ExtInfo    []ExtInfo `json:"ext_info,omitempty"` // 扩展信息
}

// 物流信息
type LogisticsInfo struct {
	TrackingNo    string `json:"tracking_no"`    // 物流单号
	LogisticsCode string `json:"logistics_code"` // 物流公司编号。 物流公司编号值请查看产品文档。该值为空时，有可能匹配不到物流信息。若有则必传
}

// 门店信息
type ShopInfo struct {
	MerchantShopId       string `json:"merchant_shop_id"`        // 商户门店id 支持英文、数字的组合
	AlipayShopId         string `json:"alipay_shop_id"`          // 蚂蚁门店shop_id
	Name                 string `json:"name"`                    // 店铺名称
	Address              string `json:"address"`                 // 店铺地址
	PhoneNum             string `json:"phone_num"`               // 联系电话-支持固话或手机号 仅支持数字、+、- 。例如 手机：1380***1111、固话：021-888**888
	MerchantShopLinkPage string `json:"merchant_shop_link_page"` // 店铺详情链接地址，例如：pages/shop/index
	Type                 string `json:"type"`                    // 仅当alipay_shop_id字段值为非标准蚂蚁门店时使用，其他场景无需传入 蚂蚁门店: ALIPAY_SHOP 饿了么门店: ELEME_SHOP
}

// 订单优惠信息
type DiscountInfo struct {
	ExternalDiscountId string  `json:"external_discount_id"`         // 外部优惠id
	DiscountName       string  `json:"discount_name"`                // 优惠名称
	DiscountQuantity   int64   `json:"discount_quantity,omitempty"`  // 优惠数量
	DiscountAmount     float64 `json:"discount_amount"`              // 优惠金额 单位为【元】
	DiscountPageLink   string  `json:"discount_page_link,omitempty"` // 优惠跳转链接地址
}

// 行程信息
type JourneyOrder struct {
	MerchantJourneyNo string    `json:"merchant_journey_no"`     // 商户行程单号 注意事项 同一个pid下的行程单号需唯一。同一个pid下外部行程单号唯一
	JourneyDesc       string    `json:"journey_desc,omitempty"`  // 行程描述
	JourneyIndex      string    `json:"journey_index,omitempty"` // 描述本行程为整个行程中的第几程
	Title             string    `json:"title,omitempty"`         // 行程标题
	Status            string    `json:"status,omitempty"`        // 行程状态 注：行程状态必须与支付宝侧进行约定
	StatusDesc        string    `json:"status_desc,omitempty"`   // 行程状态描述
	Type              string    `json:"type"`                    // 行程类型，例如：airticket
	SubType           string    `json:"sub_type"`                // 行程子类型，例如：abroad
	JourneyCreateTime string    `json:"journey_create_time"`     // 行程创建时间
	JourneyModifyTime string    `json:"journey_modify_time"`     // 行程修改时间
	Action            string    `json:"action,omitempty"`        // 操作动作 删除后的行程不再展示: DELETE
	ExtInfo           []ExtInfo `json:"ext_info"`                // 扩展信息
}

// OrderPushRsp 订单数据同步接口响应参数
type OrderPushRsp struct {
	RecordId         string `json:"record_id"`    // 同步订单记录id （自2022年5月19日起，新接入商户，除点餐场景，该字段不再返回）
	OrderId          string `json:"order_id"`     // 支付宝订单号
	OrderStatus      string `json:"order_status"` // 订单状态
	DistributeResult []struct {
		SceneCode           string `json:"scene_code"`            // 分发场景code 订单消息: SERVICE_MSG
		SceneName           string `json:"scene_name"`            // 分发场景名，对应scene_code
		NotDistributeReason string `json:"not_distribute_reason"` // 未分发到场景的具体原因
	} `json:"distribute_result"` // 分发结果 若未分发到场景侧，则会返回具体的未分发原因
	SyncSuggestions []struct {
		Type    string `json:"type"`    // 同步建议类型
		Message string `json:"message"` // 同步建议内容
	} `json:"sync_suggestions"` // 订单同步优化建议
}

// MerchantFileUpload 商品文件上传接口 https://opendocs.alipay.com/mini/510d4a72_alipay.merchant.item.file.upload?scene=common&pathHash=c08922b1
type MerchantFileUpload struct {
	AuxParam
	AppAuthToken string `json:"-"` // 可选
	Scene        string `json:"-"` // 业务场景描述。 小程序订单中心场景固定为 SYNC_ORDER。
}

func (o MerchantFileUpload) APIName() string {
	return "alipay.merchant.item.file.upload"
}

func (o MerchantFileUpload) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = o.AppAuthToken
	m["scene"] = o.Scene
	return m
}

// MerchantFileUploadRsp 商品文件上传接口响应接口
type MerchantFileUploadRsp struct {
	MaterialId  string `json:"material_id"`  // 文件在商品中心的素材标识（素材ID长期有效）
	MaterialKey string `json:"material_key"` // 文件在商品中心的素材标示，创建/更新商品时使用
}
