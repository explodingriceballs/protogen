package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	sourceDirectory string
	outDirectory    string
	languages       string
)

func main() {
	rootCmd := &cobra.Command{
		Use: "browserd",
		Run: func(cmd *cobra.Command, args []string) {
			generator := CreateGenerator(sourceDirectory, outDirectory, languages)
			if err := generator.Run(); err != nil {
				fmt.Println("failed to generate code: ", err)
			}
		},
	}

	rootCmd.PersistentFlags().StringVar(&sourceDirectory, "source", ".", "source directory / files")
	rootCmd.PersistentFlags().StringVar(&outDirectory, "out", "", "output directory")
	rootCmd.PersistentFlags().StringVar(&languages, "languages", "", "comma-separated list of languages that should be generated")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
