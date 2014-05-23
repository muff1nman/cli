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

func BuildNotificationsRequest(user_id int, token string) *PieGetRequest {
  return &PieGetRequest{
    Url: notificationsUrl(user_id),
    Token: token,
  }
}

func GetNotifications(user_id int, token string) (notifications []*Notification, err error) {
  notifications = []*Notification{}
  err = GetPieResource(BuildNotificationsRequest(user_id, token), &notifications)
  return
}

func GetRawNotifications(user_id int, token string) (res string, err error) {
  res, err = GetRawPieResource(BuildNotificationsRequest(user_id, token))
  return
}
