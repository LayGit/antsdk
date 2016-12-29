package alipay

import (
  "net/url"
  "net/http"
  "io/ioutil"
  "mime/multipart"
  "strings"
  "encoding/json"
  "reflect"
  "errors"
  "regexp"
  "time"
  "bytes"

  "github.com/LayGit/antsdk/utils"
  "github.com/LayGit/antsdk/api"
)

type AlipayClient struct {
  serverUrl               string
  appId                   string
  privateKey              string
  prodCode                string
  format                  string
  signType                string
  encryptType             string
  encryptKey              string
  alipayPublicKey         string
  charset                 string
}

func NewDefaultAlipayClient(serverUrl, appId, privateKey, alipayPublicKey string) *AlipayClient {
  return &AlipayClient{
    serverUrl: serverUrl,
    appId: appId,
    privateKey: privateKey,
    alipayPublicKey: alipayPublicKey,
    format: CONST_FORMAT_JSON,
    signType: CONST_SIGN_TYPE_RSA,
    encryptType: CONST_ENCRYPT_TYPE_AES,
  }
}

func (this *AlipayClient) Execute(request, response interface{}) error {
  return this.ExecuteWithAccessToken(request, response, "")
}

func (this *AlipayClient) ExecuteWithAccessToken(request, response interface{}, accessToken string) error {
  return this.ExecuteWithAppAuthToken(request, response, accessToken, "")
}

func (this *AlipayClient) ExecuteWithAppAuthToken(request, response interface{}, accessToken, appAuthToken string) error {
  bResult, err := this.doPost(request, accessToken, appAuthToken)
  if err != nil {
    return err
  }

  // 验签
  strResp := string(bResult)

  // 正则验签
  expResult := `(^\{\"[a-z|_]+\":)|(,\"sign\":\"[a-zA-Z0-9|\+|\/|\=]+\"\}$)`
  exptSign := `\"sign\":\"([a-zA-Z0-9|\+|\/|\=]+)\"`
  regResult := regexp.MustCompile(expResult)
  result := regResult.ReplaceAllString(strResp, "")
  regSign := regexp.MustCompile(exptSign)
  signMatchRes := regSign.FindStringSubmatch(strResp)
  if len(signMatchRes) < 2 {
    return errors.New("验签失败:签名丢失")
  }
  sign := signMatchRes[1]

  // 验证签名
  isOk, err := utils.SyncVerifySign(sign, []byte(result), []byte(this.alipayPublicKey))
  if err != nil {
    return err
  }

  if !isOk {
    return errors.New("验签失败:签名错误")
  }

  return json.Unmarshal([]byte(result), &response)
}

func (this *AlipayClient) doPost(request interface{}, accessToken, appAuthToken string) ([]byte, error) {
  requestHolder, err := this.getRequestHolderWithSign(reflect.ValueOf(request).Interface().(api.IAlipayRequest), accessToken, appAuthToken)
  if err != nil {
    return nil, err
  }

  reqUrl := this.getRequestUrl(requestHolder)

  if fileReq, ok := request.(api.IAlipayUploadRequest); ok {
    return this.postFileRequest(reqUrl, requestHolder.ApplicationParams.GetMap(), fileReq.GetFileParams())
  }
  return this.postRequest(reqUrl, requestHolder.ApplicationParams.GetMap())
}

func (this *AlipayClient) postRequest(reqUrl string, params map[string]string) ([]byte, error) {
  data := &url.Values{}

  for k, v := range params {
    if v != "" {
      data.Set(k, v)
    }
  }

  reqParams := ioutil.NopCloser(strings.NewReader(data.Encode()))
  var client http.Client
  req, _ := http.NewRequest("POST", reqUrl, reqParams)
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value")

  resp, err := client.Do(req)
  defer resp.Body.Close()
  if err != nil {
    return nil, err
  }

  return ioutil.ReadAll(resp.Body)
}

func (this *AlipayClient) postFileRequest(reqUrl string, params map[string]string, fileParams map[string]*utils.FileItem) ([]byte, error) {
  if fileParams == nil || len(fileParams) == 0 {
    return this.postRequest(reqUrl, params)
  }

  b:= &bytes.Buffer{}
  w := multipart.NewWriter(b)

  if params != nil && len(params) > 0 {
    for k, v := range params {
      w.WriteField(k, v)
    }
  }

  for k, v := range fileParams {
      fw, err := w.CreateFormFile(k, v.FileName)
      if err != nil {
        return nil, err
      }
      fw.Write(v.Content)
  }
  w.Close()

  req, err := http.NewRequest("POST", reqUrl, b)
  if err != nil {
    return nil, err
  }
  req.Header.Set("Content-Type", w.FormDataContentType())
  var client http.Client
  resp, err := client.Do(req)
  defer resp.Body.Close()
  if err != nil {
    return nil, err
  }

  return ioutil.ReadAll(resp.Body)
}

