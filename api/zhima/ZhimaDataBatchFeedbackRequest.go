package zhima

import (
  "github.com/LayGit/antsdk/api"
  "github.com/LayGit/antsdk/utils"
)

// 批量数据反馈服务
// 通过该服务接口，商户可以反馈数据到芝麻信用
type ZhimaDataBatchFeedbackRequest struct {
  api.IAlipayUploadRequest
  TerminalType      string            `json:"terminal_type"`
  TerminalInfo      string            `json:"terminal_info"`
  ProdCode          string            `json:"prod_code"`
  NotifyUrl         string            `json:"notify_url"`
  ReturnUrl         string            `json:"return_url"`
  FileCharset       string            `json:"file_charset"`           // 是反馈文件的数据编码，如果文件格式是UTF-8，则填写UTF-8，如果文件格式是GBK，则填写GBK
  Columns           string            `json:"columns"`                // 单条数据的数据列，多个列以逗号隔开
  File              *utils.FileItem   `json:"file"`                   // 反馈的json格式文件，其中{"records": 是每个文件的固定开头，文件中的字段名请使用数据反馈模板（该模板是通过“获取数据反馈模板”接口获得）中字段编码列的英文字段来组装
  FileType          string            `json:"file_type"`              // 反馈的数据类型
  Records           string            `json:"records"`                // 文件数据记录条数
  BizExtParams      string            `json:"biz_ext_params"`         // 扩展参数
  PrimaryKeyColumns string            `json:"primary_key_columns"`    // 主键列使用反馈字段进行组合，也可以使用反馈的某个单字段（确保主键稳定，而且可以很好的区分不同的数据）。例如order_no,pay_month或者order_no,bill_month组合，对于一个order_no只会有一条数据的情况，直接使用order_no作为主键列
  FileDescription   string            `json:"file_description"`       // 文件描述信息
}

func (this *ZhimaDataBatchFeedbackRequest) GetApiMethodName() string {
  return "zhima.data.batch.feedback"
}

func (this *ZhimaDataBatchFeedbackRequest) GetApiVersion() string {
  return "1.0"
}

func (this *ZhimaDataBatchFeedbackRequest) GetTerminalType() string {
  return this.TerminalType
}

func (this *ZhimaDataBatchFeedbackRequest) GetTerminalInfo() string {
  return this.TerminalInfo
}

func (this *ZhimaDataBatchFeedbackRequest) GetNotifyUrl() string {
  return this.NotifyUrl
}

func (this *ZhimaDataBatchFeedbackRequest) GetReturnUrl() string {
  return this.ReturnUrl
}

func (this *ZhimaDataBatchFeedbackRequest) GetProdCode() string {
  return this.ProdCode
}

func (this *ZhimaDataBatchFeedbackRequest) IsNeedEncrypt() bool {
  return false
}

func (this *ZhimaDataBatchFeedbackRequest) GetTextParams() *utils.AlipayHashMap {
  txtParams := utils.NewAlipayHashMap()
  txtParams.Put("biz_ext_params", this.BizExtParams)
  txtParams.Put("columns", this.Columns)
  txtParams.Put("file_charset", this.FileCharset)
  txtParams.Put("file_description", this.FileDescription)
  txtParams.Put("file_type", this.FileType)
  txtParams.Put("primary_key_columns", this.PrimaryKeyColumns)
  txtParams.Put("records", this.Records)
  return txtParams
}

func (this *ZhimaDataBatchFeedbackRequest) GetFileParams() map[string]*utils.FileItem {
  params := make(map[string]*utils.FileItem)
  params["file"] = this.File
  return params
}
