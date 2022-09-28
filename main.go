package main

import (
	"fmt"
	"github.com/integrii/flaggy"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	sourceDirectory string
	outDirectory    string
	generator       string
	includeDirs     []string
)

func main() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	// Set the name & basics
	flaggy.SetName("protogen")
	flaggy.ShowHelpOnUnexpectedDisable()

	// Flags
	flaggy.String(&sourceDirectory, "s", "source", "source directory / files")
	flaggy.String(&outDirectory, "o", "out", "output directory")
	flaggy.String(&generator, "g", "generator", "directory containing the generator")
	flaggy.StringSlice(&includeDirs, "i", "include", "include directories")

	// Parse
	flaggy.Parse()

	// Check for supplied flags
	if sourceDirectory == "" || outDirectory == "" || generator == "" {
		flaggy.ShowHelp("")
		return
	}

	// Run the generator
	generator := CreateGenerator(sourceDirectory, outDirectory, generator)
	generator.SetInclDirs(includeDirs)
	if err := generator.Run(); err != nil {
		fmt.Println("failed to generate code: ", err)
	}
}
