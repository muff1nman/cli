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
  Pending bool `json:"pending"`
}

type updatePostReq struct {
}

func buildStreamRequest(token string) *request{
  return &request{
    Url: "/posts",
    Token: token,
    ExtraParams: map[string]string{"type": "recent_stream"},
  }
}

// Returns the stream posts for the current user.
func Stream(token string) (posts []*Post, err error) {
  posts = []*Post{}
  err = buildStreamRequest(token).doGet(&posts)
  return
}

// Returns the stream posts for the current user. Returns raw response.
func RawStream(token string) (res string, err error) {
  res, err = buildStreamRequest(token).doGetRaw()
  return
}

// Creates a new post with a topic. If pending is true, it creates a post only visible by the owner.
// Use PublishPost to make it visible to everyone.
func CreatePost(topic string, token string, pending bool) (post *Post, err error) {
  payload := &newPostReq {
    Title: topic,
    Pending: pending,
  }
  post = &Post{}
  req := &request{
    Url: "/posts",
    Payload: payload,
    Token: token,
  }

  err = req.doPost(post)
  return
}

// Makes a new post published.
func PublishPost(id int, token string) (post *Post, err error) {
  payload := &updatePostReq {
  }
  post = &Post{}
  req := &request{
    Url: fmt.Sprintf("/posts/%d", id),
    Payload: payload,
    Token: token,
  }

  err = req.doPut(post)
  return
}
