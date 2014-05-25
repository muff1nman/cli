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

func userUrl (id int) string {
  return fmt.Sprintf("/users/%d", id)
}

func companyUsersUrl (company_id int) string {
  return fmt.Sprintf("/companies/%d/users", company_id)
}

func buildUserRequest(id int, token string) *pieGetRequest{
  return &pieGetRequest{
    Url: userUrl(id),
    Token: token,
  }
}

func buildCompanyUsersRequest(company_id int, token string) *pieGetRequest{
  return &pieGetRequest{
    Url: companyUsersUrl(company_id),
    Token: token,
  }
}

// Gets a user by ID.
func GetUser(id int, token string) (user *User, err error) {
  user = &User{}
  err = getPieResource(buildUserRequest(id, token), user)
  return
}

// Gets a user by ID. Returns a raw response.
func GetRawUser(id int, token string) (res string, err error) {
  res, err = getRawPieResource(buildUserRequest(id, token))
  return
}

// Gets all users for a given company.
func GetCompanyUsers(company_id int, token string) (users []*User, err error) {
  users = []*User{}
  err = getPieResource(buildCompanyUsersRequest(company_id, token), &users)
  return
}

// Gets all users for a given company. Returns a raw response.
func GetRawCompanyUsers(company_id int, token string) (res string, err error) {
  res, err = getRawPieResource(buildCompanyUsersRequest(company_id, token))
  return
}
