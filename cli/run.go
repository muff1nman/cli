package cli

import (
  "../pie"
  "os"
  "fmt"
  "errors"
  flags "github.com/jessevdk/go-flags"
)

type Options struct {
  Storage string `short:"d" long:"db" default:"pie.db" description:"The database file to use."`

  Login *struct {
    Email string `short:"e" long:"email" description:"Your e-mail address to login." required:"true"`
  } `command:"login"`

  NewPost *struct {
    Topic string `short:"t" long:"topic" description:"Topic to start a new chat." required:"true"`
    Thoughts string `long:"thoughts" description:"First thoughts for the new chat."`
  } `command:"new-post"`
}

func Run() (err error) {
  options := &Options{}
  parser := flags.NewParser(options, flags.Default)
  _, err = parser.Parse()
  if err != nil {
    os.Exit(1)
  }

  db, err := LoadDb(options.Storage)

  if db.Token == "" && options.Login == nil {
    err = errors.New("You didn't login yet. Run the app with --email a@b.c")
    return
  }

  switch parser.Command.Active.Name {
  case "login":
    err = login(options.Login.Email, options.Storage, db)
  case "new-post":
    err = newPost(options.NewPost.Topic, options.NewPost.Thoughts, db)
  }
  return
}

func login(email string, storage string, db *Db) (err error) {
  var password string
  fmt.Printf("Password: ")
  _, err = fmt.Scanf("%s", &password)
  if err != nil { return }

  session, err := pie.Login(email, password)
  if err != nil { return }

  db.Token = session.Token
  db.UserId = session.UserId
  SaveDb(db, storage)
  return
}

func newPost(topic string, thoughts string, db *Db) (err error) {
  post, err := pie.CreatePost(topic, db.Token)
  if err != nil { return }

  post, err = pie.PublishPost(post.Id, db.Token)
  if err != nil { return }

  if thoughts != "" {
    _, err = pie.CreateComment(post.Id, thoughts, db.Token)
    if err != nil { return }
  }
  return
}
