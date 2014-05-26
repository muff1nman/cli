package pie
import (
  "fmt"
  "time"
)

type User struct {
  Id int `json:"id"`
  CompanyId int `json:"company_id"`

  Email string `json:"email"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  Blurb string `json:"blurb"`
  ProfileImage string `json:"profile_image"`
  BgImage string `json:"bg_image"`
  CommentCount int `json:"comment_count"`
  CreatedAt time.Time `json:"created_at"`
}

func buildUserRequest(id int, token string) *request{
  return &request{
    Url: fmt.Sprintf("/users/%d", id),
    Token: token,
  }
}

func buildCompanyUsersRequest(company_id int, token string) *request{
  return &request{
    Url: fmt.Sprintf("/companies/%d/users", company_id),
    Token: token,
  }
}

// Gets a user by ID.
func GetUser(id int, token string) (user *User, err error) {
  user = &User{}
  err = buildUserRequest(id, token).doGet(user)
  return
}

// Gets a user by ID. Returns a raw response.
func GetRawUser(id int, token string) (res string, err error) {
  res, err = buildUserRequest(id, token).doGetRaw()
  return
}

// Gets all users for a given company.
func GetCompanyUsers(company_id int, token string) (users []*User, err error) {
  users = []*User{}
  err = buildCompanyUsersRequest(company_id, token).doGet(&users)
  return
}

// Gets all users for a given company. Returns a raw response.
func GetRawCompanyUsers(company_id int, token string) (res string, err error) {
  res, err = buildCompanyUsersRequest(company_id, token).doGetRaw()
  return
}
