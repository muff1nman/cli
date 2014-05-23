package pie
import (
  "fmt"
  "time"
  "errors"
  "github.com/jmcvetta/napping"
)

type Notification struct {
  Id int `json:"id"`

  SenderId int `json:"sender_id"`
  ObjectId int `json:"object_id"`
  ObjectType string `json:"object_type"`
  Message string `json:"message"`
  Seen bool `json:"seen"`
  CreatedAt time.Time `json:"created_at"`
}

const (
  NOTIFICATIONS_URL = URL_PREFIX + "/users/%d/notifications"
)

func notificationsUrl (user_id int) string {
  return fmt.Sprintf(NOTIFICATIONS_URL, user_id)
}

func GetNotifications(user_id int, token string) (notifications []*Notification, err error) {
  notifications = []*Notification{}
  params := &napping.Params{"token": token}
  resp, err := napping.Get(notificationsUrl(user_id), params, &notifications, nil)
  if err != nil { return }
  if resp.Status() != 200 {
    err = errors.New("Error fetching notifications")
    return
  }
  return
}
