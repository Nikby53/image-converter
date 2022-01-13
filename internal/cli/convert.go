package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	convert.PersistentFlags().StringVarP(&path, "path", "p", "", "path to image [required]")

	err := convert.MarkPersistentFlagRequired("path")
	if err != nil {
		fmt.Println(err)
		return
	}

	convert.PersistentFlags().StringVarP(&sourceFormat, "sourceformat", "s", "", "source format of the image [required]")

	err = convert.MarkPersistentFlagRequired("sourceformat")
	if err != nil {
		fmt.Println(err)
		return
	}

	convert.PersistentFlags().StringVarP(&targetFormat, "targetformat", "t", "", "target format of the image [required]")

	err = convert.MarkPersistentFlagRequired("targetformat")
	if err != nil {
		fmt.Println(err)
		return
	}

	convert.PersistentFlags().StringVarP(&ratio, "ratio", "r", "", "ratio of the image")
}

var (
	convert = &cobra.Command{
		Use:   "convert",
		Short: "Convert image [need to be authorized]",
		Long: `convert JPEG to PNG and vice versa and compress the image
with the compression ratio`,
		Run: func(cmd *cobra.Command, args []string) {
			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("error in open file:%v", err)
				return
			}

			image := &bytes.Buffer{}

			writer := multipart.NewWriter(image)

			part, err := writer.CreateFormFile("image", filepath.Base(file.Name()))
			if err != nil {
				fmt.Printf("error in create form file:%v", err)
				return
			}

			_, err = io.Copy(part, file)
			if err != nil {
				fmt.Printf("error in copy image: %v", err)
				return
			}

			params := map[string]string{
				"sourceFormat": sourceFormat,
				"targetFormat": targetFormat,
				"ratio":        ratio,
			}

			for key, val := range params {
				err = writer.WriteField(key, val)
				if err != nil {
					fmt.Printf("error in write fields: %v", err)
					return
				}
			}

			err = writer.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

			client := &http.Client{}

			req, err := http.NewRequest("POST", url+"/image/convert", image)
			if err != nil {
				fmt.Printf("error in new request: %v", err)
				return
			}

			dir, err := os.Executable()
			if err != nil {
				fmt.Printf("error in executable: %v", err)
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
				fmt.Printf("error in unmarhal json: %v", err)
				return
			}

			req.Header.Set("Authorization", "Bearer "+token.Token)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("error in client do: %v", err)
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("error in read body: %v", err)
				return
			}

			if resp.StatusCode != http.StatusOK {
				fmt.Println(string(body))
				return
			}

			fmt.Println("you successfully converted the image")
		},
	}
)
