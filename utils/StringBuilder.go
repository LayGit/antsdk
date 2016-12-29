package utils

import (
  "fmt"
  "bytes"
)

type StringBuilder struct {
    buf bytes.Buffer
}

func NewStringBuilder() *StringBuilder {
  return &StringBuilder{ buf: bytes.Buffer{} }
}

func (this *StringBuilder) Append(obj interface{}) *StringBuilder {
  this.buf.WriteString(fmt.Sprintf("%v", obj))
  return this
}

func (this *StringBuilder) ToString() string {
  return this.buf.String()
}
