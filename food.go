package meituan

// 创建或更新菜品分类
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/52
func (m MT) FoodCatUpdate(req FoodCatUpdateRequest) (StringResponse, error) {
	var data StringResponse
	_, err := m.processRquest("food/cat/update", PostMethod, nil, req, &data)
	return data, err
}

// 批量创建或更新菜品（新版）
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/59
func (m MT) FoodBatchInitData(req FoodBatchInitDataRequest) (StringResponse, error) {
	var data StringResponse
	_, err := m.processRquest("food/batchinitdata", PostMethod, nil, req, &data)
	return data, err
}

// 批量创建或更新菜品（餐饮用）
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/301
func (m MT) FoodBatchBulkSave(req FoodBatchBulkSaveRequest) (StringResponse, error) {
	var data StringResponse
	_, err := m.processRquest("food/batchbulk/save", PostMethod, nil, req, &data)
	return data, err
}

// 批量更新售卖状态
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/70
func (m MT) FoodSkuSellStatus(req FoodSkuSellStatusRequest) (StringResponse, error) {
	var data StringResponse
	_, err := m.processRquest("food/sku/sellStatus", PostMethod, nil, req, &data)
	return data, err
}

// 更新SKU价格
//
//	文档地址：https://open-shangou.meituan.com/home/docDetail/63
func (m MT) FoodSkuPrice(req FoodSkuPriceRequest) (StringResponse, error) {
	var data StringResponse
	_, err := m.processRquest("food/sku/price", PostMethod, nil, req, &data)
	return data, err
}

// 更新SKU库存
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/64
func (m MT) FoodSkuStock(req FoodSkuStockRequest) (StringResponse, error) {
	var data StringResponse
	_, err := m.processRquest("food/sku/stock", PostMethod, nil, req, &data)
	return data, err
}

// 增加SKU库存
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/65
func (m MT) FoodSkuIncStock(req FoodSkuIncStockRequest) (StringResponse, error) {
	var data StringResponse
	_, err := m.processRquest("food/sku/inc_stock", PostMethod, nil, req, &data)
	return data, err
}

// 减少SKU库存
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/66
func (m MT) FoodSkuDescStock(req FoodSkuDescStockRequest) (StringResponse, error) {
	var data StringResponse
	_, err := m.processRquest("food/sku/desc_stock", PostMethod, nil, req, &data)
	return data, err
}
