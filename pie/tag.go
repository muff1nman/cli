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

func BuildAllTagsRequest(token string) *PieGetRequest{
  return &PieGetRequest{
    Url: ALL_TAGS_URL,
    Token: token,
  }
}

func GetAllTags(token string) (tags []*Tag, err error) {
  tags = []*Tag{}
  err = GetPieResource(BuildAllTagsRequest(token), &tags)
  return
}

func GetRawAllTags(token string) (res string, err error) {
  res, err = GetRawPieResource(BuildAllTagsRequest(token))
  return
}
