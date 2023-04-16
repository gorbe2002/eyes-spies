package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
	_ "github.com/AOrps/eyes-spies/dance"
)

func main() {
	app := cli.NewApp()
	app.Name = "eyes-spies"
	app.Usage = "The ultimate CLI tool inspired by Ice Spice"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "gorbe2002/eyes-spies",
			Value: "Ultimate CLI tool",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "tutorialedge.net",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "dance",
			Usage: "Watch Ice Spice Dance",
			Flags: myFlags,
			Action: GenDisplayDefault("vid/munch.gif"),			},
		},
		
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	
}
