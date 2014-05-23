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

func GetCompany(id int, token string) (company *Company, err error) {
  company = &Company{}
  request := &PieGetRequest{
    Url: companyUrl(id),
    Token: token,
  }
  err = GetPieResource(request, company)
  return
}
