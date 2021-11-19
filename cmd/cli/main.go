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
		Use:   "signUp",
		Short: "Registration of the user",
		Long:  `If you want to convert an image, first of all need to register`,
		Run: func(cmd *cobra.Command, args []string) {
			client := &http.Client{}
			values := map[string]string{"email": Email, "password": Password}
			jsonValue, _ := json.Marshal(values)
			req, _ := http.NewRequest("POST", "http://ec2-18-193-110-163.eu-central-1.compute.amazonaws.com:8000/auth/signup", bytes.NewBufferString(string(jsonValue)))
			resp, _ := client.Do(req)
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
		},
	}
	login = &cobra.Command{
		Use:   "login",
		Short: "Authorization of the user",
		Long:  `If you want to convert an image, first of all need to login`,
		Run: func(cmd *cobra.Command, args []string) {
			client := &http.Client{}
			values := map[string]string{"email": Email, "password": Password}
			jsonValue, _ := json.Marshal(values)
			req, _ := http.NewRequest("POST", "http://ec2-18-193-110-163.eu-central-1.compute.amazonaws.com:8000/auth/login", bytes.NewBufferString(string(jsonValue)))
			resp, _ := client.Do(req)
			jwtToken, _ := ioutil.ReadAll(resp.Body)
			f, err := os.Create("user.json")
			if err != nil {
				fmt.Errorf("can't open file: %w", err)
				return
			}
			_, err = f.WriteAt(jwtToken, 0)
			if err != nil {
				fmt.Errorf("can't write into file:%w", err)
				return
			}
			fmt.Println(string(jwtToken))
		},
	}
	requests = &cobra.Command{
		Use:   "requests",
		Short: "Requests history",
		Long:  `If you want to convert an image, first of all need to login`,
		Run: func(cmd *cobra.Command, args []string) {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", "http://ec2-18-193-110-163.eu-central-1.compute.amazonaws.com:8000/requests", nil)
			token, err := os.ReadFile("user.json")

			if err != nil {
				fmt.Errorf("%w", err)
				return
			}
			req.Header.Set("Authorization", "Bearer "+string(token))
			resp, _ := client.Do(req)
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				_ = fmt.Errorf("can't read body: %w", err)
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
				fmt.Errorf("can't open file: %w", err)
				return
			}
			client := &http.Client{}
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			part, err := writer.CreateFormFile("image", filepath.Base(file.Name()))
			if err != nil {
				fmt.Errorf("can't create form file: %w", err)
				return
			}
			_, err = io.Copy(part, file)
			if err != nil {
				fmt.Errorf("can't copy file: %w", err)
				return
			}

			params := map[string]string{
				"sourceFormat": SourceFormat,
				"targetFormat": TargetFormat,
				"ratio":        Ratio,
			}
			for key, val := range params {
				_ = writer.WriteField(key, val)
			}
			writer.Close()
			req, _ := http.NewRequest("POST", "http://ec2-18-193-110-163.eu-central-1.compute.amazonaws.com:8000/image/convert", body)
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzcyNzU3OTUsImlhdCI6MTYzNzI2MTM5NSwiaWQiOjF9.BAGmJApnOGDjiV3jOv7NuFCnYNakr-NZOAbDa9kjNJc")
			req.Header.Set("Content-Type", writer.FormDataContentType())
			resp, _ := client.Do(req)
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				_ = fmt.Errorf("can't read body: %w", err)
				return
			}
			fmt.Println(string(respBody))
		},
	}
	download = &cobra.Command{
		Use:   "download",
		Short: "Download original of processed image",
		Long:  `Insert image id to download the image`,
		Run: func(cmd *cobra.Command, args []string) {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", "http://ec2-18-193-110-163.eu-central-1.compute.amazonaws.com:8000/image/download/"+ImageID, nil)
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzcyNzU3OTUsImlhdCI6MTYzNzI2MTM5NSwiaWQiOjF9.BAGmJApnOGDjiV3jOv7NuFCnYNakr-NZOAbDa9kjNJc")
			resp, _ := client.Do(req)
			image, err := os.Create(filepath.Join(Path, strings.TrimPrefix(resp.Header.Get("Content-Disposition"), "attachment; filename=")))
			if err != nil {
				fmt.Errorf("can't write file:%w", err)
				return
			}
			_, err = io.Copy(image, resp.Body)
			if err != nil {
				fmt.Errorf("error in copy image: %w", err)
				return
			}
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
