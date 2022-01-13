package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	download.PersistentFlags().StringVarP(&path, "path", "p", "", "path to file [required]")

	err := download.MarkPersistentFlagRequired("path")
	if err != nil {
		fmt.Println(err)
		return
	}

	download.PersistentFlags().StringVarP(&imageID, "id", "i", "", "id of the image [required]")

	err = download.MarkPersistentFlagRequired("id")
	if err != nil {
		fmt.Println(err)
		return
	}
}

var (
	download = &cobra.Command{
		Use:   "download",
		Short: "Download original or processed image [need to be authorized]",
		Long:  `Insert image id and path where you want to save it to download the image`,
		Run: func(cmd *cobra.Command, args []string) {
			client := &http.Client{}

			req, err := http.NewRequest("GET", url+"/image/download/"+imageID, nil)
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
				fmt.Printf("error in unmarshal json:%v", err)
				return
			}

			req.Header.Set("Authorization", "Bearer "+token.Token)

			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("error in client do:%v", err)
				return
			}

			image, err := os.Create(filepath.Join(path, strings.TrimPrefix(resp.Header.Get("Content-Disposition"), "attachment; filename=")))
			if err != nil {
				fmt.Printf("error in create file:%v", err)
				return
			}

			_, err = io.Copy(image, resp.Body)
			if err != nil {
				fmt.Printf("error in copy file:%v", err)
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

			fmt.Println("you successfully downloaded the image")
		},
	}
)
