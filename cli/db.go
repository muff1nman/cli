package cli

import(
  "encoding/json"
  "io/ioutil"
)

type Db struct {
  UserId int
  Token string
}

func SaveDb(db *Db, filename string) (err error) {
  file, err := json.Marshal(db)
  if err != nil { return }

  err = ioutil.WriteFile(filename, file, 0666)
  return
}
func LoadDb(filename string) (db *Db, err error) {
  db = &Db{}
  file, err := ioutil.ReadFile(filename)
  if err != nil {
    // Don't care, the file probably doesn't exist
    err = nil
    return
  }
  err = json.Unmarshal(file, db)
  return
}
