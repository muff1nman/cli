package pie
import (
  "fmt"
  "time"
  "errors"
  "github.com/jmcvetta/napping"
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
  COMPANY_URL = URL_PREFIX + "/companies/%d"
)

func companyUrl (id int) string {
  return fmt.Sprintf(COMPANY_URL, id)
}

func GetCompany(id int, token string) (company *Company, err error) {
  company = &Company{}
  params := &napping.Params{"token": token}
  resp, err := napping.Get(companyUrl(id), params, company, nil)
  if err != nil { return }
  if resp.Status() != 200 {
    err = errors.New("Error fetching company")
    return
  }
  return
}
