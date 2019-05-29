package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	drive "google.golang.org/api/drive/v3"
)

var (
	credFile = "credential.json"
	tokFile  = "token.json"
)

func init() {
	if os.Getenv("ENV") == "prod" {
		credFile = os.Getenv("GD_CRED")
		tokFile = os.Getenv("GD_TKN")
	}
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) (*http.Client, error) {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok, err = getTokenFromWeb(config)
		if err != nil {
			return nil, err
		}
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok), nil
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, fmt.Errorf("unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve token from web: %v", err)
	}
	return tok, nil
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
	return nil
}

func createFile(service *drive.Service, name, mimeType string,
	content io.Reader, parentID string) (*drive.File, error) {
	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  []string{parentID},
	}
	file, err := service.Files.Create(f).Media(content).Do()

	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}

	return file, nil
}

func getService() (*drive.Service, error) {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		fmt.Printf("Unable to read %s file. Err: %v\n", credFile, err)
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		return nil, err
	}

	client, err := getClient(config)
	if err != nil {
		return nil, err
	}

	service, err := drive.New(client)
	if err != nil {
		log.Printf("Cannot create the Google Drive service: %v\n", err)
		return nil, err
	}

	return service, err
}

// UploadFile uploads the specified file to a fixed google drive account
func UploadFile(filePath string) error {

	f, err := os.Open(filePath)
	if err != nil {
		return (fmt.Errorf("cannot open file: %v", err))
	}
	defer f.Close()

	service, err := getService()
	if err != nil {
		return (fmt.Errorf("could not get service: %v", err))
	}

	file, err := createFile(service, strings.Split(filePath, "/")[1], "application/x-sql", f, "root")
	if err != nil {
		return (fmt.Errorf("could not create file: %v", err))
	}

	log.Printf("File '%s' successfully uploaded\n", file.Name)
	return nil
}
