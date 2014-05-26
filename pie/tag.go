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

func buildAllTagsRequest(token string) *request{
  return &request{
    Url: "/tags",
    Token: token,
  }
}

func buildOwnTagsRequest(user_id int, token string) *request{
  return &request{
    Url: fmt.Sprintf("/users/%d/tags", user_id),
    Token: token,
  }
}

// Returns all tags for the current user.
func GetAllTags(token string) (tags []*Tag, err error) {
  tags = []*Tag{}
  err = buildAllTagsRequest(token).doGet(&tags)
  return
}

// Returns all tags for the current user. Returns raw response.
func GetRawAllTags(token string) (res string, err error) {
  res, err = buildAllTagsRequest(token).doGetRaw()
  return
}

// Returns all tags for the current user.
func GetOwnTags(user_id int, token string) (tags []*Tag, err error) {
  tags = []*Tag{}
  err = buildOwnTagsRequest(user_id, token).doGet(&tags)
  return
}

// Returns all tags for the current user. Returns raw response.
func GetRawOwnTags(user_id int, token string) (res string, err error) {
  res, err = buildOwnTagsRequest(user_id, token).doGetRaw()
  return
}
