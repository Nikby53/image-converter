package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	login = &cobra.Command{
		Use:   "login",
		Short: "Authorization of the user",
		Run: func(cmd *cobra.Command, args []string) {
			client := &http.Client{}
			values := map[string]string{"email": email, "password": password}
			jsonValue, err := json.Marshal(values)
			if err != nil {
				fmt.Printf("error in marshal:%v", err)
				return
			}
			req, err := http.NewRequest("POST", url+"/auth/login", bytes.NewBufferString(string(jsonValue)))
			if err != nil {
				fmt.Printf("error in new request:%v", err)
				return
			}
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
			if resp.StatusCode != 200 {
				fmt.Println(string(body))
				return
			}
			dir, err := os.Executable()
			if err != nil {
				fmt.Printf("error in executable:%v", err)
				return
			}
			exDir := filepath.Dir(dir)
			f, err := os.Create(exDir + tokenFile)
			if err != nil {
				fmt.Printf("error in create:%v", err)
				return
			}
			_, err = f.WriteAt(body, 0)
			if err != nil {
				fmt.Printf("error in write to file:%v", err)
				return
			}
			fmt.Println("you successfully log in")
		},
	}
)

func init() {
	login.PersistentFlags().StringVarP(&email, "email", "e", "", "your email [required]")
	err := login.MarkPersistentFlagRequired("email")
	if err != nil {
		fmt.Println(err)
		return
	}
	login.PersistentFlags().StringVarP(&password, "password", "p", "", "your password [required]")
	err = login.MarkPersistentFlagRequired("password")
	if err != nil {
		fmt.Println(err)
		return
	}
}
