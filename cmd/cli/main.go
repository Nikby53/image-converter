package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type jwtToken struct {
	Token string
}

var (
	Email        string
	Password     string
	SourceFormat string
	TargetFormat string
	Ratio        string
	Path         string
	ImageID      string
	rootCmd      = &cobra.Command{
		Use:   "imageconverter",
		Short: "An imageconverter service",
		Long: ` Service that expose a RESTful API to convert JPEG to PNG and vice versa and compress the image
with the compression ratio specified by the user. The user has the ability to view the history and status of
their requests (for example, queued, processed, completed) and upload the original image and the processed one`,
	}
	signUp = &cobra.Command{
		Use:   "signup",
		Short: "Registration of the user",
		Long:  `If you want to convert an image, first of all need to register`,
		Run: func(cmd *cobra.Command, args []string) {
			client := &http.Client{}
			values := map[string]string{"email": Email, "password": Password}
			jsonValue, _ := json.Marshal(values)
			req, err := http.NewRequest("POST", "http://ec2-18-193-110-163.eu-central-1.compute.amazonaws.com:8000/auth/signup", bytes.NewBufferString(string(jsonValue)))
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("You successfully register")
		},
	}
	login = &cobra.Command{
		Use:   "login",
		Short: "Authorization of the user",
		Long:  `If you want to convert an image, first of all need to login`,
		Run: func(cmd *cobra.Command, args []string) {
			client := &http.Client{}
			values := map[string]string{"email": Email, "password": Password}
			jsonValue, err := json.Marshal(values)
			if err != nil {
				fmt.Println(err)
				return
			}
			req, err := http.NewRequest("POST", "http://ec2-18-193-110-163.eu-central-1.compute.amazonaws.com:8000/auth/login", bytes.NewBufferString(string(jsonValue)))
			if err != nil {
				fmt.Println(err)
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(body))
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				return
			}
			f, err := os.Create(dir + "/user.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = f.WriteAt(body, 0)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("You successfully log in")
		},
	}
	requests = &cobra.Command{
		Use:   "requests",
		Short: "Requests history",
		Long:  `If you want to convert an image, first of all need to login`,
		Run: func(cmd *cobra.Command, args []string) {
			client := &http.Client{}
			req, err := http.NewRequest("GET", "http://ec2-18-193-110-163.eu-central-1.compute.amazonaws.com:8000/requests", nil)
			if err != nil {
				fmt.Errorf("%w", err)
				return
			}
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				return
			}
			data, err := os.ReadFile(dir + "/user.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			var token jwtToken
			err = json.Unmarshal(data, &token)
			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Set("Authorization", "Bearer "+token.Token)
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			if resp.StatusCode != 200 {
				fmt.Println(err)
				return
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(body))
		},
	}
	convert = &cobra.Command{
		Use:   "convert",
		Short: "Convert image",
		Long: `convert JPEG to PNG and vice versa and compress the image
with the compression ratio`,
		Run: func(cmd *cobra.Command, args []string) {
			file, err := os.Open(Path)
			if err != nil {
				fmt.Println(err)
				return
			}
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			part, err := writer.CreateFormFile("image", filepath.Base(file.Name()))
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = io.Copy(part, file)
			if err != nil {
				fmt.Println(err)
				return
			}
			params := map[string]string{
				"sourceFormat": SourceFormat,
				"targetFormat": TargetFormat,
				"ratio":        Ratio,
			}
			for key, val := range params {
				err = writer.WriteField(key, val)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			writer.Close()
			req, err := http.NewRequest("POST", "http://ec2-18-193-110-163.eu-central-1.compute.amazonaws.com:8000/image/convert", body)
			if err != nil {
				fmt.Println(err)
				return
			}
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				return
			}
			data, err := os.ReadFile(dir + "/user.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			var token jwtToken
			err = json.Unmarshal(data, &token)
			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Set("Authorization", "Bearer "+token.Token)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			fmt.Println("You successfully converted the image")
		},
	}
	download = &cobra.Command{
		Use:   "download",
		Short: "Download original of processed image",
		Long:  `Insert image id to download the image`,
		Run: func(cmd *cobra.Command, args []string) {
			client := &http.Client{}
			req, err := http.NewRequest("GET", "http://ec2-18-193-110-163.eu-central-1.compute.amazonaws.com:8000/image/download/"+ImageID, nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				return
			}
			data, err := os.ReadFile(dir + "/user.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			var token jwtToken
			err = json.Unmarshal(data, &token)
			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Set("Authorization", "Bearer "+token.Token)
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			image, err := os.Create(filepath.Join(Path, strings.TrimPrefix(resp.Header.Get("Content-Disposition"), "attachment; filename=")))
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = io.Copy(image, resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("You successfully download the image")
		},
	}
)

func init() {
	rootCmd.AddCommand(signUp, login, requests, convert, download)
	signUp.PersistentFlags().StringVarP(&Email, "email", "e", "", "pass your email")
	signUp.MarkPersistentFlagRequired("email")
	signUp.PersistentFlags().StringVarP(&Password, "password", "p", "", "pass your password")
	signUp.MarkPersistentFlagRequired("password")
	login.PersistentFlags().StringVarP(&Email, "email", "e", "", "pass your email")
	login.MarkPersistentFlagRequired("email")
	login.PersistentFlags().StringVarP(&Password, "password", "p", "", "pass your password")
	login.MarkPersistentFlagRequired("password")
	convert.PersistentFlags().StringVarP(&Path, "path", "p", "", "path to file")
	convert.MarkPersistentFlagRequired("path")
	convert.PersistentFlags().StringVarP(&SourceFormat, "sourceformat", "s", "", "source format of the image")
	convert.MarkPersistentFlagRequired("sourceformat")
	convert.PersistentFlags().StringVarP(&TargetFormat, "targetformat", "t", "", "target format of the image")
	convert.MarkPersistentFlagRequired("targetformat")
	convert.PersistentFlags().StringVarP(&Ratio, "ratio", "r", "", "ratio of the image")
	download.PersistentFlags().StringVarP(&Path, "path", "p", "", "path to file")
	download.MarkPersistentFlagRequired("path")
	download.PersistentFlags().StringVarP(&ImageID, "id", "i", "", "id of the image")
	download.MarkPersistentFlagRequired("id")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
