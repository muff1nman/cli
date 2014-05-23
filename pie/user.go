package pie
import (
  "fmt"
  "time"
  "errors"
  "github.com/jmcvetta/napping"
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

const (
  USER_URL = URL_PREFIX + "/users/%d"
  COMPANY_USERS_URL = URL_PREFIX + "/companies/%d/users"
)

func userUrl (id int) string {
  return fmt.Sprintf(USER_URL, id)
}

func companyUsersUrl (company_id int) string {
  return fmt.Sprintf(COMPANY_USERS_URL, company_id)
}

func GetUser(id int, token string) (user *User, err error) {
  user = &User{}
  params := &napping.Params{"token": token}
  resp, err := napping.Get(userUrl(id), params, user, nil)
  if err != nil { return }
  if resp.Status() != 200 {
    err = errors.New("Error fetching user")
    return
  }
  return
}

func GetCompanyUsers(company_id int, token string) (users []*User, err error) {
  users = []*User{}
  params := &napping.Params{"token": token}
  resp, err := napping.Get(companyUsersUrl(company_id), params, &users, nil)
  if err != nil { return }
  if resp.Status() != 200 {
    err = errors.New("Error fetching company users")
    return
  }
  return
}
