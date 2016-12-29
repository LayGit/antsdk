package util

// 用户登陆授权
// 用户进行账密、扫码登陆并授权
type AlipayUserInfoAuthRequest struct {
  Scopes  []string  `json:"scopes"`   // 接口权限值，目前只支持auth_user和auth_base两个值。 auth_base：以auth_base为scope发起的网页授权，是用来获取进入页面的用户的userId的，并且是静默授权并自动跳转到回调页的。用户感知的就是直接进入了回调页（通常是业务页面）。 auth_user：以auth_user为scope发起的网页授权，是用来获取用户的基本信息的（比如头像、昵称等）。但这种授权需要用户手动同意，用户同意后，就可在授权后获取到该用户的基本信息。
  State   string    `json:"state"`    // 商户自定义参数，用户授权后，重定向到redirect_uri时会原样回传给商户。 为防止CSRF攻击，建议开发者请求授权时传入state参数，该参数要做到既不可预测，又可以证明客户端和当前第三方网站的登录认证状态存在关联。
}

func (this *AlipayUserInfoAuthRequest) GetMethod() string {
  return "alipay.user.info.auth"
}
