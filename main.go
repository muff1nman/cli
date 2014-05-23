package main

import (
  "os"
  "fmt"
  "./cli"
)

func main() {
  err := cli.Run()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
