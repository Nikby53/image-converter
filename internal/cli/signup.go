package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	signUp = &cobra.Command{
		Use:   "signup",
		Short: "Registration of the user",
		Long:  `If you want to convert an image, first of all need to register`,
		Run: func(cmd *cobra.Command, args []string) {
			client := &http.Client{}
			values := map[string]string{"email": email, "password": password}
			jsonValue, err := json.Marshal(values)
			if err != nil {
				fmt.Printf("error in marshal:%v", err)
				return
			}
			req, err := http.NewRequest("POST", url+"/auth/signup", bytes.NewBufferString(string(jsonValue)))
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
			fmt.Println("you have successfully register")
		},
	}
)

func init() {
	signUp.PersistentFlags().StringVarP(&email, "email", "e", "", "your email [required]")
	err := signUp.MarkPersistentFlagRequired("email")
	if err != nil {
		fmt.Println(err)
		return
	}
	signUp.PersistentFlags().StringVarP(&password, "password", "p", "", "your password [required] ")
	err = signUp.MarkPersistentFlagRequired("password")
	if err != nil {
		fmt.Println(err)
		return
	}
}
