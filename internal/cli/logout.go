package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	logout = &cobra.Command{
		Use:   "logout",
		Short: "Exit from the current session",
		Run: func(cmd *cobra.Command, args []string) {
			dir, err := os.Executable()
			if err != nil {
				fmt.Printf("error in executable:%v", err)
				return
			}
			exDir := filepath.Dir(dir)
			err = os.Remove(exDir + tokenFile)
			if err != nil {
				fmt.Printf("error in remove file:%v", err)
				return
			}
			fmt.Println("You have successfully logged out")
		},
	}
)
