package meituan

// 文档地址：https://open-shangou.meituan.com/home/docDetail/134
type OrderDetailResponse struct {
	// 获取订单详细信息
	BaseResponse
	Data OrderDetail `json:"data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/128
type OrderViewStatusResponse struct {
	// 查询订单状态
	BaseResponse
	Data []struct {
		Status int `json:"status"`
	} `json:"data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/341
type EcommerceGetOrderIdByDaySeqResponse struct {
	// 根据流水号获取订单ID（分段）
	BaseResponse
	Data []struct {
		Result   string  `json:"result"`
		OrderIDs []int64 `json:"order_ids"`
	} `json:"data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/146
type GetOrderDaySeqResponse struct {
	// 根据订单ID获取流水号
	BaseResponse
	Data []struct {
		Result string `json:"result"`
		DaySeq int    `json:"day_seq"`
	} `json:"data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/149
type GetOrderIdByDaySeqResponse struct {
	// 根据流水号获取订单ID
	BaseResponse
	Data []struct {
		Result  string `json:"result"`
		OrderID int64  `json:"order_id"`
	} `json:"data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/321
type EcommerceOrderGetOrderRefundDetailResponse struct {
	// 获取订单退款记录
	BaseResponse
	Data []OrderRefundRecord `json:"data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/219
type OrderGetPreparationMealTimeResponse struct {
	BaseResponse
	Data struct {
		PreparationMealTime int `json:"preparationMealTime"`
	} `json:"data"`
}

type OrderDetail struct {
	OrderID                   int64                      `json:"order_id"`
	WmOrderIDView             int64                      `json:"wm_order_id_view"`
	AppPoiCode                string                     `json:"app_poi_code"`
	WmPoiName                 string                     `json:"wm_poi_name"`
	WmPoiAddress              string                     `json:"wm_poi_address"`
	WmPoiPhone                string                     `json:"wm_poi_phone"`
	RecipientAddress          string                     `json:"recipient_address"`
	DeliveryPosition          string                     `json:"delivery_position"`
	RecipientPhone            string                     `json:"recipient_phone"`
	RecipientName             string                     `json:"recipient_name"`
	BackupRecipientPhone      []string                   `json:"backup_recipient_phone"`
	ShippingFee               float64                    `json:"shipping_fee"`
	Total                     float64                    `json:"total"`
	OriginalPrice             float64                    `json:"original_price"`
	Caution                   string                     `json:"caution"`
	ShipperPhone              string                     `json:"shipper_phone"`
	Status                    int                        `json:"status"`
	CityID                    int64                      `json:"city_id"`
	HasInvoiced               int                        `json:"has_invoiced"`
	InvoiceTitle              string                     `json:"invoice_title"`
	TaxpayerID                string                     `json:"taxpayer_id"`
	Ctime                     int64                      `json:"ctime"`
	Utime                     int64                      `json:"utime"`
	DeliveryTime              int64                      `json:"delivery_time"`
	IsThirdShipping           int                        `json:"is_third_shipping"`
	PayType                   int                        `json:"pay_type"`
	PickType                  int                        `json:"pick_type"`
	Latitude                  float64                    `json:"latitude"`
	Longitude                 float64                    `json:"longitude"`
	DaySeq                    int                        `json:"day_seq"`
	IsFavorites               bool                       `json:"is_favorites"`
	IsPoiFirstOrder           bool                       `json:"is_poi_first_order"`
	IsPreOrder                bool                       `json:"is_pre_order"`
	IsPreSaleOrder            bool                       `json:"is_pre_sale_order"`
	DinnersNumber             int                        `json:"dinners_number"`
	LogisticsCode             string                     `json:"logistics_code"`
	PickupCode                *string                    `json:"pickup_code,omitempty"`
	PoiReceiveDetail          *PoiReceiveDetail          `json:"poi_receive_detail,omitempty"`
	Detail                    []OrderDetailData          `json:"detail"`
	Extras                    []OrderExtras              `json:"extras"`
	SkuBenefitDetail          *OrderSkuBenefit           `json:"sku_benefit_detail,omitempty"`
	UserMemberInfo            *OrderUserMember           `json:"user_member_info,omitempty"`
	AvgSendTime               *float64                   `json:"avg_send_time"`
	OrderSendTime             *int64                     `json:"order_send_time"`
	OrderShippingAddress      string                     `json:"order_shipping_address"`
	OrderReceiveTime          *int64                     `json:"order_receive_time"`
	OrderConfirmTime          *int64                     `json:"order_confirm_time"`
	OrderCancelTime           *int64                     `json:"order_cancel_time"`
	OrderCompletedTime        *int64                     `json:"order_completed_time"`
	LogisticsStatus           *int                       `json:"logistics_status"`
	LogisticsID               *int                       `json:"logistics_id"`
	LogisticsName             *string                    `json:"logistics_name"`
	LogisticsSendTime         *int64                     `json:"logistics_send_time"`
	LogisticsConfirmTime      *int64                     `json:"logistics_confirm_time"`
	LogisticsCancelTime       *int64                     `json:"logistics_cancel_time"`
	LogisticsFetchTime        *int64                     `json:"logistics_fetch_time"`
	LogisticsCompletedTime    *int64                     `json:"logistics_completed_time"`
	LogisticsDispatcherName   *string                    `json:"logistics_dispatcher_name"`
	LogisticsDispatcherMobile *string                    `json:"logistics_dispatcher_mobile"`
	PackageBagMoney           int                        `json:"package_bag_money"`
	EstimateArrivalTime       int                        `json:"estimate_arrival_time"`
	PackageBagMoneyYuan       string                     `json:"package_bag_money_yuan"`
	PoiReceiveDetailYuan      *OrderPoiReceiveDetailYuan `json:"poi_receive_detail_yuan"`
	TotalWeight               int                        `json:"total_weight"`
	OrderPhoneNumber          string                     `json:"order_phone_number"`
	IncmpCode                 int                        `json:"incmp_code"`
	IncmpModules              int                        `json:"incmp_modules"`
	ScanDeliveryFlag          *bool                      `json:"scan_delivery_flag"`
	ScanDeliveryQrContent     *string                    `json:"scan_delivery_qr_content"`
	IsWholeCityShip           *int                       `json:"is_whole_city_ship"`
	OrderTagList              []int                      `json:"order_tag_list"`
	CycleInfo                 OrderCycleInfo             `json:"cycle_info"`
}

type PoiReceiveDetail struct {
	ActOrderChargeByMt      []ActOrderCharge     `json:"actOrderChargeByMt"`
	ActOrderChargeByPoi     []ActOrderCharge     `json:"actOrderChargeByPoi"`
	FoodShareFeeChargeByPoi int64                `json:"foodShareFeeChargeByPoi"`
	ReconciliationExtras    ReconciliationExtras `json:"reconciliationExtras"`
	LogisticsFee            int64                `json:"logisticsFee"`
	OnlinePayment           int64                `json:"onlinePayment"`
	WmPoiReceiveCent        int64                `json:"wmPoiReceiveCent"`
}

type ActOrderCharge struct {
	Comment     string  `json:"comment"`
	FeeTypeDesc string  `json:"feeTypeDesc"`
	FeeTypeID   int     `json:"feeTypeId"`
	MoneyCent   *int64  `json:"moneyCent,omitempty"`
	Money       *string `json:"money,omitempty"`
}

type ReconciliationExtras struct {
	ChargeModel           int      `json:"chargeModel"`
	PerformanceServiceFee float64  `json:"performanceServiceFee"`
	TechnicalServiceFee   float64  `json:"technicalServiceFee"`
	DistanceFee           float64  `json:"distanceFee"`
	SaleFee               float64  `json:"saleFee"`
	BaseShippingAmount    float64  `json:"baseShippingAmount"`
	CategoryChargeFee     float64  `json:"categoryChargeFee"`
	WeightChargeFee       float64  `json:"weightChargeFee"`
	HolidayChargeFee      float64  `json:"holidayChargeFee"`
	ExplosionChargeFee    *float64 `json:"explosionChargeFee,omitempty"`
}

type OrderDetailData struct {
	BackupDetail  OrderDetailBackup `json:"backup_detail"`
	ActualPrice   string            `json:"actual_price"`
	OriginalPrice string            `json:"original_price"`
	AppSpuCode    string            `json:"app_spu_code"`
	FoodName      string            `json:"food_name"`
	SkuID         string            `json:"sku_id"`
	ItemID        int64             `json:"item_id"`
	Upc           string            `json:"upc"`
	Quanity       int               `json:"quanity"`
	Price         float64           `json:"price"`
	BoxNum        float64           `json:"box_num"`
	BoxPrice      float64           `json:"box_price"`
	Unit          string            `json:"unit"`
	FoodDiscount  float64           `json:"food_discount"`
	FoodProperty  string            `json:"food_property"`
	CartID        int               `json:"cart_id"`
	Weight        int64             `json:"weight"`
	WeightForUnit string            `json:"weight_for_unit"`
	WeightUnit    string            `json:"weight_unit"`
	Picture       string            `json:"picture"`
	LocationCode  string            `json:"location_code"`
}

type OrderDetailBackup struct {
	SkuID           string  `json:"sku_id"`
	AppSpuCode      string  `json:"app_spu_code"`
	SpuName         string  `json:"spu_name"`
	OriginalPrice   float64 `json:"original_price"`
	Spec            string  `json:"spec"`
	Unit            string  `json:"unit"`
	PoiChangedCount int     `json:"poi_changed_count"`
	PoiChangedTime  int64   `json:"poi_changed_time"`
}

type OrderExtras struct {
	ActDetailID         int                `json:"act_detail_id"`
	ReduceFee           float64            `json:"reduce_fee"`
	MtCharge            float64            `json:"mt_charge"`
	PoiCharge           float64            `json:"poi_charge"`
	PromotionServiceFee float64            `json:"promotion_service_fee"`
	Remark              float64            `json:"remark"`
	Type                int                `json:"type"`
	Type9CouponType     *int               `json:"type_9_coupon_type,omitempty"`
	ActExtendMsg        *OrderActExtendMsg `json:"act_extend_msg"`
}

type OrderActExtendMsg struct {
	AppSpuCode      string `json:"app_spu_code"`
	SkuID           string `json:"sku_id"`
	GiftsAppSpuCode string `json:"gifts_app_spu_code"`
	GiftsSkuID      string `json:"gifts_sku_id"`
	GiftsName       string `json:"gifts_name"`
	GiftNum         int    `json:"gift_num"`
}

type OrderSkuBenefit struct {
	AppSpuCode               string           `json:"app_spu_code"`
	Name                     string           `json:"name"`
	SkuID                    string           `json:"sku_id"`
	Upc                      string           `json:"upc"`
	Count                    int              `json:"count"`
	TotalOriginPrice         float64          `json:"totalOriginPrice"`
	OriginPrice              float64          `json:"originPrice"`
	TotalReducePrice         float64          `json:"totalReducePrice"`
	TotalActivityPrice       float64          `json:"totalActivityPrice"`
	ActivityPrice            float64          `json:"activityPrice"`
	TotalMtCharge            float64          `json:"totalMtCharge"`
	TotalPoiCharge           float64          `json:"totalPoiCharge"`
	TotalPromotionServiceFee float64          `json:"totalPromotionServiceFee"`
	TotalBoxPrice            float64          `json:"totalBoxPrice"`
	BoxPrice                 float64          `json:"boxPrice"`
	BoxNumber                float64          `json:"boxNumber"`
	WmAppOrderActDetails     WmOrderActDetail `json:"wmAppOrderActDetails"`
}

type WmOrderActDetail struct {
	ActID               int64   `json:"act_id"`
	SkuActID            int64   `json:"sku_act_id"`
	PoiActID            int64   `json:"poi_act_id"`
	BatchActID          int64   `json:"batch_act_id"`
	Type                int     `json:"type"`
	Remark              string  `json:"remark"`
	MtCharge            float64 `json:"mt_charge"`
	PoiCharge           float64 `json:"poi_charge"`
	PromotionServiceFee float64 `json:"promotion_service_fee"`
	Count               int     `json:"count"`
}

type OrderUserMember struct {
	IsPoiMember bool   `json:"is_poi_member"`
	CardCode    string `json:"card_code"`
	LevelCode   string `json:"level_code"`
	PoiMemberID string `json:"poi_member_id"`
}

type OrderPoiReceiveDetailYuan struct {
	ActOrderChargeByMt      []ActOrderCharge     `json:"actOrderChargeByMt"`
	ActOrderChargeByPoi     []ActOrderCharge     `json:"actOrderChargeByPoi"`
	FoodShareFeeChargeByPoi string               `json:"foodShareFeeChargeByPoi"`
	LogisticsFee            string               `json:"logisticsFee"`
	OnlinePayment           string               `json:"onlinePayment"`
	PoiReceive              string               `json:"poiReceive"`
	ReconciliationExtras    ReconciliationExtras `json:"reconciliationExtras"`
}

type OrderCycleInfo struct {
	TotalCount int    `json:"total_count"`
	Time       string `json:"time"`
	Rule       struct {
		Day  string `json:"day"`
		Type int    `json:"type"`
	} `json:"rule"`
	DeliveryInfos []OrderCycleDeliveryInfo `json:"delivery_infos"`
}

type OrderCycleDeliveryInfo struct {
	Index                    int                `json:"index"`
	EstimateArrivalDate      int                `json:"estimate_arrival_date"`
	EstimateArrivalTimeStart int                `json:"estimate_arrival_time_start"`
	EstimateArrivalTimeEnd   int                `json:"estimate_arrival_time_end"`
	LogisticsStatus          int                `json:"logistics_status"`
	CycleDetail              []OrderCycleDetail `json:"cycle_detail"`
}

type OrderCycleDetail struct {
	AppSpuCode string `json:"app_spu_code"`
	SkuID      string `json:"sku_id"`
	Name       string `json:"name"`
	Quanity    int    `json:"quanity"`
}

type OrderRefundRecord struct {
	WmOrderIDView                     int64                           `json:"wm_order_id_view"`
	OrderID                           int64                           `json:"order_id"`
	RefundID                          int64                           `json:"refund_id"`
	Ctime                             int64                           `json:"ctime"`
	Utime                             int64                           `json:"utime"`
	RefundType                        int                             `json:"refund_type"`
	ResReason                         string                          `json:"res_reason"`
	ResType                           int                             `json:"res_type"`
	ApplyType                         int                             `json:"apply_type"`
	ApplyReason                       string                          `json:"apply_reason"`
	Money                             float64                         `json:"money"`
	Pictures                          []string                        `json:"pictures"`
	RefundPartialEstimateCharge       RefundParitialEstimateCharge    `json:"refund_partial_estimate_charge"`
	WmAppRetailForOrderPartRefundList []wmAppRetailForOrderPartRefund `json:"wmAppRetailForOrderPartRefundList"`
	ServiceType                       string                          `json:"service_type"`
	Status                            string                          `json:"status"`
	ApplyOpUserType                   string                          `json:"apply_op_user_type"`
	LogisticsInfo                     struct {
		Logistics
		Reason string `json:"reason"`
	} `json:"logistics_info"`
	IncmpCode                  int    `json:"incmp_code"`
	IncmpModules               int    `json:"incmp_modules"`
	UserRefundGoodsShippingFee string `json:"user_refund_goods_shipping_fee"`
}

type RefundParitialEstimateCharge struct {
	TotalFoodAmount       *string `json:"total_food_amount"`
	BoxAmount             *string `json:"box_amount"`
	ActivityPoiAmount     *string `json:"activity_poi_amount"`
	ActivityMeituanAmount *string `json:"activity_meituan_amount"`
	ActivityAgentAmount   *string `json:"activity_agent_amount"`
	PlatformChargeFee     *string `json:"platform_charge_fee"`
	SettleAmount          *string `json:"settle_amount"`
	Productpreferences    *string `json:"productpreferences"`
	NotProductpreferences *string `json:"not_productpreferences"`
}

type wmAppRetailForOrderPartRefund struct {
	RetailRefundPartialEstimateCharge struct {
		TotalOriginPrice           string `json:"total_origin_price"`
		TotalActivityPrice         string `json:"total_activity_price"`
		TotalReducePrice           string `json:"total_reduce_price"`
		TotalPoiCharge             string `json:"total_poi_charge"`
		OrderRetailActivityDetails []struct {
			OriginActivityID int64  `json:"origin_activity_id"`
			Type             int    `json:"type"`
			PoiCharge        string `json:"poi_charge"`
		} `json:"order_retail_activity_details"`
	} `json:"retail_refund_partial_estimate_charge"`
	AppSpuCode      string  `json:"app_spu_code"`
	FoodName        string  `json:"food_name"`
	SkuID           string  `json:"sku_id"`
	ItemID          string  `json:"item_id"`
	Upc             string  `json:"upc"`
	Spec            string  `json:"spec"`
	Count           int     `json:"count"`
	BoxNum          float64 `json:"box_num"`
	BoxPrice        float64 `json:"box_price"`
	FoodPrice       float64 `json:"food_price"`
	OriginFoodPrice float64 `json:"origin_food_price"`
	RefundPrice     float64 `json:"refund_price"`
	RefundedWeight  float64 `json:"refunded_weight"`
}
