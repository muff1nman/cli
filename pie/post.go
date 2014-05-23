package pie
import (
  "fmt"
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

type NewPostReq struct {
  Title string `json:"url"`
}
type UpdatePostReq struct {
}

const (
  NEW_POST_URL = "/posts"
  UPDATE_POST_URL = "/posts/%d"
  STREAM_URL = "/posts"
)

func getUpdatePostUrl(id int) string {
  return fmt.Sprintf(UPDATE_POST_URL, id)
}

func BuildStreamRequest(token string) *PieGetRequest{
  return &PieGetRequest{
    Url: STREAM_URL,
    Token: token,
    ExtraParams: map[string]string{"type": "stream"},
  }
}

func Stream(token string) (posts []*Post, err error) {
  posts = []*Post{}
  err = GetPieResource(BuildStreamRequest(token), &posts)
  return
}

func RawStream(token string) (res string, err error) {
  res, err = GetRawPieResource(BuildStreamRequest(token))
  return
}

func CreatePost(topic string, token string) (post *Post, err error) {
  payload := &NewPostReq {
    Title: topic,
  }
  post = &Post{}
  req := &PiePostRequest{
    Url: NEW_POST_URL,
    Payload: payload,
    Token: token,
  }

  err = PostPieResource(req, post)
  return
}

func PublishPost(id int, token string) (post *Post, err error) {
  payload := &UpdatePostReq {
  }
  post = &Post{}
  req := &PiePutRequest{
    Url: getUpdatePostUrl(id),
    Payload: payload,
    Token: token,
  }

  err = PutPieResource(req, post)
  return
}
