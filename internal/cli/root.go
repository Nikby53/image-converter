package cli

import (
	"github.com/spf13/cobra"
)

const (
	url       = "http://ec2-18-193-110-163.eu-central-1.compute.amazonaws.com:8000"
	tokenFile = `\user.json`
)

type jwtToken struct {
	Token string
}

var (
	email        string
	password     string
	sourceFormat string
	targetFormat string
	ratio        string
	path         string
	imageID      string
)

var (
	rootCmd = &cobra.Command{
		Use:   "cli",
		Short: "cli client of image converter service",
		Long: ` Service that expose a RESTful API to convert JPEG to PNG and vice versa and compress the image
with the compression ratio specified by the user. The user has the ability to view the history and status of
their requests (for example, queued, processed, completed) and upload the original image and the processed one`,
	}
)

func init() {
	rootCmd.AddCommand(signUp, login, requests, convert, download, logout)
}

// New returns cobra.Command.
func New() *cobra.Command {
	return rootCmd
}
