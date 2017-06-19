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
  url := "http://localhost:9000"

  app.Commands = []cli.Command{
    {
      Name:    "clusterinfo",
      Aliases: []string{"ci"},
      Usage:   "Shows status of the cluster",
      Action:  func(c *cli.Context) error {
        resp, _ := http.Get(url + "/cluster")
        defer resp.Body.Close()
        bodyBytes, _ := ioutil.ReadAll(resp.Body)
        bodyString   := string(bodyBytes)
        fmt.Println(bodyString)
        return nil
      },
    },
    {
      Name:    "nodehealth",
      Aliases: []string{"nh"},
      Usage:   "Display health of the node",
      Action:  func(c *cli.Context) error {
        resp, _ := http.Get(url + "/health")
        defer resp.Body.Close()
        bodyBytes, _ := ioutil.ReadAll(resp.Body)
        bodyString   := string(bodyBytes)
        fmt.Println(bodyString)
        return nil
      },
    },
    {
      Name:    "buildinfo",
      Aliases: []string{"bi"},
      Usage:   "Print information about cluster build",
      Action:  func(c *cli.Context) error {
        resp, _ := http.Get(url + "/info")
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
