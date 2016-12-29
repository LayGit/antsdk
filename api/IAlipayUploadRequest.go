package api

type IAlipayUploadRequest interface {
  IAlipayRequest
  GetFileParams()
}
