package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"github.com/gorbe2002/eyes-spies/dance"
	"github.com/gorbe2002/eyes-spies/wttr"
	"github.com/gorbe2002/eyes-spies/networking"
	"github.com/gorbe2002/eyes-spies/txt"	
	logod "github.com/gorbe2002/eyes-spies/logoDisplay"	
	// "github.com/gorbe2002/eyes-spies/txt"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("err: no .env file found")
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "eyes-spies"
	app.Usage = "The ultimate CLI tool inspired by Ice Spice"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Show current Nameservers",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				networking.ShowNS(c.String("host"))
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Show IP Address",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				networking.ShowIP(c.String("host"))
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
		{
			Name:  "wttr",
			Usage: "Show the weather of where you are currently",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				wttr.ShowWTTR()
				return nil
			},
		},	
		{
			Name:  "logo",
			Usage: "Display our logo in ascii art",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				logod.ShowLogo("logos/roughLogo.png")
				return nil
			},
		},	
		{
			Name:  "text",
			Usage: "Send a text message from a Miami Number",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "msg",
					Value: "Yo what up",
				},
				cli.StringFlag{
					Name: "num",
					Value: "+1XXXYYYZZZZ",
				},
			},
			Action: func(c *cli.Context) error {
				msg := c.String("msg")
				num := c.String("num")
				txt.SendBasicSms(msg, num)
				return nil
			},
		},	
		
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	
}
