package pie
import (
  "time"
  "errors"
  "github.com/jmcvetta/napping"
)

type Tag struct {
  Name string `json:"name"`
  NumPosts int `json:"num_posts"`
  LastActivity time.Time `json:"last_activity"`
}

const (
  ALL_TAGS_URL = URL_PREFIX + "/tags"
)

func GetAllTags(token string) (tags []*Tag, err error) {
  tags = []*Tag{}
  params := &napping.Params{"token": token}
  resp, err := napping.Get(ALL_TAGS_URL, params, &tags, nil)
  if err != nil { return }
  if resp.Status() != 200 {
    err = errors.New("Error fetching tags")
    return
  }
  return
}
