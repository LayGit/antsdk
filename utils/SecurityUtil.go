package utils

import (
  "sort"
  "encoding/pem"
  "encoding/base64"
  "encoding/hex"
  "crypto"
  "crypto/x509"
  "crypto/rsa"
  "crypto/sha1"
  "errors"
  "net/url"
  "io"
  "io/ioutil"
  "os"
  "fmt"
)

func GetSignStr(m map[string]string) string {
  // 对 key 进行升序排序
  sortedKeys := make([]string, 0)
  for k, _ := range m {
    sortedKeys = append(sortedKeys, k)
  }
  sort.Strings(sortedKeys)

  // 对 key=value 的键值对用 & 连接起来，略过空值
  sbSignStr := NewStringBuilder()
  for i, k := range sortedKeys {
    if m[k] != "" {
      sbSignStr.Append(k)
      sbSignStr.Append("=")
      sbSignStr.Append(m[k])
      if i != (len(sortedKeys) - 1) {
        sbSignStr.Append("&")
      }
    }
  }
  return sbSignStr.ToString()
}

func Sign(mReq map[string]string, privateKey []byte) (string, error) {

  // 获取待签名参数
  signStr := GetSignStr(mReq)

  block, _ := pem.Decode(privateKey)
  if block == nil {
    return "", errors.New("Sign private key decode error")
  }

  prk8, err := x509.ParsePKCS8PrivateKey(block.Bytes)
  if err != nil {
    return "", err
  }

  pKey := prk8.(*rsa.PrivateKey)

  return RSASign(signStr, pKey)
}

func RSASign(origData string, privateKey *rsa.PrivateKey) (string, error) {
  h := sha1.New()
  h.Write([]byte(origData))
  digest := h.Sum(nil)

  s, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA1, []byte(digest))
  if err != nil {
      return "", err
  }
  data := base64.StdEncoding.EncodeToString(s)
  return string(data), nil
}

// 同步返回验签
func SyncVerifySign(sign string, body, alipayPublicKey []byte) (bool, error) {
  return RSAVerify(body, []byte(sign), alipayPublicKey)
}

// 异步返回验签
func AsyncVerifySign(body, alipayPublicKey []byte) (bool, error) {
  data, err := url.ParseQuery(string(body))
  if err != nil {
    return false, err
  }

  var m map[string]string
  m = make(map[string]string, 0)

  for k, v := range data {
    if k == "sign" || k == "sign_type" { //不要'sign'和'sign_type'
        continue
    }
    m[k] = v[0]
  }

  sign := data["sign"][0]

  //获取要进行计算哈希的sign string
  signStr := GetSignStr(m)

  return RSAVerify([]byte(signStr), []byte(sign), alipayPublicKey)
}

func RSAVerify(src, sign, alipayPublicKey []byte) (bool, error) {
  // 加载RSA的公钥
  block, _ := pem.Decode(alipayPublicKey)
  pub, err := x509.ParsePKIXPublicKey(block.Bytes)
  if err != nil {
      return false, err
  }
  rsaPub, _ := pub.(*rsa.PublicKey)

  // 计算代签名字串的SHA1哈希
  t := sha1.New()
  io.WriteString(t, string(src))
  digest := t.Sum(nil)

  // base64 decode,必须步骤，支付宝对返回的签名做过base64 encode必须要反过来decode才能通过验证
  data, _ := base64.StdEncoding.DecodeString(string(sign))

  hex.EncodeToString(data)

  // 调用rsa包的VerifyPKCS1v15验证签名有效性
  err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA1, digest, data)
  if err != nil {
    fmt.Println(string(src))
    fmt.Println("error")
    return false, err
  }

  return true, nil
}

func ReadPemFile(path string) []byte {
  fi,err := os.Open(path)
  if err != nil{panic(err)}
  defer fi.Close()
  fd, err := ioutil.ReadAll(fi)
  return fd
}
