package open

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 服务窗基础信息修改接口
// 适用于口碑商户签约当面付2.0默认创建的服务窗（名称为“服务窗+时间戳”，头像为空），系统商需要通过该接口帮助口碑商户修改服务窗名称、头像、欢迎语基础信息，修改名称、头像需要上传相关资质材料，触发风控审核
type AlipayOpenPublicInfoModifyRequest struct {
  api.IAlipayRequest
  TerminalType      string                                       `json:"terminal_type"`
  TerminalInfo      string                                       `json:"terminal_info"`
  ProdCode          string                                       `json:"prod_code"`
  NotifyUrl         string                                       `json:"notify_url"`
  ReturnUrl         string                                       `json:"return_url"`
  BizContent        AlipayOpenPublicInfoModifyRequestBizContent  `json:"biz_content"`
}

type AlipayOpenPublicInfoModifyRequestBizContent struct {
  AppName         string `json:"app_name"`          // 服务窗名称，2-20个字之间；不得含有违反法律法规和公序良俗的相关信息；不得侵害他人名誉权、知识产权、商业秘密等合法权利；不得以太过广泛的、或产品、行业词组来命名，如：女装、皮革批发；不得以实名认证的媒体资质账号创建服务窗，或媒体相关名称命名服务窗，如：XX电视台、XX杂志等
  LogoUrl         string `json:"logo_url"`          // 服务窗头像地址，建议尺寸 320 x 320px，支持.jpg .jpeg .png 格式，小于3M
  PublicGreeting  string `json:"public_greeting"`   // 服务窗欢迎语，200字以内，首次使用服务窗必须
  LicenseUrl      string `json:"license_url"`       // 营业执照地址，建议尺寸 320 x 320px，支持.jpg .jpeg .png 格式，小于3M
  ShopPics        []string `json:"shop_pics"`       // 门店照片Url
  AuthPic         string `json:"auth_pic"`          // 授权运营书，企业商户若为被经营方授权，需上传加盖公章的扫描件，请使用照片上传接口上传图片获得image_url
}

func (this *AlipayOpenPublicInfoModifyRequest) GetApiMethodName() string {
  return "alipay.open.public.info.modify"
}

func (this *AlipayOpenPublicInfoModifyRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOpenPublicInfoModifyRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOpenPublicInfoModifyRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOpenPublicInfoModifyRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOpenPublicInfoModifyRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOpenPublicInfoModifyRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOpenPublicInfoModifyRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOpenPublicInfoModifyRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_content", utils.ToJson(this.BizContent))
  return txtParams
}
