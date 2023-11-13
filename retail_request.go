// 闪购零售-商品类API
package meituan

// 文档地址：https://open-shangou.meituan.com/home/docDetail/71
type RetailCatUpdateRequest struct {
	// 创建/更新商品分类
	AppPoiCode            string  `json:"app_poi_code"`
	CategoryCodeOrigin    *string `json:"category_code_origin"`
	CategoryNameOrign     *string `json:"category_name_origin"`
	CategoryCode          *string `json:"category_code"`
	CategoryName          string  `json:"category_name"`
	SecondaryCategoryCode *string `json:"secondary_category_code"`
	SecondaryCategoryName *string `json:"secondary_category_name"`
	Sequence              *int    `json:"sequence"`
	TargetLevel           *int    `json:"target_level"`
	TargetParentName      *string `json:"target_parent_name"`
	TopFlag               *int    `json:"top_flag"`
	WeeksTime             *string `json:"weeks_time"`
	Period                *string `json:"period"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/72
type RetailCatDeleteRequest struct {
	// 删除商品分类
	AppPoiCode          string  `json:"app_poi_code"`
	CategoryCode        *string `json:"category_code"`
	CategoryName        *string `json:"category_name"`
	MoveProductToUncate *int    `json:"move_product_to_uncate"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/73
type RetailCatListRequest struct {
	// 查询门店商品分类列表
	AppPoiCode string `json:"app_poi_code"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/517
type RetailCatSetSequenceRequest struct {
	// 批量更新商品分类顺序及分类下商品智能排序开关
	AppPoiCode           string `json:"app_poi_code"`
	CategorySequenceData []struct {
		CategoryCode string `json:"category_code"`
		Sequence     *int   `json:"sequence"`
		SmartSwitch  *int   `json:"smart_switch"`
	} `json:"category_sequence_data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/74
type RetailInitdataRequest struct {
	// 创建/更新商品[支持商品多规格,不含删除逻辑]
	AppPoiCode  string `json:"app_poi_code"`
	OperateType *int   `json:"operate_type"`
	RetailData
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/76
type RetailBatchInitdataRequest struct {
	// 批量创建/更新商品[支持商品多规格,不含删除逻辑]
	AppPoiCode  string       `json:"app_poi_code"`
	OperateType *int         `json:"operate_type"`
	FoodData    []RetailData `json:"food_data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/271
type RetailSellStatusRequest struct {
	// 批量更新商品上/下架状态
	AppPoiCode string                 `json:"app_poi_code"`
	FoodData   []RetailSellStatusFood `json:"food_data"`
	SellStatus string                 `json:"sell_status"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/269
type RetailSkuSaveRequest struct {
	// 创建/更新SKU信息
	AppPoiCode string       `json:"app_poi_code"`
	AppSpuCode string       `json:"app_spu_code"`
	Skus       []RetailSkus `json:"skus"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/77
type RetailSkuPriceRequest struct {
	// 批量更新SKU价格
	AppPoiCode string               `json:"app_poi_code"`
	FoodData   []RetailSkuPriceFood `json:"food_data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/78
type RetailSkuStockRequest struct {
	// 批量更新SKU库存
	AppPoiCode string               `json:"app_poi_code"`
	FoodData   []RetailSkuStockFood `json:"food_data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/75
type RetailListRequest struct {
	// 查询门店商品列表
	AppPoiCode string `json:"app_poi_code"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/79
type RetailGetRequest struct {
	// 查询商品详情
	AppPoiCode string `json:"app_poi_code"`
	AppSpuCode string `json:"app_spu_code"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/276
type RetailGetSpTagIdsRequest struct {
	// 获取美团后台商品类目（末级类目id）
	AppPoiCode   *string `json:"app_poi_code"`
	CategoryType *int    `json:"category_type"`
}

// 文档地址£ohttps://open-shangou.meituan.com/home/docDetail/10100
type RetailFieldRequiredInfoRequest struct {
	// 根据末级类目id查询字段必填性
	TagID string `json:"tag_id"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/486
type RetailRecommendTagRequest struct {
	// 根据商品UPC或商品名称查询平台推荐类目信息
	UpcCode *string `json:"upc_code"`
	Name    *string `json:"name"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/272
type RetailDeleteRequest struct {
	// 删除商品
	AppPoiCode        string `json:"app_poi_code"`
	AppSpuCode        string `json:"app_spu_code"`
	IsDeleteRetailCat *int   `json:"is_delete_retail_cat"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/273
type RetailBindPropertyRequest struct {
	// 批量绑定商品属性
	AppPoiCode   string `json:"app_poi_code"`
	FoodProperty []struct {
		AppSpuCode string           `json:"app_spu_code"`
		Properties []RetailProperty `json:"properties"`
	} `json:"food_property"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/270
type RetailPropertyListRequest struct {
	// 获取商品属性
	RetailGetRequest
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/274
type RetailSkuDeleteRequest struct {
	// 批量删除SKU
	RetailGetRequest
	SkuID             string `json:"sku_id"`
	IsDeleteRetailCat *int   `json:"is_delete_retail_cat,omitempty"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/286
type RetailCatBatchdeleteCatandretailRequest struct {
	// 批量删除商品分类及商品
	AppPoiCode             string  `json:"app_poi_code"`
	CategoryCodes          *string `json:"category_codes,omitempty"`
	CategoryNames          *string `json:"category_names,omitempty"`
	SecondaryCategoryCodes *string `json:"secondary_category_codes,omitempty"`
	SecondaryCategoryNames *string `json:"secondary_category_names,omitempty"`
	AppSpuCodes            *string `json:"app_spu_codes,omitempty"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/380
type EcommerceVideoUploadRequest struct {
	// 上传商品视频
	AppPoiCodes string  `json:"app_poi_codes"`
	VideoName   *string `json:"video_name,omitempty"`
	VideoData   *[]byte `json:"video_data,omitempty"`
	VideoUrl    *string `json:"video_url,omitempty"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/381
type GwVideoListRequest struct {
	// 查询门店商品视频
	AppPoiCode string `json:"app_poi_code"`
	PageNum    int    `json:"page_num"`
	PageSize   int    `json:"page_size"`
}

type RetailSkus struct {
	SkuID                 string         `json:"sku_id"`
	Spec                  *string        `json:"spec,omitempty"`
	AvailableTimes        AvailableTimes `json:"available_times"`
	Unit                  *string        `json:"unit,omitempty"`
	MinOrderCount         *string        `json:"min_order_count,omitempty"`
	Price                 *string        `json:"price,omitempty"`
	Stock                 *string        `json:"stock,omitempty"`
	Upc                   *string        `json:"upc,omitempty"`
	LocationCode          *string        `json:"location_code,omitempty"`
	Weight                *string        `json:"weight,omitempty"`
	WeightForUnit         *string        `json:"weight_for_unit,omitempty"`
	WeightUnit            *string        `json:"weight_unit,omitempty"`
	BoxNum                *string        `json:"box_num,omitempty"`
	BoxPrice              *string        `json:"box_price,omitempty"`
	LadderBoxNum          *string        `json:"ladder_box_num,omitempty"`
	LadderBoxPrice        *string        `json:"ladder_box_price,omitempty"`
	IsSellFlag            *int           `json:"is_sell_flag,omitempty"`
	LimitOpenSyncStockNow *bool          `json:"limit_open_sync_stock_now,omitempty"`
	OpenSaleAttrValueList *[]struct {
		AttrID  int64  `json:"attrId"`
		ValueID int64  `json:"valueId"`
		Value   string `json:"value"`
	} `json:"openSaleAttrValueList,omitempty"`
}

type RetailData struct {
	AppPoiCode            *string                  `json:"app_poi_code,omitempty"`
	AppSpuCode            string                   `json:"app_spu_code"`
	Name                  *string                  `json:"name,omitempty"`
	Description           *string                  `json:"description,omitempty"`
	StandardUpc           *string                  `json:"standard_upc,omitempty"`
	UseStandardCate       *int                     `json:"use_standard_cate,omitempty"`
	Skus                  *[]RetailSkus            `json:"skus,omitempty"`
	Price                 *float64                 `json:"price,omitempty"`
	MinOrderCount         *int                     `json:"min_order_count,omitempty"`
	Unit                  *string                  `json:"unit,omitempty"`
	BoxNum                *float64                 `json:"box_num,omitempty"`
	BoxPrice              *float64                 `json:"box_price,omitempty"`
	CategoryName          *string                  `json:"category_name,omitempty"`
	CategoryCodeList      *[]string                `json:"category_code_list,omitempty"`
	CategoryNameList      *[]string                `json:"category_name_list,omitempty"`
	IsSoldOut             *int                     `json:"is_sold_out,omitempty"`
	Picture               *string                  `json:"picture,omitempty"`
	Sequence              *int                     `json:"sequence,omitempty"`
	TagID                 *int64                   `json:"tag_id,omitempty"`
	ZhID                  *int64                   `json:"zh_id,omitempty"`
	OriginName            *string                  `json:"origin_name,omitempty"`
	PictureContents       *string                  `json:"picture_contents,omitempty"`
	Properties            *[]RetailProperties      `json:"properties,omitempty"`
	IsSpecialty           *int                     `json:"is_specialty,omitempty"`
	IsSingleNoDelivery    *int                     `json:"is_single_no_delivery,omitempty"`
	VideoID               *int64                   `json:"video_id,omitempty"`
	CommonAttrValue       *[]RetailCommonAttrValue `json:"common_attr_value,omitempty"`
	IsShowUpcPicContents  *int                     `json:"is_show_upc_pic_contents,omitempty"`
	LimitSaleInfo         *RetailLimitSaleInfo     `json:"limit_sale_info,omitempty"`
	UpcImage              *string                  `json:"upc_image,omitempty"`
	SellPoint             *string                  `json:"sell_point,omitempty"`
	SaleType              *int                     `json:"sale_type,omitempty"`
	DeliveryTime          *int                     `json:"delivery_time,omitempty"`
	DeliveryEndTime       *int                     `json:"delivery_end_time,omitempty"`
	AvailableTimes        *AvailableTimes          `json:"available_times,omitempty"`
	AppointAvailableTimes *string                  `json:"appoint_available_times,omitempty"`
	AllAvailableTimes     *string                  `json:"all_available_times,omitempty"`
	QuaPictures           *string                  `json:"qua_pictures,omitempty"`
	QuaEffectiveDate      *int                     `json:"qua_effective_date,omitempty"`
	QuaApprovalDate       *int                     `json:"qua_approval_date,omitempty"`
	ExpireDate            *string                  `json:"expire_date,omitempty"`
}

type RetailCommonAttrValue struct {
	AttrID    int64  `json:"attrId"`
	AttrName  string `json:"attrName"`
	ValueList []struct {
		ValueID int64  `json:"valueId"`
		Value   string `json:"value"`
	} `json:"value"`
}

type RetailLimitSaleInfo struct {
	LimitSale bool    `json:"limitSale"`
	Type      *int    `json:"type,omitempty"`
	Frequency *int    `json:"frequency,omitempty"`
	Begin     *string `json:"begin,omitempty"`
	End       *string `json:"end,omitempty"`
	Count     *int    `json:"count,omitempty"`
}

type RetailSellStatusFood struct {
	AppSpuCode string `json:"app_spu_code"`
}

type RetailSkuPriceFood struct {
	AppSpuCode string `json:"app_spu_code"`
	Skus       []struct {
		SkuID string `json:"sku_id"`
		Price string `json:"price"`
	} `json:"skus"`
}

type RetailSkuStockFood struct {
	AppSpuCode string `json:"app_spu_code"`
	Skus       []struct {
		SkuID string `json:"sku_id"`
		Stock string `json:"stock"`
	} `json:"skus"`
}

type RetailProperties struct {
	PropertyName string `json:"property_name"`
	Values       string `json:"values"`
}
