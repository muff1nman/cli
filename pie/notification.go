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

func buildNotificationsRequest(user_id int, token string) *request {
  return &request{
    Url: fmt.Sprintf("/users/%d/notifications", user_id),
    Token: token,
  }
}

// Gets all notifications for the given user.
func GetNotifications(user_id int, token string) (notifications []*Notification, err error) {
  notifications = []*Notification{}
  err = buildNotificationsRequest(user_id, token).doGet(&notifications)
  return
}

// Gets all notifications for the given user. Returns the raw response
func GetRawNotifications(user_id int, token string) (res string, err error) {
  res, err = buildNotificationsRequest(user_id, token).doGetRaw()
  return
}
