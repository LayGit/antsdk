package marketing

type TemplateBenefitInfoDTO struct {
  Title       string `json:"title"`         // 权益描述
  BenefitDesc string `json:"benefit_desc"`  // 权益描述信息
  StartDate   string `json:"start_date"`    // 开始时间
  EndDate     string `json:"end_date"`      // 权益结束时间
}
