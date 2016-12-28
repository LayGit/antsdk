package alipay

import (
  "net/url"
  "net/http"
  "io/ioutil"
  "strings"
  "encoding/json"
  "reflect"
  "errors"
  "regexp"
  "time"

  "github.com/LayGit/antsdk/utils"
  "github.com/LayGit/antsdk/api"
)

type AlipayClient struct {
  gatewayUrl              string
  appId                   string
  rsaPrivateKeyFilePath   string
  format                  string
  charset                 string
  alipayPublicKeyFilePath string
  signType                string
  version                 string
  notifyUrl               string
}

func NewDefaultAlipayClient(gatewayUrl, appId, rsaPrivateKeyFilePath, alipayPublicKeyFilePath, notifyUrl string) *AlipayClient {
  return &AlipayClient{
    gatewayUrl: gatewayUrl,
    appId: appId,
    rsaPrivateKeyFilePath: rsaPrivateKeyFilePath,
    format: "JSON",
    charset: "utf-8",
    alipayPublicKeyFilePath: alipayPublicKeyFilePath,
    signType: "RSA",
    version: "1.0",
    notifyUrl: notifyUrl,
  }
}

func (this *AlipayClient) getCommonRequest(method string) *api.CommonRequest {
  commonRequest := &api.CommonRequest{}
  commonRequest.AppId = this.appId
  commonRequest.Method = method
  commonRequest.Format = this.format
  commonRequest.Charset = this.charset
  commonRequest.SignType = this.signType
  commonRequest.Version = this.version
  commonRequest.NotifyUrl = this.notifyUrl
  commonRequest.Timestamp = time.Now().Format("2006-01-02 15:04:05")
  return commonRequest
}

func (this *AlipayClient) Execute(request interface{}, response interface{}) error {
  // 通过反射获取方法
  m := reflect.ValueOf(request).MethodByName("GetMethod")
  values := m.Call([]reflect.Value{})

  param := this.getCommonRequest(values[0].String())
  b, _ := json.Marshal(request)
  param.BizContent = string(b)

  var err error

  // 签名
  param.Sign, err = utils.Sign(param.ToMap(), utils.ReadPemFile(this.rsaPrivateKeyFilePath))
  if err != nil {
    return err
  }

  params := param.ToMap()

  data := &url.Values{}
  for k, v := range params {
    if v != "" {
      data.Set(k, v)
    }
  }

  reqParams := ioutil.NopCloser(strings.NewReader(data.Encode()))
  client := &http.Client{}
  req, _ := http.NewRequest("POST", this.gatewayUrl + "?charset=" + this.charset, reqParams)
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value")

  resp, err := client.Do(req)
  defer resp.Body.Close()
  if err != nil {
    return err
  }

  bResult, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }

  strResp := string(bResult)

  // 正则验签
  expResult := `(^\{\"[a-z|_]+\":)|(,\"sign\":\"[a-zA-Z0-9|\+|\/|\=]+\"\}$)`
  exptSign := `\"sign\":\"([a-zA-Z0-9|\+|\/|\=]+)\"`
  regResult := regexp.MustCompile(expResult)
  result := regResult.ReplaceAllString(strResp, "")
  regSign := regexp.MustCompile(exptSign)
  signMatchRes := regSign.FindStringSubmatch(strResp)
  if len(signMatchRes) < 2 {
    return errors.New("loss of sign")
  }
  sign := signMatchRes[1]

  // 验证签名
  isOk, err := utils.SyncVerifySign(sign, []byte(result), utils.ReadPemFile(this.alipayPublicKeyFilePath))
  if err != nil {
    return err
  }

  if !isOk {
    return errors.New("Response sign verify error")
  }

  return json.Unmarshal(bResult, &response)
}
