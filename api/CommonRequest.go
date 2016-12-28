package api

type CommonRequest struct {
  AppId         string  `json:"app_id"`         // 支付宝分配给开发者的应用ID
  Method        string  `json:"method"`         // 接口名称
  Format        string  `json:"format"`         // 仅支持JSON
  Charset       string  `json:"charset"`        // 请求使用的编码格式，如utf-8,gbk,gb2312等
  SignType      string  `json:"sign_type"`      // 商户生成签名字符串所使用的签名算法类型，目前支持RSA
  Sign          string  `json:"sign"`           // 商户请求参数的签名串，详见签名(https://doc.open.alipay.com/doc2/detail.htm?treeId=200&articleId=105351&docType=1)
  Timestamp     string  `json:"timestamp"`      // 送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
  Version       string  `json:"version"`        // 调用的接口版本，固定为：1.0
  NotifyUrl     string  `json:"notify_url"`     // 支付宝服务器主动通知商户服务器里指定的页面http/https路径。
  AppAuthToken  string  `json:"app_auth_token"` // 详见应用授权概述(https://doc.open.alipay.com/doc2/detail.htm?treeId=216&articleId=105193&docType=1)
  BizContent    string  `json:"biz_content"`    // 请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
}

func (this *CommonRequest) ToMap() map[string]string {
  var m map[string]string
  m = make(map[string]string)
  m["app_id"] = this.AppId
  m["method"] = this.Method
  m["format"] = this.Format
  m["charset"] = this.Charset
  m["sign_type"] = this.SignType
  m["sign"] = this.Sign
  m["timestamp"] = this.Timestamp
  m["version"] = this.Version
  m["app_auth_token"] = this.AppAuthToken
  m["biz_content"] = this.BizContent
  return m
}
