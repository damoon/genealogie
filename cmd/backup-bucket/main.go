package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	err := run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	app := &cli.App{
		Action: doctor,
	}

	return app.Run(os.Args)
}

func doctor(c *cli.Context) error {
	fmt.Println("Hello, world.")

	return nil
}

/*
func backup(c *cli.Context) error {

}

func verify(c *cli.Context) error {

}

func restore(c *cli.Context) error {

}
*/
