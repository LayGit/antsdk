package shop

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 上传门店照片和视频接口
// 系统商需要先将商户需要使用的图片和视频，上传支付宝服务器,生成对应的图片ID后才能够在口碑平台上配置相应图片
type AlipayOfflineMaterialImageUploadRequest struct {
  api.IAlipayUploadRequest
  TerminalType  string                              `json:"terminal_type"`
  TerminalInfo  string                              `json:"terminal_info"`
  ProdCode      string                              `json:"prod_code"`
  NotifyUrl     string                              `json:"notify_url"`
  ReturnUrl     string                              `json:"return_url"`
  ImageType     string                              `json:"image_type"`      // 图片/视频格式
  ImageName     string                              `json:"image_name"`      // 图片/视频名称
  ImageContent  *utils.FileItem                     `json:"image_content"`   // 图片/视频二进制内容，图片/视频大小不能超过5M
  ImagePid      string                              `json:"image_pid"`       // 用于显示指定图片/视频所属的partnerId（支付宝内部使用，外部商户无需填写此字段）
}

func (this *AlipayOfflineMaterialImageUploadRequest) GetApiMethodName() string {
  return "alipay.offline.material.image.upload"
}

func (this *AlipayOfflineMaterialImageUploadRequest) GetApiVersion() string {
  return "1.0"
}

func (this *AlipayOfflineMaterialImageUploadRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *AlipayOfflineMaterialImageUploadRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *AlipayOfflineMaterialImageUploadRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *AlipayOfflineMaterialImageUploadRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *AlipayOfflineMaterialImageUploadRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *AlipayOfflineMaterialImageUploadRequest) IsNeedEncrypt() bool {
  return false
}

func (this *AlipayOfflineMaterialImageUploadRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("image_name", this.ImageName)
  txtParams.Put("image_pid", this.ImagePid)
  txtParams.Put("image_type", this.ImageType)
  return txtParams
}

func (this *AlipayOfflineMaterialImageUploadRequest) GetFileParams() map[string]*utils.FileItem {
  params := make(map[string]*utils.FileItem)
  params["image_content"] = this.ImageContent
  return params
}
