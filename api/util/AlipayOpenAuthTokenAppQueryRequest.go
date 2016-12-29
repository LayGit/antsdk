package util

// 查询某个应用授权AppAuthToken的授权信息
// 当商户把服务窗、店铺等接口的权限授权给ISV之后，支付宝会给ISV颁发一个app_auth_token。如若授权成功之后，ISV想知道用户的授权信息，如授权者、授权接口列表等信息，可以调用本接口查询某个app_auth_token对应的授权信息
type AlipayOpenAuthTokenAppQueryRequest struct {
  AppAuthToken  string  `json:"app_auth_token"` // 应用授权令牌
}

func (this *AlipayOpenAuthTokenAppQueryRequest) GetMethod() string {
  return "alipay.open.auth.token.app.query"
}
