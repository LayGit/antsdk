package open

type StdPublicBindAccount struct {
  FromUserId    string `json:"from_user_id"`      // 绑定的商户会员对应的支付宝用户号，以2088 开头的16位数字。
  AppId         string `json:"app_id"`            // 公众账号ID
  AgreementId   string `json:"agreement_id"`      // 协议号是商户会员在支付宝公众账号中的唯一标识。
  BindAccountNo string `json:"bind_account_no"`   // 绑定的商户会员号
  RealName      string `json:"real_name"`         // 绑定的商户会员的真实姓名，最长10个汉字。
  DisplayName   string `json:"display_name"`      // 公众账号期望支付宝用户在公众账号首页看到的关于该用户的显示信息，最长10个汉字。
}
