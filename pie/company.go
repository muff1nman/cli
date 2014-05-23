package pie
import (
  "fmt"
  "time"
)

type Company struct {
  Id int `json:"id"`
  GroupId int `json:"group_id"`
  Domain string `json:"domain"`
  Logo string `json:"logo"`
  Name string `json:"name"`
  CreatedAt time.Time `json:"created_at"`
}

const (
  COMPANY_URL = "/companies/%d"
)

func companyUrl (id int) string {
  return fmt.Sprintf(COMPANY_URL, id)
}

func BuildCompanyRequest(id int, token string) *PieGetRequest {
  return &PieGetRequest{
    Url: companyUrl(id),
    Token: token,
  }
}

func GetCompany(id int, token string) (company *Company, err error) {
  company = &Company{}
  err = GetPieResource(BuildCompanyRequest(id, token), company)
  return
}

func GetRawCompany(user_id int, token string) (res string, err error) {
  res, err = GetRawPieResource(BuildCompanyRequest(user_id, token))
  return
}
