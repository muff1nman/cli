package pie
import (
  "fmt"
  "time"
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
  NOTIFICATIONS_URL = "/users/%d/notifications"
)

func notificationsUrl (user_id int) string {
  return fmt.Sprintf(NOTIFICATIONS_URL, user_id)
}

func GetNotifications(user_id int, token string) (notifications []*Notification, err error) {
  notifications = []*Notification{}
  err = GetPieResource(notificationsUrl(user_id), token, &notifications, nil)
  return
}
