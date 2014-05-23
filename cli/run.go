package cli

import (
  "../pie"
  "fmt"
  "flag"
)

func Run() {
  email := flag.String("email", "", "Your e-mail address to login.")
  storage := flag.String("db", "pie.db", "The database file to use.")
  flag.Parse()

  db, err := LoadDb(*storage)
  if err != nil { panic(err) }

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
    db.UserId = session.UserId
    SaveDb(db, *storage)
  }

  user, err := pie.GetUser(db.UserId, db.Token)
  if err != nil { panic(err) }
  fmt.Println(user)

  notifications, err := pie.GetNotifications(user.Id, db.Token)
  if err != nil { panic(err) }
  fmt.Println(len(notifications))
  fmt.Println(notifications[0])

  tags, err := pie.GetAllTags(db.Token)
  if err != nil { panic(err) }
  fmt.Println(len(tags))
  fmt.Println(tags[0])

  company, err := pie.GetCompany(user.CompanyId, db.Token)
  if err != nil { panic(err) }
  fmt.Println(company)

  users, err := pie.GetCompanyUsers(user.CompanyId, db.Token)
  if err != nil { panic(err) }
  fmt.Println(len(users))
  fmt.Println(users[0])

  posts, err := pie.Stream(db.Token)
  if err != nil { panic(err) }
  fmt.Printf("Found %d posts\n", len(posts))

  post := posts[0]
  fmt.Println(post)
}
