package cli

import (
  "../pie"
  "os"
  "fmt"
  "errors"
  flags "github.com/jessevdk/go-flags"
)

type Options struct {
  Storage string `long:"db" default:"pie.db" description:"The database file to use."`
  UrlPrefix string `long:"url" default:"https://api.piethis.com/v1" description:"The API url prefix, including version."`
  Raw bool `long:"raw" default:"false" description:"Returns raw (json) responses."`

  Login struct {
    Email string `short:"e" long:"email" description:"Your e-mail address to login." required:"true"`
  } `command:"login"`

  NewPost struct {
    Topic string `short:"t" long:"topic" description:"Topic to start a new chat." required:"true"`
    Thoughts string `long:"thoughts" description:"First thoughts for the new chat."`
  } `command:"new-post"`
  Stream struct {
  } `command:"stream"`
  Notifications struct {
  } `command:"notifications"`
  Comments struct {
    PostId int `short:"p" long:"post" description:"ID of the post to add the comment" required:"true"`
  } `command:"comments"`
  NewComment struct {
    Text string `short:"t" long:"text" description:"Text for your new comment" required:"true"`
    PostId int `short:"p" long:"post" description:"ID of the post to add the comment" required:"true"`
  } `command:"new-comment"`
  AllTags struct {
  } `command:"all-tags"`
  MyTags struct {
  } `command:"my-tags"`
  Company struct {
    CompanyId int `short:"c" long:"company" description:"ID of the company." required:"true"`
  } `command:"company"`
}

func Run() (err error) {
  options := &Options{}
  parser := flags.NewParser(options, flags.Default)
  _, err = parser.Parse()
  if err != nil {
    os.Exit(1)
  }

  pie.UrlPrefix = options.UrlPrefix

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
    if options.Raw {
      err = rawStream(options, db)
    } else {
      err = stream(options, db)
    }
  case "comments":
    if options.Raw {
      err = rawComments(options, db)
    } else {
      err = comments(options, db)
    }
  case "notifications":
    if options.Raw {
      err = rawNotifications(options, db)
    } else {
      err = notifications(options, db)
    }
  case "new-comment":
    err = newComment(options, db)
  case "all-tags":
    if options.Raw {
      err = rawAllTags(options, db)
    } else {
      err = allTags(options, db)
    }
  case "my-tags":
    if options.Raw {
      err = rawMyTags(options, db)
    } else {
      err = myTags(options, db)
    }
  case "company":
    if options.Raw {
      err = rawCompany(options, db)
    } else {
      err = company(options, db)
    }
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
  post, err := pie.CreatePost(options.NewPost.Topic, db.Token, false)
  if err != nil { return }

  if options.NewPost.Thoughts != "" {
    _, err = pie.CreateComment(post.Id, options.NewPost.Thoughts, db.Token)
    if err != nil { return }
  }
  return
}

func newComment(options *Options, db *Db) (err error) {
  _, err = pie.CreateComment(options.NewComment.PostId, options.NewComment.Text, db.Token)
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

func rawStream(options *Options, db *Db) (err error) {
  posts, err := pie.RawStream(db.Token)
  if err != nil { return }

  fmt.Println(posts)
  return
}

func comments(options *Options, db *Db) (err error) {
  comments, err := pie.GetComments(options.Comments.PostId, db.Token)
  if err != nil { return }

  for _, comment := range comments {
    fmt.Printf("From: %d\n%s\n\n", comment.UserId, comment.Text)
  }
  return
}

func rawComments(options *Options, db *Db) (err error) {
  comments, err := pie.GetRawComments(options.Comments.PostId, db.Token)
  if err != nil { return }

  fmt.Println(comments)
  return
}

func notifications(options *Options, db *Db) (err error) {
  notifications, err := pie.GetNotifications(db.UserId, db.Token)
  if err != nil { return }

  for _, notification := range notifications {
    new_msg := ""
    if notification.Seen {
      new_msg = "NEW! "
    }
    fmt.Printf("%sFrom: %d\n %s %s(%d)\n\n",
      new_msg,
      notification.SenderId,
      notification.Message,
      notification.ObjectType,
      notification.ObjectId)
  }
  return
}

func rawNotifications(options *Options, db *Db) (err error) {
  notifications, err := pie.GetRawNotifications(db.UserId, db.Token)
  if err != nil { return }

  fmt.Println(notifications)
  return
}

func allTags(options *Options, db *Db) (err error) {
  tags, err := pie.GetAllTags(db.Token)
  if err != nil { return }

  for _, tag := range tags {
    fmt.Printf("%s (%d)\n", tag.Name, tag.NumPosts)
  }
  return
}

func rawAllTags(options *Options, db *Db) (err error) {
  tags, err := pie.GetRawAllTags(db.Token)
  if err != nil { return }

  fmt.Println(tags)
  return
}

func myTags(options *Options, db *Db) (err error) {
  tags, err := pie.GetOwnTags(db.UserId, db.Token)
  if err != nil { return }

  for _, tag := range tags {
    fmt.Printf("%s (%d)\n", tag.Name, tag.NumPosts)
  }
  return
}

func rawMyTags(options *Options, db *Db) (err error) {
  tags, err := pie.GetRawOwnTags(db.UserId, db.Token)
  if err != nil { return }

  fmt.Println(tags)
  return
}

func company(options *Options, db *Db) (err error) {
  company, err := pie.GetCompany(options.Company.CompanyId, db.Token)
  if err != nil { return }

  fmt.Printf("%s (%s)\n", company.Name, company.Domain)
  return
}

func rawCompany(options *Options, db *Db) (err error) {
  company, err := pie.GetRawCompany(options.Company.CompanyId, db.Token)
  if err != nil { return }

  fmt.Println(company)
  return
}
