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

  _, token, err := pie.Login(*email, password)
  if err != nil {
    panic(err)
  }

  posts, err := pie.Stream(token)
  if err != nil {
    panic(err)
  }
  for _, post := range posts {
    fmt.Println(post.Title)
  }
}
