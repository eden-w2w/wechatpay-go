package coupons

import "time"

type GetStocksRequest struct {
	Offset            uint32     `json:"offset" default:"0"`
	Limit             uint32     `json:"limit" default:"10"`
	StockCreatorMchID string     `json:"stock_creator_mchid"`
	CreateStartTime   *time.Time `json:"create_start_time,omitempty"`
	CreateEndTime     *time.Time `json:"create_end_time,omitempty"`
	Status            *string    `json:"status,omitempty"`
}

type RuleFixedNormalCoupon struct {
	// 面额
	CouponAmount *int64 `json:"coupon_amount"`
	// 门槛
	TransactionMinimum *int64 `json:"transaction_minimum"`
}

type StockUseRule struct {
	// 发放总上限
	MaxCoupons *int64 `json:"max_coupons"`
	// 总预算
	MaxAmount *int64 `json:"max_amount"`
	// 单日发放上限金额
	MaxAmountByDay *int64 `json:"max_amount_by_day"`
	// 固定面额批次特定信息
	FixedNormalCoupon *RuleFixedNormalCoupon `json:"fixed_normal_coupon,omitempty"`
	// 单个用户可领个数
	MaxCouponsPerUser *int32 `json:"max_coupons_per_user"`
	// 券类型
	// NORMAL：满减券
	// CUT_TO：减至券
	CouponType *string `json:"coupon_type,omitempty"`
	// 支持的支付方式
	// MICROAPP：小程序支付
	// APPPAY：APP支付
	// PPAY：免密支付
	// CARD：付款码支付
	// FACE：人脸支付
	// OTHER：（公众号、扫码等）
	TradeType []string `json:"trade_type"`
	// 是否可叠加其他优惠
	CombineUse *bool `json:"combine_use,omitempty"`
}

type StockCutToMessage struct {
	// 可用优惠的商品最高单价
	SinglePriceMax *int64 `json:"single_price_max"`
	// 减至后的优惠单价
	CutToPrice *int64 `json:"cut_to_price"`
}

type Stock struct {
	// 批次号
	// 微信为每个代金券批次分配的唯一id。
	StockID *string `json:"stock_id"`
	// 创建批次的商户号
	// 微信为创建方商户分配的商户号。
	StockCreatorMchID *string `json:"stock_creator_mchid"`
	// 批次名称
	StockName *string `json:"stock_name"`
	// 批次状态
	// unactivated：未激活
	// audit：审核中
	// running：运行中
	// stoped：已停止
	// paused：暂停发放
	Status *string `json:"status"`
	// 创建时间
	CreateTime *time.Time `json:"create_time"`
	// 使用说明
	Description *string `json:"description"`
	// 满减券批次使用规则
	StockUseRule *StockUseRule `json:"stock_use_rule,omitempty"`
	// 可用开始时间
	AvailableBeginTime *time.Time `json:"available_begin_time"`
	// 可用结束时间
	AvailableEndTime *time.Time `json:"available_end_time"`
	// 已发券数量
	DistributedCoupons *int64 `json:"distributed_coupons"`
	// 是否无资金流
	NoCash *bool `json:"no_cash"`
	// 激活批次的时间
	StartTime *time.Time `json:"start_time"`
	// 批次终止的时间
	StopTime *time.Time `json:"stop_time"`
	// 减至批次特定信息
	CutToMessage *StockCutToMessage `json:"cut_to_message,omitempty"`
	// 是否单品优惠
	SingleItem *bool `json:"singleitem"`
	// 批次类型
	// NORMAL：代金券批次
	// DISCOUNT_CUT：立减与折扣
	// OTHER：其他
	StockType *string `json:"stock_type"`
}

type StocksPagination struct {
	Total  int64   `json:"total_count"`
	Data   []Stock `json:"data,omitempty"`
	Limit  uint32  `json:"limit"`
	Offset uint32  `json:"offset"`
}

type GiveCouponRequest struct {
	StockID           string  `json:"stock_id"`
	OpenID            *string `json:"openid"`
	OutRequestNo      string  `json:"out_request_no"`
	AppID             string  `json:"appid"`
	StockCreatorMchID string  `json:"stock_creator_mchid"`
}

type GiveCouponResponse struct {
	// 代金券ID
	CouponID string `json:"coupon_id"`
}
