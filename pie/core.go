package pie

import (
  "errors"
  "github.com/jmcvetta/napping"
)
const (
  URL_PREFIX = "http://localhost:3000/v1"
)

func GetPieResource(url string, token string, response interface{}, extra_params *map[string]string) (err error) {
  params := napping.Params{}
  if extra_params != nil {
    for k, v := range *extra_params {
      params[k] = v
    }
  }
  if token != "" {
    params["token"] = token
  }
  res, err := napping.Get(URL_PREFIX + url, &params, response, nil)
  if err != nil { return }
  if res.Status() != 200 {
    err = errors.New("Status code id not 200")
  }
  return
}

