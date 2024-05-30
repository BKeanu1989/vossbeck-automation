package backup

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	// "golang.org/x/oauth2/google"
	// "google.golang.org/api/drive/v3"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func main() {
	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := getClient(config)

	srv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	fileMetadata := &drive.File{
		Name:     "NewFolder",
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{"your_shared_folder_id_here"},
	}

	file, err := srv.Files.Create(fileMetadata).Do()
	if err != nil {
		log.Fatalf("Unable to create folder: %v", err)
	}

	fmt.Printf("Folder ID: %s\n", file.Id)
}

func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}

	return config.Client(ctx, tok)
}

// func tokenFromFile(tokenPath string) (*oauth2.Token, error) {
// 	f, err := os.Open(tokenPath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()
// 	tok, err := google.FindDefaultTokenFromReader(f)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return tok, nil
// }

func tokenFromFile(tokenPath string) (*oauth2.Token, error) {
	f, err := os.Open(tokenPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read the entire file into a byte slice
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON into an oauth2.Token
	var tok oauth2.Token
	if err := json.Unmarshal(b, &tok); err != nil {
		return nil, err
	}

	return &tok, nil
}

func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"resulting code with a blank line:\n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read length from stdin: %v", err)
	}

	tok, err := config.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("Unable to retrieve tokens: %v", err)
	}

	return tok, nil
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	if contents, err := json.MarshalIndent(token, "", "  "); err != nil {
		log.Fatalf("Unable to marshal token: %v", err)
	} else {
		ioutil.WriteFile(path, contents, 0600)
	}
}
