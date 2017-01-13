package marketing

type CardUserInfo struct {
  UserUniId     string `json:"user_uni_id"`       // 用户唯一标识, 根据user_id_type类型来定 （目前暂支持支付宝userId） 支付宝userId说明：支付宝用户号是以2088开头的16位纯数字组成
  UserUniIdType string `json:"user_uni_id_type"`  // ID类型：UID， 即传值UID即可
}
