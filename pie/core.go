package pie

import (
  "errors"
  "github.com/jmcvetta/napping"
)
var (
  UrlPrefix = "https://api.piethis.com/v1"
)

type pieGetRequest struct {
  Url string
  Token string
  ExtraParams map[string]string
}

type piePostRequest struct {
  Url string
  Token string
  Payload interface{}
}

type piePutRequest struct {
  Url string
  Token string
  Payload interface{}
}

func (this pieGetRequest) getUrl() string {
  return UrlPrefix + this.Url
}
func (this piePostRequest) getUrl() string {
  if this.Token == "" {
    return UrlPrefix + this.Url
  } else {
    return UrlPrefix + this.Url + "?token=" + this.Token
  }
}
func (this piePutRequest) getUrl() string {
  if this.Token == "" {
    return UrlPrefix + this.Url
  } else {
    return UrlPrefix + this.Url + "?token=" + this.Token
  }
}

func (this pieGetRequest) getParams() (params napping.Params) {
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

func getPieResource(request *pieGetRequest, response interface{}) (err error) {
  params := request.getParams()
  res, err := napping.Get(request.getUrl(), &params, response, nil)
  if err != nil { return }
  if res.Status() != 200 {
    err = errors.New("Status code is not 200")
  }
  return
}

func getRawPieResource(request *pieGetRequest) (body string, err error) {
  params := request.getParams()
  res, err := napping.Get(request.getUrl(), &params, nil, nil)
  if err != nil { return }
  if res.Status() != 200 {
    err = errors.New("Status code is not 200")
  }
  body = res.RawText()
  return
}

func postPieResource(request *piePostRequest, response interface{}) (err error) {
  resp, err := napping.Post(request.getUrl(), request.Payload, response, nil)
  if err != nil { return }
  if resp.Status() != 201 {
    err = errors.New("Status code is not 201")
  }
  return
}

func putPieResource(request *piePutRequest, response interface{}) (err error) {
  resp, err := napping.Put(request.getUrl(), request.Payload, response, nil)
  if err != nil { return }
  if resp.Status() != 200 {
    err = errors.New("Status code is not 200")
  }
  return
}
