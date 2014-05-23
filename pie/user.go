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

const (
  USER_URL = "/users/%d"
  COMPANY_USERS_URL = "/companies/%d/users"
)

func userUrl (id int) string {
  return fmt.Sprintf(USER_URL, id)
}

func companyUsersUrl (company_id int) string {
  return fmt.Sprintf(COMPANY_USERS_URL, company_id)
}

func GetUser(id int, token string) (user *User, err error) {
  user = &User{}
  request := &PieGetRequest{
    Url: userUrl(id),
    Token: token,
  }
  err = GetPieResource(request, user)
  return
}

func GetCompanyUsers(company_id int, token string) (users []*User, err error) {
  users = []*User{}
  request := &PieGetRequest{
    Url: companyUsersUrl(company_id),
    Token: token,
  }
  err = GetPieResource(request, &users)
  return
}
