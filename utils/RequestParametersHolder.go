package utils

type RequestParametersHolder struct {
  ProtocalMustParams *AlipayHashMap
  ProtocalOptParams  *AlipayHashMap
  ApplicationParams  *AlipayHashMap
}

func NewRequestParametersHolder() *RequestParametersHolder {
  return &RequestParametersHolder{}
}
