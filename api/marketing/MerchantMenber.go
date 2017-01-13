package marketing

type MerchantMenber struct {
  Name  string `json:"name"`  // 姓名
  Gende string `json:"gende"` // 性别（男：MALE；女：FEMALE）
  Birth string `json:"birth"` // 生日 yyyy-MM-dd
  Cell  string `json:"cell"`  // 手机号
}
