package finance

import (
  "github.com/LayGit/antsdk/api"
)

type MybankFinanceYulibaoCapitalPurchaseResponse struct {
  api.AlipayResponse
  TransResult string `json:"trans_result"`  // 交易结果，目前会有3种状态值，1：success，表示交易成功、2：fail，表示交易失败:、3：inprocess，表示交易处理中。其中交易处理中的状态可以使用回查交易历史的方式查看其处理结果。
  InnerBizNo  string `json:"inner_biz_no"`  // 余利宝内部的交易流水号。
  Remark      string `json:"remark"`        // 交易结果的备注信息。
}
