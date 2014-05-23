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

type PiePostRequest struct {
  Url string
  Token string
  Payload interface{}
}

type PiePutRequest struct {
  Url string
  Token string
  Payload interface{}
}

func (this PieGetRequest) GetUrl() string {
  return URL_PREFIX + this.Url
}
func (this PiePostRequest) GetUrl() string {
  if this.Token == "" {
    return URL_PREFIX + this.Url
  } else {
    return URL_PREFIX + this.Url + "?token=" + this.Token
  }
}
func (this PiePutRequest) GetUrl() string {
  if this.Token == "" {
    return URL_PREFIX + this.Url
  } else {
    return URL_PREFIX + this.Url + "?token=" + this.Token
  }
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
  res, err := napping.Get(request.GetUrl(), &params, response, nil)
  if err != nil { return }
  if res.Status() != 200 {
    err = errors.New("Status code is not 200")
  }
  return
}

func GetRawPieResource(request *PieGetRequest) (body string, err error) {
  params := request.GetParams()
  res, err := napping.Get(request.GetUrl(), &params, nil, nil)
  if err != nil { return }
  if res.Status() != 200 {
    err = errors.New("Status code is not 200")
  }
  body = res.RawText()
  return
}

func PostPieResource(request *PiePostRequest, response interface{}) (err error) {
  resp, err := napping.Post(request.GetUrl(), request.Payload, response, nil)
  if err != nil { return }
  if resp.Status() != 201 {
    err = errors.New("Status code is not 201")
  }
  return
}

func PutPieResource(request *PiePutRequest, response interface{}) (err error) {
  resp, err := napping.Put(request.GetUrl(), request.Payload, response, nil)
  if err != nil { return }
  if resp.Status() != 200 {
    err = errors.New("Status code is not 200")
  }
  return
}
