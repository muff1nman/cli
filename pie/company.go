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

func companyUrl (id int) string {
  return fmt.Sprintf("/companies/%d", id)
}

func buildCompanyRequest(id int, token string) *pieGetRequest {
  return &pieGetRequest{
    Url: companyUrl(id),
    Token: token,
  }
}

// Get a company by ID
func GetCompany(id int, token string) (company *Company, err error) {
  company = &Company{}
  err = getPieResource(buildCompanyRequest(id, token), company)
  return
}

// Get a company by ID. Returns a row response.
func GetRawCompany(id int, token string) (res string, err error) {
  res, err = getRawPieResource(buildCompanyRequest(id, token))
  return
}
