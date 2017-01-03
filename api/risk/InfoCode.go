package risk

type InfoCode struct {
  RiskFactorCode  string `json:"risk_factor_code"`  // 风险因素编码
  RiskFactorName  string `json:"risk_factor_name"`  // 风险因素名称
  RiskDescription string `json:"risk_description"`  // 风险描述
  RiskMagnitude   string `json:"risk_magnitude"`    // 风险度量
}
