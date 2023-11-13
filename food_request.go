package meituan

// 文档地址：https://open-shangou.meituan.com/home/docDetail/52
type FoodCatUpdateRequest struct {
	// 创建或更新菜品分类
	AppPoiCode          string                   `json:"app_poi_code"`
	CategoryNameOrigin  *string                  `json:"category_name_origin,omitempty"`
	CategoryName        string                   `json:"category_name"`
	Sequence            *int                     `json:"sequence,omitempty"`
	CategoryDescription *string                  `json:"category_description,omitempty"`
	CategoryMode        *int                     `json:"category_mode,omitempty"`
	TopFlag             *int                     `json:"top_flag,omitempty"`
	TimeZone            *map[int]FoodCatTimeZone `json:"time_zone,omitempty"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/59
type FoodBatchInitDataRequest struct {
	// 批量创建或更新菜品（新版）
	AppPoiCode   string   `json:"app_poi_code"`
	SkuOverwrite *bool    `json:"sku_overwrite,omitempty"`
	FoodData     FoodData `json:"food_data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/301
type FoodBatchBulkSaveRequest struct {
	// 批量创建或更新菜品（餐饮用）
	AppPoiCode string   `json:"app_poi_code"`
	FoodData   FoodData `json:"food_data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/70
type FoodSkuSellStatusRequest struct {
	// 更新售卖状态
	AppPoiCode string               `json:"app_poi_code"`
	FoodData   []SellStatusFoodData `json:"food_data"`
	SellStatus int                  `json:"sell_status"`
}

// 文档地址: https://open-shangou.meituan.com/home/docDetail/63
type FoodSkuPriceRequest struct {
	// 更新SKU价格
	AppPoiCode string             `json:"app_poi_code"`
	FoodData   []SkuPriceFoodData `json:"food_data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/64
type FoodSkuStockRequest struct {
	// 更新SKU库存
	AppPoiCode string             `json:"app_poi_code"`
	FoodData   []SkuStockFoodData `json:"food_data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/65
type FoodSkuIncStockRequest struct {
	// 增加SKU库存
	AppPoiCode string             `json:"app_poi_code"`
	FoodData   []SkuStockFoodData `json:"food_data"`
}

// 文档地址：https://open-shangou.meituan.com/home/docDetail/66
type FoodSkuDescStockRequest struct {
	// 减少SKU库存
	AppPoiCode string             `json:"app_poi_code"`
	FoodData   []SkuStockFoodData `json:"food_data"`
}

type FoodCatTimeZone struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type FoodData struct {
	AppFoodCode     string   `json:"app_food_code"`
	Name            *string  `json:"name,omitempty"`
	Description     *string  `json:"description,omitempty"`
	SpuAttr         *SpuAttr `json:"spu_attr,omitempty"`
	Skus            *Skus    `json:"skus,omitempty"`
	Price           *float64 `json:"price,omitempty"`
	MinOrderCount   *int     `json:"min_order_count,omitempty"`
	Unit            *string  `json:"unit,omitempty"`
	BoxNum          *float64 `json:"box_num,omitempty"`
	BoxPrice        *float64 `json:"box_price,omitempty"`
	CategoryName    *string  `json:"category_name,omitempty"`
	IsSoldOut       *int     `json:"is_sold_out,omitempty"`
	Picture         *string  `json:"picture,omitempty"`
	Sequence        *int     `json:"sequence,omitempty"`
	Pictures        *string  `json:"pictures,omitempty"`
	Speciality      *int     `json:"speciality,omitempty"`
	IsNotSingle     *int     `json:"is_not_single,omitempty"`
	OnlySellInCombo *bool    `json:"only_sell_in_combo,omitempty"`
}

type SpuAttr struct {
	No            *int           `json:"no,omitempty"`
	Mode          *int           `json:"mode,omitempty"`
	Naem          *string        `json:"name,omitempty"`
	Value         *string        `json:"value,omitempty"`
	Price         *float64       `json:"price,omitempty"`
	SellStatus    *int           `json:"sell_status,omitempty"`
	ValueSequence *int           `json:"value_sequence,omitempty"`
	ExcludeAttr   *[]ExcludeAttr `json:"exclude_attr,omitempty"`
}

type ExcludeAttr struct {
	AttrName   *string   `json:"attr_name,omitempty"`
	AttrValues *[]string `json:"attr_values,omitempty"`
}

type Skus struct {
	SkuID          string          `json:"sku_id"`
	Spec           string          `json:"spec"`
	Upc            *string         `json:"upc,omitempty"`
	Price          string          `json:"price"`
	Stock          string          `json:"stock"`
	SkuSequence    *int            `json:"sku_sequence,omitempty"`
	AvailableTimes *AvailableTimes `json:"available_times,omitempty"`
	LocationCode   *string         `json:"location_code,omitempty"`
	BoxNum         *string         `json:"box_num,omitempty"`
	BoxPrice       *string         `json:"box_price,omitempty"`
	LadderNum      *string         `json:"ladder_num,omitempty"`
	LadderPrice    *string         `json:"ladder_price,omitempty"`
	Weight         *int64          `json:"weight,omitempty"`
	WeightUnit     *string         `json:"weight_unit,omitempty"`
	SkuAttr        *SkuAttr        `json:"sku_attr,omitempty"`
}

type SkuAttr struct {
	SkuID *string `json:"sku_id,omitempty"`
	No    *int    `json:"no,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type SellStatusFoodData struct {
	AppFoodCode string   `json:"app_food_code"`
	Skus        []string `json:"skus"`
}

type SkuPriceFoodData struct {
	AppFoodCode string                 `json:"app_food_code"`
	Skus        []SkuPriceFoodDataSkus `json:"skus"`
}

type SkuPriceFoodDataSkus struct {
	SkuID string  `json:"sku_id"`
	Price float64 `json:"price"`
}

type SkuStockFoodData struct {
	AppFoodCode string                 `json:"app_food_code"`
	Skus        []SkuStockFoodDataSkus `json:"skus"`
}

type SkuStockFoodDataSkus struct {
	SkuID string `json:"sku_id"`
	Stock int    `json:"stock"`
}
