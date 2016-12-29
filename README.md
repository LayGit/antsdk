# antsdk
蚂蚁金服(支付宝)开放平台 go-sdk

已支持
 - 统一收单交易撤销接口     trade.AlipayTradeCancelRequest
 - 统一收单交易关闭接口     trade.AlipayTradeCloseRequest
 - 统一收单交易创建         trade.AlipayTradeCreateRequest
 - 统一收单交易退款查询     trade.AlipayTradeFastpayRefundQueryRequest
 - 统一收单交易结算接口     trade.AlipayTradeOrderSettleRequest
 - 统一收单交易支付         trade.AlipayTradePayRequest
 - 统一收单线下交易预创建   trade.AlipayTradePreCreateRequest
 - 统一收单线下交易查询     trade.AlipayTradeQueryRequest
 - 统一收单交易退款         trade.AlipayTradeRefundRequest

使用示例：
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
        // todo:
        fmt.Println(response.Code)
    }
}
```
