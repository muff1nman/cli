package pie

import (
  "errors"
  "github.com/jmcvetta/napping"
)
var (
  UrlPrefix = "https://api.piethis.com/v1"
)

type request struct {
  Url string
  Token string
  ExtraParams map[string]string
  Payload interface{}
}


func (this request) getUrl() string {
  return UrlPrefix + this.Url
}

func (this request) getUrlWithToken() (url string) {
  url = this.getUrl()
  if this.Token != "" {
    url = url + "?token=" + this.Token
  }
  return
}

func (this request) getParams() (params napping.Params) {
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

func (this request) doGet(response interface{}) (err error) {
  params := this.getParams()
  res, err := napping.Get(this.getUrl(), &params, response, nil)
  if err != nil { return }
  if res.Status() != 200 {
    err = errors.New("Status code is not 200")
  }
  return
}

func (this request) doGetRaw() (body string, err error) {
  params := this.getParams()
  res, err := napping.Get(this.getUrl(), &params, nil, nil)
  if err != nil { return }
  if res.Status() != 200 {
    err = errors.New("Status code is not 200")
  }
  body = res.RawText()
  return
}

func (this request) doPost(response interface{}) (err error) {
  resp, err := napping.Post(this.getUrlWithToken(), this.Payload, response, nil)
  if err != nil { return }
  if resp.Status() != 201 {
    err = errors.New("Status code is not 201")
  }
  return
}

func (this request) doPut(response interface{}) (err error) {
  resp, err := napping.Put(this.getUrlWithToken(), this.Payload, response, nil)
  if err != nil { return }
  if resp.Status() != 200 {
    err = errors.New("Status code is not 200")
  }
  return
}
