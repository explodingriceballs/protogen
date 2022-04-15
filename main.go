package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	sourceDirectory string
	outDirectory    string
	generator       string
)

func main() {
	rootCmd := &cobra.Command{
		Use: "protogen",
		Run: func(cmd *cobra.Command, args []string) {
			generator := CreateGenerator(sourceDirectory, outDirectory, generator)
			if err := generator.Run(); err != nil {
				fmt.Println("failed to generate code: ", err)
			}
		},
	}

	rootCmd.PersistentFlags().StringVar(&sourceDirectory, "source", "", "source directory / files")
	rootCmd.PersistentFlags().StringVar(&outDirectory, "out", "", "output directory")
	rootCmd.PersistentFlags().StringVar(&generator, "generator", "", "directory containing the generator")
	_ = rootCmd.MarkPersistentFlagRequired("source")
	_ = rootCmd.MarkPersistentFlagRequired("out")
	_ = rootCmd.MarkPersistentFlagRequired("generator")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
