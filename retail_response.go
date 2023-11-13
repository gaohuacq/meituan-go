package meituan

import (
	"encoding/json"
)

type RetailBaseResponse struct {
	ResultCode int `json:"result_code"`
}

func (r RetailBaseResponse) GetResultCode() int {
	return 1
}

func (r RetailBaseResponse) GetErrorList() []ErrorData {
	return []ErrorData{}
}

func (r RetailBaseResponse) GetError() Error {
	return Error{}
}

func (r RetailBaseResponse) GetErrorString() string {
	return r.GetError().Msg
}

type RetailBaseMsg struct {
	Success *[]RetailBaseInfo `json:"success,omitempty"`
	Error   *[]RetailBaseInfo `json:"error,omitempty"`
}

type RetailBaseInfo struct {
	AppSpuCode string `json:"app_spu_code"`
	Name       string `json:"name"`
	Skus       []struct {
		SkuID string  `json:"sku_id"`
		Price float64 `json:"price"`
	} `json:"skus"`
	UpcCode string `json:"upc_code"`
}

type RetailStringResponse struct {
	RetailBaseResponse
	Data string `json:"data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/73
type RetailCatListResponse struct {
	RetailBaseResponse
	Data []RetailCatData `json:"data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/75
type RetailListResponse struct {
	RetailBaseResponse
	Data      []RetailDataResponse `json:"data"`
	ExtraInfo struct {
		TotalCount int `json:"total_count"`
	} `json:"extra_info"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/79
type RetailGetResponse struct {
	RetailBaseResponse
	Data RetailDataResponse `json:"data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/276
type RetailGetSpTagIdsResponse struct {
	RetailBaseResponse
	Data []RetailSpTag `json:"data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/10100
type RetailFieldRequiredInfoResponse struct {
	RetailBaseResponse
	Data struct {
		Upc    int `json:"upc"`
		Weight int `json:"weight"`
	} `json:"success_map"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/486
type RetailRecommendTagResponse struct {
	RetailBaseResponse
	Data UpcData `json:"success_map"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/270
type RetailPropertyListResponse struct {
	RetailBaseResponse
	Data []RetailProperty `json:"data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/380
type EcommerceVideoUploadResponse struct {
	RetailBaseResponse
	Data []ProductVideo `json:"data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/381
type GwVideoListResponse struct {
	RetailBaseResponse
	Data []ProductVideo `json:"data"`
}

type RetailResponse struct {
	RetailBaseResponse
	Data       string          `json:"data"`
	ExtraInfo  *RetailTaskInfo `json:"extra_info,omitempty"`
	SuccessMap *RetailTaskInfo `json:"success_map,omitempty"`
	ErrorList  *[]RetailError  `json:"error_list,omitempty"`
}

type RetailCatData struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Sequence    int    `json:"sequence"`
	TopFlag     int    `json:"top_flag"`
	WeeksTime   string `json:"weeks_time"`
	Period      string `json:"period"`
	SmartSwitch int    `json:"smart_switch"`
	Children    string `json:"children"`
}

type RetailTaskInfo struct {
	TaskID         string `json:"task_id"`
	TaskDetailLink string `json:"task_detail_link"`
}

type RetailError struct {
	AppSpuCode string `json:"app_spu_code"`
	SkuID      string `json:"sku_id"`
	Msg        string `json:"msg"`
	Code       int    `json:"code"`
}

type RetailSpTag struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Level    int    `json:"level"`
	NamePath string `json:"name_path"`
}

type UpcData struct {
	UpcCode string `json:"upc_code"`
	Name    string `json:"name"`
	TagID   int64  `json:"tag_id"`
}

type RetailDataResponse struct {
	AppPoiCode            string                  `json:"app_poi_code"`
	AppSpuCode            string                  `json:"app_spu_code"`
	AppointAvailableTimes string                  `json:"appoint_available_times"`
	AduitStatus           int                     `json:"aduit_status,string"`
	AvailableTimes        AvailableTimes          `json:"available_times"`
	BoxNum                float64                 `json:"box_num"`
	BoxPrice              float64                 `json:"box_price"`
	CategoryCode          string                  `json:"category_code"`
	CategoryList          *[]RetailCategory       `json:"category_list,omitempty"`
	CategoryName          string                  `json:"category_name"`
	CommonAttrValue       []RetailCommonAttrValue `json:"common_attr_value"`
	Ctime                 int64                   `json:"ctime"`
	Description           string                  `json:"description"`
	IsSp                  int                     `json:"is_sp"`
	IsComplete            int                     `json:"is_complete"`
	IsShowUpcPicContents  int                     `json:"is_show_upc_pic_contents"`
	IsSingleNoDelivery    int                     `json:"is_single_no_delivery"`
	IsSoldOut             int                     `json:"is_sold_out"`
	IsSpecialty           int                     `json:"is_specialty"`
	MinOrderCount         int                     `json:"min_order_count"`
	Name                  string                  `json:"name"`
	NormAduitStatus       int                     `json:"norm_aduit_status"`
	NormAduitType         int                     `json:"norm_aduit_type"`
	Picture               string                  `json:"picture"`
	PictureContents       string                  `json:"picture_contents"`
	Price                 float64                 `json:"price"`
	Properties            []RetailProperties      `json:"properties"`
	SaleType              int                     `json:"sale_type"`
	Sequence              int                     `json:"sequence"`
	Skus                  []RetailSkus            `json:"skus"`
	SpID                  int                     `json:"spId"`
	StockConfig           RetailStockConfig       `json:"stock_config"`
	TagID                 int64                   `json:"tag_id"`
	Unit                  string                  `json:"unit"`
	UpcCode               string                  `json:"upc_code"`
	Utime                 int64                   `json:"utime"`
	VideoID               int                     `json:"video_id"`
	ZhName                string                  `json:"zh_name"`
}

func (pc *RetailDataResponse) UnmarshalJSON(p []byte) (err error) {
	if p == nil || string(p) == `""` {
		return nil
	}

	type x RetailDataResponse
	if err = json.Unmarshal(escapeJsonCharactor(p), (*x)(pc)); err != nil {
		return
	}

	return
}

type RetailCategory struct {
	CategoryCode string `json:"category_code"`
	CategoryName string `json:"category_name"`
}

type RetailStockConfig struct {
	LimitOpenSyncStock bool   `json:"limit_open_sync_stock"`
	Reset              bool   `json:"reset"`
	Schedule           string `json:"schedule"`
	SyncCount          int    `json:"sync_count"`
	SyncNextDay        bool   `json:"sync_next_day"`
	Type               []int  `json:"type"`
}

type RetailProperty struct {
	PropertyName string   `json:"property_name"`
	Values       []string `json:"values"`
}

type ProductVideo struct {
	AppPoiCode  *string   `json:"app_poi_code,omitempty"`
	VideoID     *int64    `json:"video_id,omitempty"`
	VideoUrlMp4 *string   `json:"video_url_mp4,omitempty"`
	VideoName   *string   `json:"video_name,omitempty"`
	AppSpuCodes *[]string `json:"app_spu_codes,omitempty"`
	VideoState  *int      `json:"video_state,omitempty"`
	VideoLength *int64    `json:"video_length,omitempty"`
	VideoSize   *float64  `json:"video_size,omitempty"`
}
