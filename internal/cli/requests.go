package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	requests = &cobra.Command{
		Use:   "requests",
		Short: "Requests history [need to be authorized]",
		Long:  `If you want to convert an image, first of all need to login`,
		Run: func(cmd *cobra.Command, args []string) {
			client := &http.Client{}
			req, err := http.NewRequest("GET", url+"/requests", nil)
			if err != nil {
				fmt.Printf("error in new request:%v", err)
				return
			}
			dir, err := os.Executable()
			if err != nil {
				fmt.Printf("error in executable:%v", err)
				return
			}
			exDir := filepath.Dir(dir)
			_, err = os.Open(exDir + tokenFile)
			if err != nil {
				fmt.Println("you need to be authorized for accessing this item")
				return
			}
			data, err := os.ReadFile(exDir + tokenFile)
			if err != nil {
				fmt.Printf("error in read file:%v", err)
				return
			}
			var token jwtToken
			err = json.Unmarshal(data, &token)
			if err != nil {
				fmt.Printf("error in unmarshal:%v", err)
				return
			}
			req.Header.Set("Authorization", "Bearer "+token.Token)
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("error in client do:%v", err)
				return
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("error in read body:%v", err)
				return
			}
			if resp.StatusCode != http.StatusOK {
				fmt.Println(string(body))
				return
			}
			fmt.Println(string(body))
		},
	}
)
