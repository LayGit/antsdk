package utils

import (
  "encoding/json"
)

func ToJson(i interface{}) string {
  b, _ := json.Marshal(i)
  return string(b)
}
