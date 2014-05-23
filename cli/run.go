package cli

import (
  "../pie"
  "os"
  "fmt"
  flags "github.com/jessevdk/go-flags"
)

type Options struct {
  Email string `short:"e" long:"email" description:"Your e-mail address to login."`
  Storage string `short:"d" long:"db" default:"pie.db" description:"The database file to use."`
  Topic string `short:"t" long:"topic" description:"Topic to start a new chat."`
  Thoughts string `long:"thoughts" description:"First thoughts for the new chat."`
}

func Run() {
  options := &Options{}
  flags.Parse(options)

  db, err := LoadDb(options.Storage)
  if err != nil { panic(err) }

  if db.Token == "" && options.Email == "" {
    fmt.Println("You didn't login yet. Run the app with --email a@b.c")
    os.Exit(1)
  }

  if options.Email != "" || db.Token == "" {
    login(options.Email, options.Storage, db)
  }

  if options.Topic != "" {
    post, err := pie.CreatePost(options.Topic, db.Token)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    post, err = pie.PublishPost(post.Id, db.Token)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    if options.Thoughts != "" {
      _, err := pie.CreateComment(post.Id, options.Thoughts, db.Token)

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
