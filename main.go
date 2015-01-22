package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/m0a/easyjson"
	"net/http"
	"os"
	"strconv"
)

var Version string = "0.0.1"

const (
	Success  = 0
	Warning  = 1
	Critical = 2
	Unknown  = 3
)

func main() {
	app := cli.NewApp()
	app.Name = "check_growthforecast_value"
	app.Version = Version
	app.Usage = "Nagios Monitoring for growthforecast value"
	app.Author = "hirocaster"
	app.Email = "hohtsuka@gmail.com"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url, u",
			Value: "http://localhost",
			Usage: "growthforecast server url",
		},
		cli.StringFlag{
			Name:  "item, i",
			Value: "/service/section/name",
			Usage: "growthforecast item",
		},
		cli.StringFlag{
			Name:  "direction, d",
			Value: "upward",
			Usage: "direction default:upward or downward",
		},
		cli.StringFlag{
			Name:  "warning, w",
			Value: "70",
			Usage: "warning value",
		},
		cli.StringFlag{
			Name:  "critical, c",
			Value: "90",
			Usage: "critical value",
		},
	}

	app.Action = func(c *cli.Context) {
		warning_value, _ := strconv.Atoi(c.String("warning"))
		critical_value, _ := strconv.Atoi(c.String("critical"))

		url := c.String("url") + "/summary/" + c.String("item")

		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		if resp.StatusCode != http.StatusOK {
			panic(resp.StatusCode)
		}

		json, err := easyjson.NewEasyJson(resp.Body)
		if err != nil {
			panic("json convert err")
		}

		current_value, _ := json.K(c.String("item")).K(0).AsInt()
		if c.String("direction") == "upward" {
			switch {
			case current_value >= critical_value:
				fmt.Println(Critical)
			case current_value >= warning_value:
				fmt.Println(Warning)
			default:
				fmt.Println(Success)
			}
		} else if c.String("direction") == "downward" {
			switch {
			case current_value <= critical_value:
				fmt.Println(Critical)
			case current_value <= warning_value:
				fmt.Println(Warning)
			default:
				fmt.Println(Success)
			}
		} else {
			fmt.Println(Unknown)
			panic("unknown direction option, only 'upward' or 'downward'")
		}
	}

	app.Run(os.Args)
}
