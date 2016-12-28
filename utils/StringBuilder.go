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

func (this *StringBuilder) Append(obj interface{}) {
  this.buf.WriteString(fmt.Sprintf("%v", obj))
}

func (this *StringBuilder) ToString() string {
  return this.buf.String()
}
