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

func buildCommentsRequest(post_id int, token string) *pieGetRequest {
  return &pieGetRequest{
    Url: getCommentsUrl(post_id),
    Token: token,
  }
}


func getNewCommentUrl(post_id int) string {
  return fmt.Sprintf("/posts/%d/comments", post_id)
}

func getCommentsUrl(post_id int) string {
  return fmt.Sprintf("/posts/%d/comments", post_id)
}

// Creates a new comment in a given post.
func CreateComment(post_id int, text string, token string) (comment *Comment, err error) {
  payload := &newCommentReq {
    Text: text,
  }
  comment = &Comment{}
  req := &piePostRequest{
    Url: getNewCommentUrl(post_id),
    Payload: payload,
    Token: token,
  }

  err = postPieResource(req, comment)
  return
}

// Returns all comments for a given post.
func GetComments(post_id int, token string) (comments []*Comment, err error) {
  comments = []*Comment{}
  err = getPieResource(buildCommentsRequest(post_id, token), &comments)
  return
}

// Returns all comments for a given post in raw format.
func GetRawComments(post_id int, token string) (res string, err error) {
  res, err = getRawPieResource(buildCommentsRequest(post_id, token))
  return
}
