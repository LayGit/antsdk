package open

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOpenPublicTemplateMessageGetResponse struct {
  api.AlipayResponse
  MsgTemplateId string `json:"msg_template_id"` // 消息模板id--商户领取母版后生成的唯一模板id
  Template      string `json:"template"`        // 模板内容
}
