package pie
import (
  "time"
  "errors"
  "github.com/jmcvetta/napping"
)

type Post struct {
  Id int `json:"id"`
  CompanyId int `json:"company_id"`
  UserId int `json:"user_id"`
  Title string `json:"title"`
  Description string `json:"description"`
  Url string `json:"url"`
  Image string `json:"image"`
  EmbedUri string `json:"embed_uri"`
  Provider string `json:"provider"`
  Format string `json:"format"`
  Mimetype string `json:"mimetype"`
  Source string`json:"source"`
  Tags []string `json:"tags"`
  CreatedAt time.Time `json:"created_at"`
  Pending bool `json:"pending"`
  Secret bool `json:"secret"`
}

const (
  STREAM_URL = URL_PREFIX + "/posts"
)

func Stream(token string) (posts []*Post, err error) {
  posts = []*Post{}
  params := &napping.Params{"type": "stream", "token": token}
  resp, err := napping.Get(STREAM_URL, params, &posts, nil)
  if err != nil { return }
  if resp.Status() != 200 {
    err = errors.New("Error fetching stream")
    return
  }
  return
}
