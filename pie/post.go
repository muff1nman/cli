package pie
import (
  "time"
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
  STREAM_URL = "/posts"
)

func Stream(token string) (posts []*Post, err error) {
  posts = []*Post{}
  request := &PieGetRequest{
    Url: STREAM_URL,
    Token: token,
    ExtraParams: map[string]string{"type": "stream"},
  }
  err = GetPieResource(request, &posts)
  return
}
