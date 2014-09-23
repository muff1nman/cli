package main

import (
  "os"
  "fmt"
  "github.com/piethis/cli/cli"
)

func main() {
  err := cli.Run()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
