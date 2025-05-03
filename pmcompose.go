package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/papermerge/pmcompose/ask"
	"github.com/papermerge/pmcompose/logs"
	"github.com/papermerge/pmcompose/utils"
)

type ComposeConfig struct {
	SecretKey            string
	AppVersion           string
	UserLoginCredentials ask.Credentials
	WebAppPort           int
	S3StorageBackend     ask.S3StorageBackend
	LoggingConfigs       bool // generate and include logging configs?
}

var interactive = flag.Bool("i", false, "start interactive session")
var appVersion = flag.String("av", "3.5", "app version e.g. 3.5")
var appPort = flag.Int("ap", 12000, "web app port")
var username = flag.String("u", "admin", "username")
var password = flag.String("p", "", "password")
var withS3Backend = flag.Bool("s3", false, "with S3 backend")
var withLoggingConfigs = flag.Bool("lc", false, "with logging configs")

func main() {

	flag.Parse()

	var cfg *ComposeConfig
	var err error

	if *interactive {
		cfg, err = InteractiveSession()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	} else {
		cfg, err = GetComposeConfig()
	}

	dir, err := utils.GetExecutableDir()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	templatePath := fmt.Sprintf("%s/pmcompose_templates/docker-compose.yaml.tmpl", dir)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("Error loading template:", err)
		return
	}

	outputFile, err := os.Create("docker-compose.yml")
	if err != nil {
		fmt.Println("Error creating docker-compose.yml:", err)
		return
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, cfg)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("✅ docker-compose.yml generated successfully.")

	if cfg.LoggingConfigs {
		err = os.WriteFile("webapp_logging.yaml", []byte(logs.WEBAPP_LOGGING), 0644)
		if err != nil {
			fmt.Println("Error writing webapp_logging.yaml file")
		}
		fmt.Println("✅ webapp_logging.yaml generated successfully.")
	}

	if cfg.S3StorageBackend.S3BucketName != "" {
		if cfg.LoggingConfigs {
			err = os.WriteFile("s3worker_logging.yaml", []byte(logs.S3WORKER_LOGGING), 0644)
			if err != nil {
				fmt.Println("Error writing s3worker_logging.yaml file")
			}
			fmt.Println("✅ s3worker_logging.yaml generated successfully.")
		}
	}

}

func GetComposeConfig() (*ComposeConfig, error) {

	var cfg ComposeConfig

	secretKey, err := utils.GenerateSecretString(32)
	if err != nil {
		return nil, err
	}

	pass, err := utils.GenerateSecretString(5)
	if err != nil {
		return nil, err
	}

	if *password == "" {
		*password = pass
	}

	cfg = ComposeConfig{
		AppVersion: *appVersion,
		SecretKey:  secretKey,
		WebAppPort: *appPort,
		UserLoginCredentials: ask.Credentials{
			Username: *username,
			Password: *password,
		},
		LoggingConfigs: *withLoggingConfigs,
	}

	if *withS3Backend == true {
		cfg.S3StorageBackend.S3BucketName = "<your bucket name>"
		cfg.S3StorageBackend.AWSRegionName = "eu-central-1"
		cfg.S3StorageBackend.AWSAccessKeyID = "<your access key ID>"
		cfg.S3StorageBackend.AWSSecretAccessKey = "<your secret access key>"
	}

	return &cfg, nil
}

func InteractiveSession() (*ComposeConfig, error) {

	cfg := ComposeConfig{}

	fmt.Println(`
###################################################
###   Papermerge DMS docker compose generator   ###
###################################################
  `)

	version, err := ask.AppVersion("3.5")

	if err != nil {
		return nil, err
	}

	cfg.AppVersion = version

	appPort, err := ask.WebAppPort(12000)

	if err != nil {
		return nil, err
	}

	cfg.WebAppPort = appPort

	creds, err := ask.LoginCredentials()

	if err != nil {
		return nil, err
	}

	cfg.UserLoginCredentials = *creds

	s3StorageBackend, err := ask.StorageBackend()

	if err != nil {
		return nil, err
	}

	if s3StorageBackend != nil {
		cfg.S3StorageBackend = *s3StorageBackend
	}

	cfg.LoggingConfigs = ask.WithLogging()

	secretKey, err := utils.GenerateSecretString(32)
	if err != nil {
		return nil, err
	}
	cfg.SecretKey = secretKey

	return &cfg, nil
}
