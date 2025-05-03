package ask

import (
	"bufio"
	"crypto/subtle"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"

	"golang.org/x/term"
)

func ReadInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func WebAppPort(defaultValue int) (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(
		`The port on which the Papermerge DMS webserver
will listen for incoming connections.`)

	fmt.Printf("Port [%d]:", defaultValue)
	appPort := ReadInput(reader)
	if appPort != "" {
		result, err := strconv.Atoi(appPort)

		if err != nil {
			return -1, err
		}

		return result, nil
	}

	return defaultValue, nil
}

func AppVersion(defaultValue string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Papermerge DMS version")

	fmt.Printf("Version [%s]:", defaultValue)
	version := ReadInput(reader)
	if version == "" {
		return defaultValue, nil
	}

	return version, nil
}

func LoginCredentials() (*Credentials, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Specify initial login credentials.")

	creds := Credentials{}
	var username string

	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Username [%s]:", currentUser.Username)
	username = ReadInput(reader)
	if username == "" {
		creds.Username = currentUser.Username
	} else {
		creds.Username = username
	}

	passwordMatch := false

	for passwordMatch == false {
		fmt.Print("Password:")
		password1, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return nil, err
		}
		if len(password1) == 0 {
			fmt.Println("Cannot be empty")
			continue
		}
		fmt.Println()
		fmt.Print("Password (again):")
		password2, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		if err != nil {
			return nil, err
		}

		passwordMatch = comparePasswords(password1, password2)

		if passwordMatch {
			creds.Password = string(password1)
			break
		} else {
			fmt.Println("Password did not match")
		}
	}

	return &creds, nil
}

func StorageBackend() (*S3StorageBackend, error) {
	fmt.Println(`
  S3 Storage
  ===========

  When using S3 storage, Papermerge will upload/download documents
  to/from your private S3 location. This may be useful for example
  when you plan run multiple webapp containers or you want
  to use CDN (currently works only with AWS CloudFront) for serving
  document files.
  `)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Would you use S3 storage? (yes no) [no]:")

	yesno := ""

	for yesno != "yes" || yesno != "no" {
		yesno := ReadInput(reader)

		if yesno == "" || yesno == "no" {
			return nil, nil
		}

		if yesno != "yes" {
			fmt.Print("Answer can be either yes or no")
		}

		if yesno == "yes" {
			break
		}
	}

	backend := S3StorageBackend{}
	fmt.Print("AWS negion name? [eu-central-1]:")
	region := ReadInput(reader)
	if region == "" {
		region = "eu-central-1"
	}
	backend.AWSRegionName = region

	fmt.Print("S3 bucket name:")
	bucketName := ReadInput(reader)
	backend.S3BucketName = bucketName

	fmt.Print("AWS_ACCESS_KEY_ID:")
	accessKeyID := ReadInput(reader)
	backend.AWSAccessKeyID = accessKeyID

	fmt.Print("AWS_SECRET_ACCESS_KEY:")
	secretAccessKey := ReadInput(reader)
	backend.AWSSecretAccessKey = secretAccessKey

	return &backend, nil

}

func comparePasswords(password1, password2 []byte) bool {
	// Use subtle.ConstantTimeCompare for secure comparison
	return subtle.ConstantTimeCompare(password1, password2) == 1
}
