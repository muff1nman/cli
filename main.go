package main

import (
  "./pie"
  "fmt"
  "flag"
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

func main() {
  email := flag.String("email", "", "Your e-mail address to login.")
  storage := flag.String("db", "pie.db", "The database file to use.")
  flag.Parse()

  db, err := LoadDb(*storage)
  if err != nil {
    panic(err)
  }

  if db.Token == "" && *email == "" {
    panic("You didn't login yet. Run the app with --email a@b.c")
  }

  if db.Token == ""{
    var password string
    fmt.Printf("Password: ")
    _, err = fmt.Scanf("%s", &password)
    if err != nil {
      panic(err)
    }

    session, err := pie.Login(*email, password)
    if err != nil {
      panic(err)
    }
    // db.UserId = session.UserId
    db.Token = session.Token
    SaveDb(db, *storage)
  }

  posts, err := pie.Stream(db.Token)
  if err != nil {
    panic(err)
  }
  fmt.Printf("Found %d posts\n", len(posts))

  post := posts[0]
  fmt.Println(post)
}
