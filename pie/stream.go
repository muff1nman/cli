package pie
import (
  "errors"
  "github.com/jmcvetta/napping"
)

type PostRes struct {
  Title string `json:"title"`
}

const (
  STREAM_URL = URL_PREFIX + "/posts"
)

func Stream(token string) (posts []*PostRes, err error) {
  posts = []*PostRes{}
  params := &napping.Params{"type": "stream", "token": token}
  resp, err := napping.Get(STREAM_URL, params, &posts, nil)
  if err != nil { return }
  if resp.Status() != 200 {
    err = errors.New("Error fetching stream")
    return
  }
  return
}
