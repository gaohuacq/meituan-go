package meituan

type TagListResponse struct {
	BaseResponse
	Data []TagInfo
}

type TagInfo struct {
	Name string
	ID   int
}

// 文档地址： https://developer.waimai.meituan.com/home/docDetail/7
type PoiMgetResponse struct {
	// 批量获取门店详细信息
	BaseResponse
	Data []PoiInfoResponse `json:"data"`
}

type PoiInfoResponse struct {
	AppID              int     `json:"app_id"`
	AppPoiCode         string  `json:"app_poi_code"`
	CityID             int     `json:"city_id"`
	Name               string  `json:"name"`
	Address            string  `json:"address"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	PicUrl             string  `json:"pic_url"`
	PicUrlLarge        *string `json:"pic_url_large,omitempty"`
	Phone              string  `json:"phone"`
	StandbyTel         string  `json:"standby_tel"`
	ShippingFee        float64 `json:"shipping_fee"`
	ShippingTime       string  `json:"shipping_time"`
	PromotionInfo      string  `json:"promotion_info"`
	InvoiceSupport     int     `json:"invoice_support"`
	InvoiceMinPrice    int     `json:"invoice_min_price"`
	InvoiceDescription string  `json:"invoice_description"`
	OpenLevel          int     `json:"open_level"`
	IsOnline           int     `json:"is_online"`
	Ctime              int64   `json:"ctime"`
	Utime              int64   `json:"utime"`
	TagName            string  `json:"tag_name"`
	ThirdTagName       *string `json:"third_tag_name,omitempty"`
	PreBook            int     `json:"pre_book"`
	TimeSelect         int     `json:"time_select"`
	LogisticsCodes     string  `json:"logistics_codes"`
	LocationID         int     `json:"location_id"`
	PreBookMinDays     int     `json:"pre_book_min_days"`
	PreBookMaxDays     int     `json:"pre_book_max_days"`
	Remark             string  `json:"remark"`
}
