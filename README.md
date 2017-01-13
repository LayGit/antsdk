# antsdk
蚂蚁金服(支付宝)开放平台 go-sdk
支付宝API文档:[传送门](https://doc.open.alipay.com/doc2/apiList?docType=4)

## 安装
```bash
go get github.com/LayGit/antsdk
```

## 使用示例

```go
import (
  "fmt"
  "github.com/LayGit/antsdk/alipay"
  "github.com/LayGit/antsdk/api/trade"
)

func main() {
    client := alipay.NewDefaultAlipayClient("https://openapi.alipay.com/gateway.do", "商户AppId", "商户密钥", "支付宝公钥")
    // 创建请求
    request := &trade.AlipayTradeQueryRequest{}
    // 设置参数
    request.BizContent.OutTradeNo = "L123456"
    // 请求响应
    var response trade.AlipayTradeQueryResponse
    err := client.Execute(request, &response)
    if err != nil {
        // 错误处理
        fmt.Println(err)
    } else {
        if response.IsSuccess() {
          fmt.Println("调用成功")
        } else {
          fmt.Println("调用失败")
        }
    }
}
```

## API 支持情况
- [x] [支付 API](#支付-api)
- [x] [会员 API](#会员-api)
- [x] [店铺 API](#店铺-api)
- [ ] 营销 API
- [x] [服务窗 API](#服务窗-api)
- [x] [芝麻信用 API](#芝麻信用-api)
- [x] [工具类 API](#工具类-api)
- [x] [风险控制 API](#风险控制-api)
- [x] [服务市场 API](#服务市场-api)
- [x] [账务 API](#账务-api)
- [x] [生活缴费 API](#生活缴费-api)
- [x] [车主服务 API](#车主服务-api)
- [x] [数据服务 API](#数据服务-api)
- [x] [卡券 API](#卡券-api)
- [x] [广告 API](#广告-api)
- [x] [地铁购票 API](#地铁购票-api)
- [x] [理财 API](#理财-api)

## 支付 API

描述 | API
---|---
统一收单交易撤销接口 | trade.AlipayTradeCancelRequest
统一收单交易关闭接口 | trade.AlipayTradeCloseRequest
统一收单交易创建  | trade.AlipayTradeCreateRequest
统一收单交易退款查询 | trade.AlipayTradeFastpayRefundQueryRequest
统一收单交易结算接口 | trade.AlipayTradeOrderSettleRequest
统一收单交易支付 | trade.AlipayTradePayRequest
统一收单线下交易预创建 | trade.AlipayTradePreCreateRequest
统一收单线下交易查询 | trade.AlipayTradeQueryRequest
统一收单交易退款 | trade.AlipayTradeRefundRequest

## 会员 API

描述 | API
---|---
支付宝钱包用户信息共享 | user.AlipayUserUserinfoShareRequest
支付宝会员授权信息查询接口 | user.AlipayUserInfoShareRequest

## 店铺 API

描述 | API
---|---
创建门店信息 | shop.AlipayOfflineMarketShopCreateRequest
修改门店信息 | shop.AlipayOfflineMarketShopModifyRequest
查询商户的门店编号列表 | shop.AlipayOfflineMarketShopBatchqueryRequest
查询单个门店信息接口 | shop.AlipayOfflineMarketShopQuerydetailRequest
上传门店照片和视频接口 | shop.AlipayOfflineMaterialImageUploadRequest
业务流水批量查询接口 | shop.AlipayOfflineMarketApplyorderBatchqueryRequest
门店摘要信息批量查询接口 | shop.AlipayOfflineMarketShopSummaryBatchqueryRequest
门店类目配置查询接口 | shop.AlipayOfflineMarketShopCategoryQueryRequest

## 服务窗 API
描述 | API
---|---
模板消息行业设置修改接口 | open.AlipayOpenPublicTemplateMessageIndustryModifyRequest
异步单发消息 | open.AlipayOpenPublicMessageCustomSendRequest
查询绑定商户会员号 | open.AlipayOpenPublicAccountQueryRequest
解除绑定商户会员号 | open.AlipayOpenPublicAccountDeleteRequest
添加绑定商户会员号 | open.AlipayOpenPublicAccountCreateRequest
服务窗基础信息查询接口 | open.AlipayOpenPublicInfoQueryRequest
单发模板消息 | open.AlipayOpenPublicMessageSingleSendRequest
查询服务窗联系人关注列表 | open.AlipayOpenPublicContactFollowBatchqueryRequest
服务窗基础信息修改接口 | open.AlipayOpenPublicInfoModifyRequest
公众服务平台查询菜单 | open.AlipayOpenPublicMenuQueryRequest
获取用户地理位置 | open.AlipayOpenPublicGisQueryRequest
创建菜单 | open.AlipayOpenPublicMenuCreateRequest
更新菜单 | open.AlipayOpenPublicMenuModifyRequest
带参推广二维码接口 | open.AlipayOpenPublicQrcodeCreateRequest
消息模板领取接口 | open.AlipayOpenPublicTemplateMessageGetRequest
群发消息 | open.AlipayOpenPublicMessageTotalSendRequest
重新设置绑定商家会员号 | open.AlipayOpenPublicAccountResetRequest
获取关注者列表 | open.AlipayOpenPublicFollowBatchqueryRequest
服务窗短链自主生成接口 | open.AlipayOpenPublicShortlinkCreateRequest

## 芝麻信用 API

描述 | API
---|---
获取数据反馈模板 | zhima.ZhimaDataFeedbackurlQueryRequest
批量数据反馈服务 | zhima.ZhimaDataBatchFeedbackRequest
芝麻信用评分普惠版 | zhima.ZhimaCreditScoreBriefGetRequest
行业关注名单普惠版 | zhima.ZhimaCreditWatchlistBriefGetRequest
欺诈信息验证 | zhima.ZhimaCreditAntifraudVerifyRequest
认证初始化 | zhima.ZhimaCustomerCertificationInitializeRequest
芝麻认证开始认证 | zhima.ZhimaCustomerCertificationCertifyRequest
芝麻认证查询 | zhima.ZhimaCustomerCertificationQueryRequest

## 工具类 API

描述 | API
---|---
换取应用授权令牌 | util.AlipayOpenAuthTokenAppRequest
换取授权访问令牌 | util.AlipaySystemOauthTokenRequest
查询某个应用授权AppAuthToken的授权信息 | util.AlipayOpenAuthTokenAppQueryRequest
用户登陆授权 | util.AlipayUserInfoAuthRequest

## 风险控制 API
描述 | API
---|---
“蚁盾”风险评分服务 | risk.AlipaySecurityRiskRainscoreQueryRequest
蚁盾风险评分服务上数版 | risk.SsdataDataserviceRiskRainscoreQueryRequest

## 服务市场 API
描述 | API
---|---
订购插件订单明细查询 | market.AlipayOpenServicemarketOrderQueryRequest
服务市场商户确认订购通知 | market.AlipayOpenServicemarketOrderNotifyRequest
服务商接单操作 | market.AlipayOpenServicemarketOrderAcceptRequest
服务商完成订单内单个明细实施项 | market.AlipayOpenServicemarketOrderItemCompleteRequest
服务订单明细实施项单项取消 | market.AlipayOpenServicemarketOrderItemCancelRequest
服务商拒绝接单 | market.AlipayOpenServicemarketOrderRejectRequest
门店插件下架操作 | market.AlipayOpenServicemarketCommodityShopOfflineRequest
服务商代商家确认实施完成 | market.AlipayOpenServicemarketOrderItemConfirmRequest
门店插件上架操作 | market.AlipayOpenServicemarketCommodityShopOnlineRequest

## 账务 API

描述 | API
---|---
查询对账单下载地址 | bill.AlipayDataDataserviceBillDownloadurlQueryRequest

## 生活缴费 API
描述 | API
---|---
查询账单 | ebpp.AlipayEbppBillGetRequest
创建账单 | ebpp.AlipayEbppBillAddRequest
缴费直连代扣订单支付状态查询 | ebpp.AlipayEbppPdeductBillPayStatusRequest
直连代扣协议查询接口 | ebpp.AlipayEbppPdeductSignQueryRequest
缴费直连代扣取消签约 | ebpp.AlipayEbppPdeductSignCancelRequest
公共事业缴费直连代扣扣款支付接口 | ebpp.AlipayEbppPdeductPayRequest
缴费直连代扣签约 | ebpp.AlipayEbppPdeductSignAddRequest

## 车主服务 API
描述 | API
---|---
订单同步接口 | car.AlipayEcoMycarParkingOrderSyncRequest
车辆驶出接口 | car.AlipayEcoMycarParkingExitinfoSyncRequest
修改停车场信息 | car.AlipayEcoMycarParkingParkinglotinfoUpdateRequest
车辆驶入接口 | car.AlipayEcoMycarParkingEnterinfoSyncRequest
录入停车场信息 | car.AlipayEcoMycarParkingParkinglotinfoCreateRequest
订单更新接口 | car.AlipayEcoMycarParkingOrderUpdateRequest
车牌查询接口 | car.AlipayEcoMycarParkingVehicleQueryRequest

## 数据服务 API
描述 | API
---|---
isv 回传的商户操作行为信息调用接口 | data.AlipayOfflineProviderShopactionRecordRequest
isv 回传的用户操作行为信息调用接口 | data.AlipayOfflineProviderUseractionRecordRequest
营销活动人群组规则修改接口 | data.KoubeiMarketingCampaignCrowdModifyRequest
营销活动人群组规则创建接口 | data.KoubeiMarketingCampaignCrowdCreateRequest
商户会员标签列表查询接口 | data.KoubeiMarketingCampaignTagsQueryRequest
营销活动人群组规则详情查询接口 | data.KoubeiMarketingCampaignCrowdDetailQueryRequest
营销活动人群组规则列表分页查询接口 | data.KoubeiMarketingCampaignCrowdBatchqueryRequest
营销活动人群组规则删除接口 | data.KoubeiMarketingCampaignCrowdDeleteRequest
营销活动人群组人数统计接口 | data.KoubeiMarketingCampaignCrowdCountRequest
口碑商户经营数据查询接口 | data.KoubeiMarketingDataIndicatorQueryRequest
自定义报表规则删除接口 | data.KoubeiMarketingDataCustomreportDeleteRequest
自定义报表数据查询接口 | data.KoubeiMarketingDataCustomreportQueryRequest
自定义报表规则详情查询接口 | data.KoubeiMarketingDataCustomreportDetailQueryRequest
自定义报表规则列表分页查询接口 | data.KoubeiMarketingDataCustomreportBatchqueryRequest
自定义报表规则创建及更新接口 | data.KoubeiMarketingDataCustomreportSaveRequest
口碑菜品热度查询 | data.AlipayOfflineProviderDishQueryRequest

## 卡券 API
描述 | API
---|---
卡券模板创建 | coupon.AlipayPassTemplateAddRequest
支付宝pass更新模版接口 | coupon.AlipayPassTemplateUpdateRequest
支付宝pass新建卡券实例接口 | coupon.AlipayPassInstanceAddRequest
支付宝pass更新卡券实例接口 | coupon.AlipayPassInstanceUpdateRequest
集点查询 | coupon.KoubeiMarketingToolPointsQueryRequest
更新卡集点 | coupon.KoubeiMarketingToolPointsUpdateRequest
发券授权 | coupon.KoubeiMarketingToolPrizesendAuthRequest

## 广告 API
描述 | API
---|---
修改广告接口 | ad.AlipayMarketingCdpAdvertiseModifyRequest
操作广告接口 | ad.AlipayMarketingCdpAdvertiseOperateRequest
查询广告接口 | ad.AlipayMarketingCdpAdvertiseQueryRequest
创建广告接口 | ad.AlipayMarketingCdpAdvertiseCreateRequest

## 地铁购票 API
描述 | API
---|---
地铁购票站点数据查询 | subway.AlipayCommerceCityfacilitatorStationQueryRequest
地铁购票核销码发码 | subway.AlipayCommerceCityfacilitatorVoucherGenerateRequest
地铁购票订单批量查询 | subway.AlipayCommerceCityfacilitatorVoucherBatchqueryRequest
地铁购票发码退款 | subway.AlipayCommerceCityfacilitatorVoucherRefundRequest

## 理财 API
描述 | API
---|---
余利宝账户和收益查询 | finance.MybankFinanceYulibaoAccountQueryRequest
余利宝历史交易查询 | finance.MybankFinanceYulibaoTransHistoryQueryRequest
余利宝申购 | finance.MybankFinanceYulibaoCapitalPurchaseRequest
余利宝赎回 | finance.MybankFinanceYulibaoCapitalRansomRequest
查询余利宝行情信息 | finance.MybankFinanceYulibaoPriceQueryRequest
