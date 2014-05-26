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

func buildCompanyRequest(id int, token string) *request {
  return &request{
    Url: fmt.Sprintf("/companies/%d", id),
    Token: token,
  }
}

// Get a company by ID
func GetCompany(id int, token string) (company *Company, err error) {
  company = &Company{}
  err = buildCompanyRequest(id, token).doGet(company)
  return
}

// Get a company by ID. Returns a row response.
func GetRawCompany(id int, token string) (res string, err error) {
  res, err = buildCompanyRequest(id, token).doGetRaw()
  return
}
