package drivemanagement

import (
	"log"
	"net/http"

	custom_error "meli-challenge-compliance/pkg/errors"

	"github.com/go-chi/chi/v5"
)

type DriveHandler struct {
	drivemanagementComponent *GoogleDriveComponent
}

// NewDriveHandler
func NewDriveHandler(drivemanagementComponent *GoogleDriveComponent) *DriveHandler {
	return &DriveHandler{
		drivemanagementComponent: drivemanagementComponent,
	}
}

// GetInvoiceFromUser crea la configuracion para buscar una palabra en un documento
func (c *DriveHandler) GetInvoiceFromUser(w http.ResponseWriter, r *http.Request) {
	documentId := chi.URLParam(r, "id")
	//si no existe el parametro lanza el error 404
	if documentId == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	// toma la palabra a buscar
	word := r.URL.Query().Get("word")

	// llama a la función que busca el archivo a partir de su ID
	err := c.drivemanagementComponent.GetFileFromId(word, documentId)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// PostFile crea la configuracion para subir un archivo en el drive
func (c *DriveHandler) PostFile(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("titulo")
	description := r.URL.Query().Get("descripcion")

	// chequea los datos
	if title == "" {
		log.Println(custom_error.ErrEmptyParameter.Error())
		http.Error(w, custom_error.ErrEmptyParameter.Error(), http.StatusNotFound)
		return
	}
	if description == "" {
		log.Println(custom_error.ErrEmptyParameter.Error())
		http.Error(w, custom_error.ErrEmptyParameter.Error(), http.StatusNotFound)
		return
	}
	// llama a la función que crea un archivo en el drive
	resp, err := c.drivemanagementComponent.UploadFile(title, description)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// crea la respuesta en JSON
	jsonResponse, err := resp.Marshal()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//devuelve el estado correctamente
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
