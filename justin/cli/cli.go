package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "JustinDB-Cli"
  app.Usage = "Manage JustinDB from Command Line"
  app.Action = func(c *cli.Context) error {
    fmt.Println("Cli!")
    return nil
  }

  app.Run(os.Args)
}
