package pie
import (
  "fmt"
  "time"
)

type Comment struct {
  Id int `json:"id"`
  PostId int `json:"post_id"`
  UserId int `json:"user_id"`
  Text string `json:"text"`
  CreatedAt time.Time `json:"created_at"`
}

type newCommentReq struct {
  Text string `json:"text"`
}

func buildCommentsRequest(post_id int, token string) *request {
  return &request{
    Url: fmt.Sprintf("/posts/%d/comments", post_id),
    Token: token,
  }
}

// Creates a new comment in a given post.
func CreateComment(post_id int, text string, token string) (comment *Comment, err error) {
  payload := &newCommentReq {
    Text: text,
  }
  comment = &Comment{}
  req := &request{
    Url: fmt.Sprintf("/posts/%d/comments", post_id),
    Payload: payload,
    Token: token,
  }

  err = req.doPost(comment)
  return
}

// Returns all comments for a given post.
func GetComments(post_id int, token string) (comments []*Comment, err error) {
  comments = []*Comment{}
  err = buildCommentsRequest(post_id, token).doGet(&comments)
  return
}

// Returns all comments for a given post in raw format.
func GetRawComments(post_id int, token string) (res string, err error) {
  res, err = buildCommentsRequest(post_id, token).doGetRaw()
  return
}
