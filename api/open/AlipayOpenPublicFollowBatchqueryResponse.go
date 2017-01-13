package open

import (
  "github.com/LayGit/antsdk/api"
)

type AlipayOpenPublicFollowBatchqueryResponse struct {
  api.AlipayResponse
  Count       string    `json:"count"`         // 本次调用获取的userId个数，最大值为10000
  NextUserId  string    `json:"next_user_id"`  // 查询分组的userid
  UserIdList  []string  `json:"user_id_list"`  // 用户的userId列表
}