func (this *AlipayClient) getRequestUrl(requestHolder *utils.RequestParametersHolder) string {
  sbUrl := utils.NewStringBuilder()
  sbUrl.Append(this.serverUrl)
  sysMustQuery := utils.BuildQuery(requestHolder.ProtocalMustParams.GetMap())
  sysOptQuery := utils.BuildQuery(requestHolder.ProtocalOptParams.GetMap())

  sbUrl.Append("?")
  sbUrl.Append(sysMustQuery)
  if sysOptQuery != "" {
    sbUrl.Append("&")
    sbUrl.Append(sysOptQuery)
  }

  return sbUrl.ToString()
}

func (this *AlipayClient) getRequestHolderWithSign(request api.IAlipayRequest, accessToken, appAuthToken string) (*utils.RequestParametersHolder, error) {
  requestHolder := utils.NewRequestParametersHolder()
  appParams := request.GetTextParams()

  // 只有新接口和设置密钥才能支持加密
  if request.IsNeedEncrypt() {
    if appParams.Get(CONST_BIZ_CONTENT_KEY) == "" {
      return nil, errors.New("当前API不支持加密请求")
    }
    // 需要加密必须设置密钥和加密算法
    if this.encryptKey == "" || this.encryptType == "" {
      return nil, errors.New("API请求要求加密，则必须设置密钥和密钥类型：encryptKey=" + this.encryptKey + ",encryptType=" + this.encryptType)
    }

    encryptContent, err := utils.EncryptContent(appParams.Get(CONST_BIZ_CONTENT_KEY), this.encryptType, this.encryptKey)
    if err != nil {
      return nil, err
    }

    appParams.Put(CONST_BIZ_CONTENT_KEY, encryptContent)
  }

  if appAuthToken != "" {
    appParams.Put(CONST_APP_AUTH_TOKEN, appAuthToken)
  }

  requestHolder.ApplicationParams = appParams

  if this.charset == "" {
    this.charset = CONST_CHARSET_UTF8
  }

  protocalMustParams := utils.NewAlipayHashMap()
  protocalMustParams.Put(CONST_METHOD, request.GetApiMethodName())
  protocalMustParams.Put(CONST_VERSION, request.GetApiVersion())
  protocalMustParams.Put(CONST_APP_ID, this.appId)
  protocalMustParams.Put(CONST_SIGN_TYPE, this.signType)
  protocalMustParams.Put(CONST_TERMINAL_TYPE, request.GetTerminalType())
  protocalMustParams.Put(CONST_TERMINAL_INFO, request.GetTerminalInfo())
  protocalMustParams.Put(CONST_NOTIFY_URL, request.GetNotifyUrl())
  protocalMustParams.Put(CONST_RETURN_URL, request.GetReturnUrl())
  protocalMustParams.Put(CONST_CHARSET, this.charset)

  if request.IsNeedEncrypt() {
    protocalMustParams.Put(CONST_ENCRYPT_TYPE, this.encryptType)
  }
  protocalMustParams.Put(CONST_TIMESTAMP, time.Now().Format("2006-01-02 15:04:05"))
  requestHolder.ProtocalMustParams = protocalMustParams

  protocalOptParams := utils.NewAlipayHashMap()
  protocalOptParams.Put(CONST_FORMAT, this.format)
  protocalOptParams.Put(CONST_ACCESS_TOKEN, accessToken)
  protocalOptParams.Put(CONST_ALIPAY_SDK, CONST_SDK_VERSION)
  protocalOptParams.Put(CONST_PROD_CODE, request.GetProdCode())
  requestHolder.ProtocalOptParams = protocalOptParams

  if this.signType != "" {
    signMap := utils.GetSignMap(requestHolder)
    sign, err := utils.Sign(signMap, []byte(this.privateKey))
    if err != nil {
      return nil, err
    }
    protocalMustParams.Put(CONST_SIGN, sign)
  } else {
    protocalMustParams.Put(CONST_SIGN, "")
  }

  return requestHolder, nil
}
