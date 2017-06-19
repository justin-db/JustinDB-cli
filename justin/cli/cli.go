package main

import (
  "os"
  "fmt"
  "sort"
  "io/ioutil"
  "net/http"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "JustinDB-Cli"
  app.Usage = "Manage JustinDB from Command Line"
  url := "http://localhost:9000/cluster"

  app.Commands = []cli.Command{
    {
      Name:    "clusterinfo",
      Aliases: []string{"ci"},
      Usage:   "Status of the cluster",
      Action:  func(c *cli.Context) error {
        resp, _ := http.Get(url)
        defer resp.Body.Close()
        bodyBytes, _ := ioutil.ReadAll(resp.Body)
        bodyString   := string(bodyBytes)
        fmt.Println(bodyString)
        return nil
      },
    },
  }

  sort.Sort(cli.FlagsByName(app.Flags))
  sort.Sort(cli.CommandsByName(app.Commands))

  app.Run(os.Args)
}
