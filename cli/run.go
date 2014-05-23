package cli

import (
  "../pie"
  "os"
  "fmt"
  "flag"
)

func Run() {
  email := flag.String("email", "", "Your e-mail address to login.")
  storage := flag.String("db", "pie.db", "The database file to use.")
  new_chat_topic := flag.String("topic", "", "Topic to start a new chat.")
  new_chat_comment := flag.String("thoughts", "", "First thoughts for the new chat.")
  flag.Parse()

  db, err := LoadDb(*storage)
  if err != nil { panic(err) }

  if db.Token == "" && *email == "" {
    fmt.Println("You didn't login yet. Run the app with --email a@b.c")
    os.Exit(1)
  }

  if *email != "" || db.Token == "" {
    login(*email, *storage, db)
  }

  if *new_chat_topic != "" {
    post, err := pie.CreatePost(*new_chat_topic, db.Token)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    post, err = pie.PublishPost(post.Id, db.Token)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    if new_chat_comment != nil {
      _, err := pie.CreateComment(post.Id, *new_chat_comment, db.Token)

      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
    }

    fmt.Println(post)
  }
  // user, err := pie.GetUser(db.UserId, db.Token)
  // if err != nil { panic(err) }
  // fmt.Println(user)

  // notifications, err := pie.GetNotifications(user.Id, db.Token)
  // if err != nil { panic(err) }
  // fmt.Println(len(notifications))
  // fmt.Println(notifications[0])

  // tags, err := pie.GetAllTags(db.Token)
  // if err != nil { panic(err) }
  // fmt.Println(len(tags))
  // fmt.Println(tags[0])

  // company, err := pie.GetCompany(user.CompanyId, db.Token)
  // if err != nil { panic(err) }
  // fmt.Println(company)

  // users, err := pie.GetCompanyUsers(user.CompanyId, db.Token)
  // if err != nil { panic(err) }
  // fmt.Println(len(users))
  // fmt.Println(users[0])

  // posts, err := pie.Stream(db.Token)
  // if err != nil { panic(err) }
  // fmt.Printf("Found %d posts\n", len(posts))

  // post := posts[0]
  // fmt.Println(post)
}

func login(email string, storage string, db *Db) {
  var password string
  fmt.Printf("Password: ")
  _, err := fmt.Scanf("%s", &password)
  if err != nil {
    panic(err)
  }

  session, err := pie.Login(email, password)
  if err != nil {
    panic(err)
  }
  // db.UserId = session.UserId
  db.Token = session.Token
  db.UserId = session.UserId
  SaveDb(db, storage)
}
