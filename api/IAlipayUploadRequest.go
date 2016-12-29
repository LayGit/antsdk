package api

import (
  "github.com/LayGit/antsdk/utils"
)

type IAlipayUploadRequest interface {
  IAlipayRequest
  GetFileParams() map[string]*utils.FileItem
}
