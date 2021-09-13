package main

import (
	"fmt"
	"log"
	"netcmd/pkg"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "Website Lookup CLI",
		Usage:   "Let's you query IPs, CNAMEs and Name Servers!",
		Version: "0.0.0",
		Commands: []*cli.Command{
			{
				Name:  "ns",
				Usage: "Looks Up the NameServers for a Particular Host",
				Action: func(c *cli.Context) error {
					ns, _ := pkg.GetNS(c.Args().First())
					for _, v := range ns {
						fmt.Println(v.Host)
					}
					return nil
				},
			},
			{
				Name:  "cname",
				Usage: "Looks up the CNAME for a particular host",
				Action: func(c *cli.Context) error {
					cname, _ := pkg.GetCNAME(c.Args().First())
					fmt.Println(cname)
					return nil
				},
			},
			{
				Name:  "ip",
				Usage: "Looks up the IP addresses for a particular host",
				Action: func(c *cli.Context) error {
					ips, _ := pkg.GetIP(c.Args().First())
					for _, v := range ips {
						fmt.Println(v.String())
					}
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
