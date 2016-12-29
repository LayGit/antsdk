package user

type AlipayUserDeliverAddress struct {
  DeliverFullname        string `json:"deliver_fullname"`          // 收货人全名
  DefaultDeliverAddress  string `json:"default_deliver_address"`   // 是否默认收货地址
  DeliverPhone           string `json:"deliver_phone"`             // 收货地址的联系人固定电话
  DeliverMobile          string `json:"deliver_mobile"`            // 收货地址的联系人移动电话
  Address                string `json:"address"`                   // 地址
  Zip                    string `json:"zip"`                       // 邮政编码
  DeliverProvince        string `json:"deliver_province"`          // 收货人所在省份
  DeliverCity            string `json:"deliver_city"`              // 收货人所在城市
  DeliverArea            string `json:"deliver_area"`              // 收货人所在区县
  AddressCode            string `json:"address_code"`              // 区域编码
}
