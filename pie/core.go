package pie

import (
  "errors"
  "github.com/jmcvetta/napping"
)
const (
  URL_PREFIX = "http://localhost:3000/v1"
)

type PieGetRequest struct {
  Url string
  Token string
  ExtraParams map[string]string
}

func (this PieGetRequest) GetParams() (params napping.Params) {
  params = napping.Params{}
  if this.ExtraParams != nil {
    for k, v := range this.ExtraParams {
      params[k] = v
    }
  }
  if this.Token != "" {
    params["token"] = this.Token
  }
  return
}

func GetPieResource(request *PieGetRequest, response interface{}) (err error) {
  params := request.GetParams()
  res, err := napping.Get(URL_PREFIX + request.Url, &params, response, nil)
  if err != nil { return }
  if res.Status() != 200 {
    err = errors.New("Status code id not 200")
  }
  return
}

