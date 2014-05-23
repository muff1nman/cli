package pie
import (
  "time"
)

type Tag struct {
  Name string `json:"name"`
  NumPosts int `json:"num_posts"`
  LastActivity time.Time `json:"last_activity"`
}

const (
  ALL_TAGS_URL = "/tags"
)

func GetAllTags(token string) (tags []*Tag, err error) {
  tags = []*Tag{}
  request := &PieGetRequest{
    Url: ALL_TAGS_URL,
    Token: token,
  }
  err = GetPieResource(request, &tags)
  return
}
