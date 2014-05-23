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

type NewCommentReq struct {
  Text string `json:"text"`
}

const (
  NEW_COMMENT_URL = "/posts/%d/comments"
)

func getNewCommentUrl(post_id int) string {
  return fmt.Sprintf(NEW_COMMENT_URL, post_id)
}

func CreateComment(post_id int, text string, token string) (comment *Comment, err error) {
  payload := &NewCommentReq {
    Text: text,
  }
  comment = &Comment{}
  req := &PiePostRequest{
    Url: getNewCommentUrl(post_id),
    Payload: payload,
    Token: token,
  }

  err = PostPieResource(req, comment)
  return
}
