package pie
import (
  "fmt"
  "time"
)

type Post struct {
  Id int `json:"id"`
  CompanyId int `json:"company_id"`
  UserId int `json:"user_id"`
  CommentsCount int `json:"comments_count"`
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

type newPostReq struct {
  Title string `json:"url"`
}

type updatePostReq struct {
}

func getUpdatePostUrl(id int) string {
  return fmt.Sprintf("/posts/%d", id)
}

func buildStreamRequest(token string) *pieGetRequest{
  return &pieGetRequest{
    Url: "/posts",
    Token: token,
    ExtraParams: map[string]string{"type": "stream"},
  }
}

// Returns the stream posts for the current user.
func Stream(token string) (posts []*Post, err error) {
  posts = []*Post{}
  err = getPieResource(buildStreamRequest(token), &posts)
  return
}

// Returns the stream posts for the current user. Returns raw response.
func RawStream(token string) (res string, err error) {
  res, err = getRawPieResource(buildStreamRequest(token))
  return
}

// Creates a new post with a topic.
func CreatePost(topic string, token string) (post *Post, err error) {
  payload := &newPostReq {
    Title: topic,
  }
  post = &Post{}
  req := &piePostRequest{
    Url: "/posts",
    Payload: payload,
    Token: token,
  }

  err = postPieResource(req, post)
  return
}

// Makes a new post published.
func PublishPost(id int, token string) (post *Post, err error) {
  payload := &updatePostReq {
  }
  post = &Post{}
  req := &piePutRequest{
    Url: getUpdatePostUrl(id),
    Payload: payload,
    Token: token,
  }

  err = putPieResource(req, post)
  return
}
