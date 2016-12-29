package utils

import (
  "fmt"
)

type AlipayHashMap struct {
  mMap map[string]string
  Length int
}

func NewAlipayHashMap() *AlipayHashMap {
  return &AlipayHashMap{ mMap: make(map[string]string), Length: 0 }
}

func (this *AlipayHashMap) Put(key string, value interface{}) {
  var strValue string
  if value == nil {
    strValue = ""
  } else {
    strValue = fmt.Sprintf("%v", value)
  }

  this.mMap[key] = strValue
  this.Length = len(this.mMap)
}

func (this *AlipayHashMap) Get(key string) string {
  return this.mMap[key]
}

func (this *AlipayHashMap) Remove(key string) {
  delete(this.mMap, key)
}

func (this *AlipayHashMap) GetMap() map[string]string {
  return this.mMap
}
