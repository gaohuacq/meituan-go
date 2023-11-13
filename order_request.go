package meituan

type OrderBaseRequest struct {
	// 订单基础信息
	OrderID string `json:"order_id"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/134
type OrderDetailRequest struct {
	OrderBaseRequest
	IsMtLogistics *int `json:"is_mt_logistics,omitempty"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/112
type OrderCancelRequest struct {
	// 商家取消订单
	OrderBaseRequest
	Reason     string `json:"reason"`
	ReasonCode int    `json:"reason_code"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/166
type OrderRemindReplyRequest struct {
	// 催单回复
	OrderBaseRequest
	RemindID     int64  `json:"remind_id"`
	ReplyID      int64  `json:"reply_id"`
	ReplyContent string `json:"reply_content"`
}

type OrderRefundRequest struct {
	// 订单退款拒绝或同意通用结构体
	OrderBaseRequest
	Reason string `json:"reason"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/227
type OrderApplyCompensationRequest struct {
	// 申请餐损赔付
	OrderBaseRequest
	Reason      string  `json:"reason"`
	ApplyStatus int     `json:"apply_status"`
	Amount      float64 `json:"amount"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/252
type OrderBatchFetchAbnormalOrderRequest struct {
	// 批量获取异常订单
	StartTime int `json:"start_time"`
	EndTime   int `json:"end_time"`
	Type      int `json:"type"`
	Offset    int `json:"offset"`
	Limit     int `json:"limit"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/351
type EcommerceOrderLogisticsSyncRequest struct {
	// 自配送商家同步发货状态和配送信息（推荐）
	OrderBaseRequest
	ThirdCarrierOrderID       string  `json:"third_carrier_order_id"`
	CourierName               *string `json:"courier_name,omitempty"`
	CourierPhone              *string `json:"courier_phone,omitempty"`
	PhoneType                 *int    `json:"phone_type,omitempty"`
	PrivacyNumValiditySeconds *int    `json:"privacy_num_validity_seconds,omitempty"`
	LogisticsProviderCode     string  `json:"logistics_provider_code"`
	LogisticsStatus           int     `json:"logistics_status"`
	Latitude                  *string `json:"latitude,omitempty"`
	Longitude                 *string `json:"longitude,omitempty"`
	DeliveryPostion           *string `json:"delivery_position,omitempty"`
	DeliveryPositionPictures  *string `json:"delivery_position_pictures,omitempty"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/410
type EcommerceOrderReviewAfterSalesRequest struct {
	// 售后单（退款/退货退款）审核接口
	WmOrderIDView     string  `json:"wm_order_id_view"`
	ReviewType        int     `json:"review_type"`
	RejectReasonCode  *int    `json:"reject_reason_code,omitempty"`
	RejectOtherReason *string `json:"reject_other_reason,omitempty"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/341
type EcommerceGetOrderIdByDaySeqRequest struct {
	// 根据流水号获取订单ID（分段）
	AppPoiCode  string `json:"app_poi_code"`
	DateTime    string `json:"date_time"`
	DaySeqStart int    `json:"day_seq_start"`
	DaySeqEnd   int    `json:"day_seq_end"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/146
type GetOrderDaySeqRequest struct {
	// 获取当日最新订单流水号
	AppPoiCode string `json:"app_poi_code"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/149
type GetOrderIdByDaySeqRequest struct {
	// 根据流水号获取订单ID
	AppPoiCode string `json:"app_poi_code"`
	DateTime   string `json:"date_time"`
	DaySeq     int    `json:"day_seq"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/321
type EcommerceOrderGetOrderRefundDetailRequest struct {
	// 获取订单退款详情
	WmOrderIDView string `json:"wm_order_id_view"`
	RefundType    int    `json:"refund_type"`
}
