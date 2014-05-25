package pie
import (
  "fmt"
  "time"
)

type Tag struct {
  Name string `json:"name"`
  NumPosts int `json:"num_posts"`
  LastActivity time.Time `json:"last_activity"`
}

func getOwnTagsUrl(user_id int) string {
  return fmt.Sprintf("/users/%d/tags", user_id)
}

func buildAllTagsRequest(token string) *pieGetRequest{
  return &pieGetRequest{
    Url: "/tags",
    Token: token,
  }
}

func buildOwnTagsRequest(user_id int, token string) *pieGetRequest{
  return &pieGetRequest{
    Url: getOwnTagsUrl(user_id),
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

// Returns all tags for the current user.
func GetOwnTags(user_id int, token string) (tags []*Tag, err error) {
  tags = []*Tag{}
  err = getPieResource(buildOwnTagsRequest(user_id, token), &tags)
  return
}

// Returns all tags for the current user. Returns raw response.
func GetRawOwnTags(user_id int, token string) (res string, err error) {
  res, err = getRawPieResource(buildOwnTagsRequest(user_id, token))
  return
}
