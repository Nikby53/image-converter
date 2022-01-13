package main

import (
	"fmt"
	"os"

	"github.com/Nikby53/image-converter/internal/cli"
)

func main() {
	rootCmd := cli.New()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
