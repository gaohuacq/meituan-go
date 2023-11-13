// 闪购零售-商品类API
package meituan

// 创建/更新商品分类
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/71
func (m MT) RetailCatUpdate(req RetailCatUpdateRequest) (data RetailStringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retailCat/update", PostMethod, nil, req, &data)
	return
}

// 删除商品分类
//
// 文档地址：ohttps://open-shangou.meituan.com/home/docDetail/72
func (m MT) RetailCatDelete(req RetailCatDeleteRequest) (data RetailStringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retailCat/delete", PostMethod, nil, req, &data)
	return
}

// 查询门店商品分类列表
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/73
func (m MT) RetailCatList(req RetailCatListRequest) (data RetailCatListResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retailCat/list", GetMethod, req, nil, &data)
	return
}

// 批量更新商品分类顺序及分类下商品智能排序开关
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/517
func (m MT) RetailCatSetSequence(req RetailCatSetSequenceRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retailCat/setSequence", PostMethod, nil, req, &data)
	return
}

// 创建/更新商品[支持商品多规格,不含删除逻辑]
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/74
func (m MT) RetailInitdata(req RetailInitdataRequest) (data RetailResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/initdata", PostMethod, nil, req, &data)
	return
}

// 批量创建/更新商品[支持商品多规格,不含删除逻辑]
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/76
func (m MT) RetailBatchInitdata(req RetailBatchInitdataRequest) (data RetailResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/batchinitdata", PostMethod, nil, req, &data)
	return
}

// 批量更新商品上/下架状态
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/271
func (m MT) RetailSellStatus(req RetailSellStatusRequest) (data RetailResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/sellStatus", PostMethod, nil, req, &data)
	return
}

// 创建/更新SKU信息
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/269
func (m MT) RetailSkuSave(req RetailSkuSaveRequest) (data RetailResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/sku/save", PostMethod, nil, req, &data)
	return
}

// 批量更新SKU价格
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/77
func (m MT) RetailSkuPrice(req RetailSkuPriceRequest) (data RetailResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/sku/price", PostMethod, nil, req, &data)
	return
}

// 批量更新SKU库存
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/78
func (m MT) RetailSkuStock(req RetailSkuStockRequest) (data RetailResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/sku/stock", PostMethod, nil, req, &data)
	return
}

// 查询门店商品列表
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/75
func (m MT) RetailList(req RetailListRequest) (data RetailListResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/list", GetMethod, req, nil, &data)
	return
}

// 查询商品详情
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/79
func (m MT) RetailGet(req RetailGetRequest) (data RetailGetResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/get", GetMethod, req, nil, &data)
	return
}

// 获取美团后台商品类目（末级类目id）
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/276
func (m MT) RetailGetSpTagIds(req RetailGetSpTagIdsRequest) (data RetailGetSpTagIdsResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/getSpTagIds", GetMethod, req, nil, &data)
	return
}

// 根据末级类目id查询字段必填性
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/10100
func (m MT) RetailFieldRequiredInfo(req RetailFieldRequiredInfoRequest) (data RetailFieldRequiredInfoResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/field/required/info", GetMethod, req, nil, &data)
	return
}

// 根据商品UPC或商品名称查询平台推荐类目信息
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/486
func (m MT) RetailRecommendTag(req RetailRecommendTagRequest) (data RetailRecommendTagResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/recommend/tag", GetMethod, req, nil, &data)
	return
}

// 删除商品
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/272
func (m MT) RetailDelete(req RetailDeleteRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/delete", PostMethod, nil, req, &data)
	return
}

// 批量绑定商品属性
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/273
func (m MT) RetailBindProperty(req RetailBindPropertyRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/bind/property", PostMethod, nil, req, &data)
	return
}

// 获取商品属性
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/270
func (m MT) RetailPropertyList(req RetailPropertyListRequest) (data RetailPropertyListResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/property/list", GetMethod, req, nil, &data)
	return
}

// 删除SKU信息
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/274
func (m MT) RetailSkuDelete(req RetailSkuDeleteRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retail/sku/delete", PostMethod, nil, req, &data)
	return
}

// 批量删除商品分类及商品
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/286
func (m MT) RetailCatBatchdeleteCatandretail(req RetailCatBatchdeleteCatandretailRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("retailCat/batchdelete/catandretail", PostMethod, nil, req, &data)
	return
}

// 上传商品视频
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/380
func (m MT) EcommerceVideoUpload(req EcommerceVideoUploadRequest) (data EcommerceVideoUploadResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("ecommerce/video/upload", PostMethod, nil, req, &data)
	return
}

// 查询门店商品视频
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/381
func (m MT) GwVideoList(req GwVideoListRequest) (data GwVideoListResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("gw/video/list", GetMethod, req, nil, &data)
	return
}
