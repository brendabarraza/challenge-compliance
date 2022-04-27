package drivemanagement_test

import (
	googleauth "meli-challenge-compliance/pkg/google-auth"
	drivepkg "meli-challenge-compliance/pkg/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UploadFileSuccess(t *testing.T) {

	title := "gusano"
	description := "feo"
	googleClient, _ := googleauth.GetGoogleClient("../../credentials.json", "../../token.json")

	driveComponent := drivepkg.NewGoogleDriveComponent(googleClient)
	resp, err := driveComponent.UploadFile(title, description)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, title, resp.Titulo)
	assert.Equal(t, description, resp.Descripcion)

}
func Test_SearchTextInDocumentSuccess(t *testing.T) {
	googleClient, _ := googleauth.GetGoogleClient("../../credentials.json", "../../token.json")

	driveComponent := drivepkg.NewGoogleDriveComponent(googleClient)
	err := driveComponent.GetFileFromId("prueba", "1-HXD930Fs4KD5r0kPw0DeOowlBw1fdmL")

	assert.Nil(t, err)

}
