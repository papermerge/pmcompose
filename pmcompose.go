package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/papermerge/pmcompose/ask"
	"github.com/papermerge/pmcompose/utils"
)

type ComposeConfig struct {
	SecretKey            string
	AppVersion           string
	UserLoginCredentials ask.Credentials
	WebAppPort           int
	S3StorageBackend     ask.S3StorageBackend
}

func main() {

	fmt.Println(`
###################################################
###   Papermerge DMS docker compose generator   ###
###################################################
  `)

	cfg := ComposeConfig{}

	version, err := ask.AppVersion("3.5")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	cfg.AppVersion = version

	appPort, err := ask.WebAppPort(12000)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	cfg.WebAppPort = appPort

	creds, err := ask.LoginCredentials()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	cfg.UserLoginCredentials = *creds

	s3StorageBackend, err := ask.StorageBackend()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if s3StorageBackend != nil {
		cfg.S3StorageBackend = *s3StorageBackend
	}

	secretKey, err := utils.GenerateSecretString(32)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	cfg.SecretKey = secretKey

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
}
