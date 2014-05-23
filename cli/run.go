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

  Login struct {
    Email string `short:"e" long:"email" description:"Your e-mail address to login." required:"true"`
  } `command:"login"`

  NewPost struct {
    Topic string `short:"t" long:"topic" description:"Topic to start a new chat." required:"true"`
    Thoughts string `long:"thoughts" description:"First thoughts for the new chat."`
  } `command:"new-post"`
  Stream struct {
  } `command:"stream"`
}

func Run() (err error) {
  options := &Options{}
  parser := flags.NewParser(options, flags.Default)
  _, err = parser.Parse()
  if err != nil {
    os.Exit(1)
  }
  action := parser.Command.Active.Name

  db, err := LoadDb(options.Storage)

  if db.Token == "" && action != "login" {
    err = errors.New("You didn't login yet.")
    return
  }

  switch action {
  case "login":
    err = login(options, db)
  case "new-post":
    err = newPost(options, db)
  case "stream":
    err = stream(options, db)
  }
  return
}

func login(options *Options, db *Db) (err error) {
  var password string
  fmt.Printf("Password: ")
  _, err = fmt.Scanf("%s", &password)
  if err != nil { return }

  session, err := pie.Login(options.Login.Email, password)
  if err != nil { return }

  db.Token = session.Token
  db.UserId = session.UserId
  SaveDb(db, options.Storage)
  return
}

func newPost(options *Options, db *Db) (err error) {
  post, err := pie.CreatePost(options.NewPost.Topic, db.Token)
  if err != nil { return }

  post, err = pie.PublishPost(post.Id, db.Token)
  if err != nil { return }

  if options.NewPost.Thoughts != "" {
    _, err = pie.CreateComment(post.Id, options.NewPost.Thoughts, db.Token)
    if err != nil { return }
  }
  return
}

func stream(options *Options, db *Db) (err error) {
  posts, err := pie.Stream(db.Token)
  if err != nil { return }

  for _, post := range posts {
    fmt.Printf("{%d} %s\ncomments: %d\n\n", post.Id, post.Title, post.CommentsCount)
  }
  return
}
