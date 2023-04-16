package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"

	"github.com/urfave/cli"
	"github.com/gorbe2002/eyes-spies/dance"
	// "github.com/gorbe2002/eyes-spies/txt"
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
			Action: func(c *cli.Context) error {
				dance.GenDisplayDefault("vid/munch.gif")
				return nil
			},
		},
		{
			Name:  "os",
			Usage: "Check Operating System & Architecture",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				os := runtime.GOOS
				arch := runtime.GOARCH
				fmt.Printf("\tOS : {  %s  }\tARCH: {  %s  }\n", os, arch)
				return nil
			},
		},	
		
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	
}
