package main

import (
  "./pie"
  "fmt"
  "flag"
)

func main() {
  email := flag.String("email", "", "Your e-mail address to login.")
  flag.Parse()

  var password string
  fmt.Printf("Password: ")
  _, err := fmt.Scanf("%s", &password)
  if err != nil {
    panic(err)
  }

  session, err := pie.Login(*email, password)
  if err != nil {
    panic(err)
  }

  posts, err := pie.Stream(session.Token)
  if err != nil {
    panic(err)
  }
  fmt.Printf("Found %d posts\n", len(posts))

  post := posts[0]
  fmt.Println(post.CreatedAt)
  fmt.Println(post.Title)
}
