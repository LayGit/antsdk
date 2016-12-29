package utils

import (
  "net/url"
)


func BuildQuery(params map[string]string) string {
  if params == nil {
    return ""
  }

  query := NewStringBuilder()
  hasParam := false

  for k, v := range params {
    if v != "" {
      if hasParam {
        query.Append("&")
      } else {
        hasParam = true
      }

      query.Append(k).Append("=").Append(url.QueryEscape(v))
    }
  }
  return query.ToString()
}
