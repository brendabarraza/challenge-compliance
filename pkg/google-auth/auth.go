package googleauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

// GetGoogleClient Recupera un token, guarda el token y luego devuelve el cliente generado.
func GetGoogleClient(credentialFileLocation, tokenFileLocation string) (*http.Client, error) {
	file, err := ioutil.ReadFile(credentialFileLocation)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(file, drive.DriveScope, drive.DriveFileScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	tokenFile := tokenFileLocation
	token, err := tokenFromFile(tokenFile)
	if err != nil {
		token = getTokenFromWeb(config)
		// El archivo token.json almacena el tokens de acceso y se crea automáticamente cuando el flujo de autorización se completa por primera vez.
		saveToken(tokenFile, token)
	}

	return config.Client(context.Background(), token), nil
}

// getTokenFromWeb Solicita un token de la web
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenFromFile Recupera un token de un archivo local.
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

// saveToken Guarda el token en un archivo.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
