package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/hosackm/aoc2025/internal/commands"
)

var CLI struct {
	Run struct {
		Day int `arg:"" help:"Day from 1 to 25"`
	} `cmd:"" name:"run" short:"r" help:"Run a specific day of the Advent of Code"`
	Scaffold struct {
		Day       int    `arg:"" help:"Day from 1 to 25"`
		Directory string `name:"directory" help:"Directory to write files to"`
	} `cmd:"" name:"scaffold" short:"s" help:"Generate boilerplate for solving a new day"`
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Run() error {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "run <day>":
		return commands.HandleRun(CLI.Run.Day)
	case "scaffold <day>":
		err := commands.HandleScaffold(CLI.Scaffold.Day)
		if err != nil {
			return err
		}
		fmt.Printf("created day %d. make sure to register it and add an input.txt\n", CLI.Scaffold.Day)
	}

	return nil
}
