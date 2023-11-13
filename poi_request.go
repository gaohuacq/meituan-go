package meituan

type PoiCommonRequest struct {
	// 门店操作通用结构体
	AppPoiCode string `json:"app_poi_code"`
}

type PoiMgetRequest struct {
	AppPoiCodes string `json:"app_poi_codes"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/26
type PoiUpdatePromotionInfoRequest struct {
	// 更新门店公告信息
	AppPoiCode    string `json:"app_poi_code"`
	PromotionInfo string `json:"promotion_info"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/2
type PoiSaveRequest struct {
	// 更新门店信息
	AppPoiCode         string   `json:"app_poi_code"`
	Name               *string  `json:"name,omitempty"`
	Address            *string  `json:"address,omitempty"`
	Latitude           *float64 `json:"latitude,omitempty"`
	Longitude          *float64 `json:"longitude,omitempty"`
	PicUrl             *string  `json:"pic_url,omitempty"`
	PicUrlLarge        *string  `json:"pic_url_large,omitempty"`
	Phone              *string  `json:"phone,omitempty"`
	StandbyTel         *string  `json:"standby_tel,omitempty"`
	ShippingFee        *float64 `json:"shipping_fee,omitempty"`
	ShippingTime       *string  `json:"shipping_time,omitempty"`
	PromotionInfo      *string  `json:"promotion_info,omitempty"`
	OpenLevel          *int     `json:"open_level,omitempty"`
	IsOnline           *int     `json:"is_online,omitempty"`
	InvoiceSupport     *int     `json:"invoice_support,omitempty"`
	InvoiceMinPrice    *int     `json:"invoice_min_price,omitempty"`
	InvoiceDescription *string  `json:"invoice_description,omitempty"`
	ThirdTagName       *string  `json:"third_tag_name,omitempty"`
	PreBook            *int     `json:"pre_book,omitempty"`
	TimeSelect         *int     `json:"time_select,omitempty"`
	AppBrandCode       *string  `json:"app_brand_code,omitempty"`
	PreBookMinDays     *int     `json:"pre_book_min_days,omitempty"`
	PreBookMaxDays     *int     `json:"pre_book_max_days,omitempty"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/32
type PoiShippingTimeUpdateRequest struct {
	// 更新门店营业时间
	AppPoiCode   string `json:"app_poi_code"`
	ShippingTime string `json:"shipping_time"`
}
