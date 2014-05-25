package pie
import (
  "time"
)

type Tag struct {
  Name string `json:"name"`
  NumPosts int `json:"num_posts"`
  LastActivity time.Time `json:"last_activity"`
}

func buildAllTagsRequest(token string) *pieGetRequest{
  return &pieGetRequest{
    Url: "/tags",
    Token: token,
  }
}

// Returns all tags for the current user.
func GetAllTags(token string) (tags []*Tag, err error) {
  tags = []*Tag{}
  err = getPieResource(buildAllTagsRequest(token), &tags)
  return
}

// Returns all tags for the current user. Returns raw response.
func GetRawAllTags(token string) (res string, err error) {
  res, err = getRawPieResource(buildAllTagsRequest(token))
  return
}
