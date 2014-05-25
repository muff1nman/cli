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

func notificationsUrl (user_id int) string {
  return fmt.Sprintf("/users/%d/notifications", user_id)
}

func buildNotificationsRequest(user_id int, token string) *pieGetRequest {
  return &pieGetRequest{
    Url: notificationsUrl(user_id),
    Token: token,
  }
}

// Gets all notifications for the given user.
func GetNotifications(user_id int, token string) (notifications []*Notification, err error) {
  notifications = []*Notification{}
  err = getPieResource(buildNotificationsRequest(user_id, token), &notifications)
  return
}

// Gets all notifications for the given user. Returns the raw response
func GetRawNotifications(user_id int, token string) (res string, err error) {
  res, err = getRawPieResource(buildNotificationsRequest(user_id, token))
  return
}
