package drivemanagement

import (
	"context"
	"encoding/json"
	"io/ioutil"
	custom_error "meli-challenge-compliance/pkg/errors"
	"net/http"
	"strings"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type GoogleDriveComponent struct {
	googleClient *http.Client
}
type createResponse struct {
	ID          string `json:"id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
}

//NewGoogleDriveComponent Crea el componente que permite administrar Google Drive
func NewGoogleDriveComponent(googleClient *http.Client) *GoogleDriveComponent {
	return &GoogleDriveComponent{
		googleClient: googleClient,
	}
}

//UploadFile Sube un archivo al Drive
func (g *GoogleDriveComponent) UploadFile(fileName, textContent string) (*createResponse, error) {
	ctx := context.TODO()
	textDescription := strings.NewReader(textContent)
	driveFile := &drive.File{
		Name:     fileName,
		MimeType: "text/plain",
	}
	srv, err := drive.NewService(ctx, option.WithHTTPClient(g.googleClient))
	if err != nil {
		return nil, custom_error.ErrDriveClient
	}
	//crea el archivo
	resDrive, err := srv.Files.Create(driveFile).Media(textDescription).Do()
	if err != nil {
		return nil, err
	}
	//cambia los permisos del archivo para que pueda ser accedido desde la aplicacion
	permissiondata := &drive.Permission{
		Type:               "domain",
		Role:               "writer",
		Domain:             "localhost.com",
		AllowFileDiscovery: true,
	}
	_, err = srv.Permissions.Create(resDrive.Id, permissiondata).Do()
	if err != nil {
		return nil, err
	}
	//retorna la respuesta a la funcion que crea el JSON marshall
	return &createResponse{
		ID:          resDrive.Id,
		Titulo:      resDrive.Name,
		Descripcion: textContent,
	}, nil
}

//GetFileFromId Busca un archivo a partir de su ID
func (g *GoogleDriveComponent) GetFileFromId(word, id string) error {
	ctx := context.TODO()

	//crea un cliente de google drive utilizando las credenciales y el token de usuario
	srv, err := drive.NewService(ctx, option.WithHTTPClient(g.googleClient))
	if err != nil {
		return custom_error.ErrDriveClient
	}

	//descarga el archivo de google drive
	r, err := srv.Files.Get(id).Download()
	if err != nil {
		return custom_error.ErrUnableFiles
	}
	//lee la respuesta del request-body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	//convierte body a tipo string
	sb := string(body)
	//devuelve el resultado de la busqueda
	isWordInDoc := strings.Contains(sb, word)
	if !isWordInDoc {
		return custom_error.ErrNotDefined
	}
	return nil
}

func (r *createResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
