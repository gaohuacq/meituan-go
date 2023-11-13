package meituan

// 设订单为商家已收到
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/107
func (m MT) OrderPoiReceived(req OrderBaseRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/poi_received", GetMethod, nil, req, &data)
	return
}

// 获取订单详细信息
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/134
func (m MT) OrderDetail(req OrderDetailRequest) (data OrderDetailResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/getOrderDetail", GetMethod, req, nil, &data)
	return
}

// 商家确认订单（必接）
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/110
func (m MT) OrderConfirm(req OrderBaseRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/confirm", GetMethod, nil, req, &data)
	return
}

// 商家取消订单（必接）
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/113
func (m MT) OrderCancel(req OrderCancelRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/cancel", GetMethod, nil, req, &data)
	return
}

// 订单已送达
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/118
func (m MT) OrderDelivered(req OrderBaseRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/delivered", GetMethod, nil, req, &data)
	return
}

// 催单回复接口
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/166
func (m MT) OrderRemindReply(req OrderRemindReplyRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/remindReply", PostMethod, nil, req, &data)
	return
}

// 订单确认退款请求
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/121
func (m MT) OrderRefundAgree(req OrderRefundRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/refund/agree", GetMethod, req, nil, &data)
	return
}

// 驳回订单退款申请
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/125
func (m MT) OrderRefundReject(req OrderRefundRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/refund/reject", GetMethod, req, nil, &data)
	return
}

// 申请餐损赔付
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/227
func (m MT) OrderApplyCompensation(req OrderApplyCompensationRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/applyCompensation", PostMethod, nil, req, &data)
	return
}

// 批量拉取异常订单
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/252
func (m MT) OrderBatchFetchAbnormalOrder(req OrderBatchFetchAbnormalOrderRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/batchFetchAbnormalOrder", PostMethod, nil, req, &data)
	return
}

// 自配送商家同步发货状态和配送信息（推荐）
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/351
func (m MT) EcommerceOrderLogisticsSync(req EcommerceOrderLogisticsSyncRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("ecommerce/order/logistics/sync", PostMethod, nil, req, &data)
	return
}

// 自配订单配送中（不支持同步骑手信息）
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/116
func (m MT) OrderdDelivering(req OrderBaseRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/delivering", GetMethod, req, nil, &data)
	return
}

// 自配订单已送达（不支持同步骑手信息）
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/119
func (m MT) OrderArrived(req OrderBaseRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/arrived", GetMethod, req, nil, &data)
	return
}

// 售后单（退款/退货退款）审核接口
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/410
func (m MT) EcommerceOrderReviewAfterSales(req EcommerceOrderReviewAfterSalesRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("ecommerce/order/reviewAfterSales", PostMethod, nil, req, &data)
	return
}

// 查询订单状态
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/128
func (m MT) OrderViewStatus(req OrderBaseRequest) (data OrderViewStatusResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/viewStatus", GetMethod, req, nil, &data)
	return
}

// 根据流水号获取订单ID（分段）
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/341
func (m MT) EcommerceGetOrderIdByDaySeq(req EcommerceGetOrderIdByDaySeqRequest) (data EcommerceGetOrderIdByDaySeqResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("ecommerce/order/getOrderIdByDaySeq", GetMethod, req, nil, &data)
	return
}

// 获取当日最新订单流水号
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/146
func (m MT) GetOrderDaySeq(req GetOrderDaySeqRequest) (data GetOrderDaySeqResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/getOrderDaySeq", GetMethod, req, nil, &data)
	return
}

// 根据流水号获取订单ID
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/149
func (m MT) GetOrderIdByDaySeq(req GetOrderIdByDaySeqRequest) (data GetOrderIdByDaySeqResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/getOrderIdByDaySeq", GetMethod, req, nil, &data)
	return
}

// 获取订单退款记录
//
// 文档地址£ohttps://open-shangou.meituan.com/home/docDetail/321
func (m MT) EcommerceOrderGetOrderRefundDetail(req EcommerceOrderGetOrderRefundDetailRequest) (data EcommerceOrderGetOrderRefundDetailResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("ecommerce/order/getOrderRefundDetail", GetMethod, req, nil, &data)
	return
}

// 商家确认已完成拣货
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/216
func (m MT) OrderPreparationMealComplete(req OrderBaseRequest) (data StringResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/preparationMealComplete", GetMethod, req, nil, &data)
	return
}

// 商家获取备货时长
//
// 文档地址：https://open-shangou.meituan.com/home/docDetail/219
func (m MT) OrderGetPreparationMealTime(req OrderBaseRequest) (data OrderGetPreparationMealTimeResponse, errResp *ErrorResponse, err error) {
	errResp, err = m.processRquest("order/preparationMealTime", GetMethod, req, nil, &data)
	return
}
