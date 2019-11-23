package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var app = cli.NewApp()
var beer = []string{"Enjoy your beer with some nice"}

func info() {
	app.Name = "Simple BEER CLI"
	app.Usage = "An example CLI for ordering beer's"
	app.Author = "ozombo"
	app.Version = "1.0.0"
}

//commands called from CLI package
func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "peanuts",     //commands name
			Aliases: []string{"p"}, //command alias
			Usage:   "Add peanuts to your beer",
			Action: func(c *cli.Context) { //action to be taken, append command name to beer var
				pe := "peanuts"
				peanuts := append(beer, pe)
				m := strings.Join(peanuts, " ")
				fmt.Println(m)
			},
		},
		{
			Name:    "peppersoup",
			Aliases: []string{"ps"},
			Usage:   "Add pepper soup to your beer",
			Action: func(c *cli.Context) {
				ps := "pepper soup"
				peppersoup := append(beer, ps)
				m := strings.Join(peppersoup, " ")
				fmt.Println(m)
			},
		},
		{
			Name:    "suya",
			Aliases: []string{"s"},
			Usage:   "Add suya to your beer",
			Action: func(c *cli.Context) {
				s := "suya"
				suya := append(beer, s)
				m := strings.Join(suya, " ")
				fmt.Println(m)
			},
		},
	}
}

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
