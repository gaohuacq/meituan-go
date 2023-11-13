// 门店类api
package meituan

// 获取门店品类列表
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/29
func (m MT) TagList() (data TagListResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("poiTag/list", GetMethod, nil, nil, &data)
	return
}

// 更改门店公告信息
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/26
func (m MT) PoiUpdatePromotionInfo(req PoiUpdatePromotionInfoRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("poi/updatepromoteinfo", PostMethod, nil, req, &data)
	return
}

// 更新门店信息
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/2
func (m MT) PoiSave(req PoiSaveRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("poi/save", PostMethod, nil, req, &data)
	return
}

// 批量获取门店详细信息
//
// 文档地址：https://developer.waimai.meituan.com/home/docDetail/7
func (m MT) PoiMget(req PoiMgetRequest) (data PoiMgetResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("poi/mget", GetMethod, req, nil, &data)
	return
}

// 门店设置为上线状态
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/20
func (m MT) PoiOnline(req PoiCommonRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("poi/online", GetMethod, req, nil, &data)
	return
}

// 门店设置为下线状态
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/17
func (m MT) PoiOffline(req PoiCommonRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("poi/offline", GetMethod, req, nil, &data)
	return
}

// 门店设置为休息状态
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/14
func (m MT) PoiClose(req PoiCommonRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("poi/close", GetMethod, req, nil, &data)
	return
}

// 门店设置为营业状态
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/11
func (m MT) PoiOpen(req PoiCommonRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("poi/open", GetMethod, req, nil, &data)
	return
}

// 获取应用已绑定门店的三方门店ID
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/10083
func (m MT) PoiGetIDs() (data SliceStringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("poi/getids", GetMethod, nil, nil, &data)
	return
}

// 更新门店营业时间
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/32
func (m MT) PoiShippingTimeUpdate(req PoiShippingTimeUpdateRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("poi/shippingtime/update", PostMethod, nil, req, &data)
	return
}
