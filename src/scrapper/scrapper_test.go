package scrapper

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPage_Success(t *testing.T) {
	const url = "https://google.com"

	page, err := GetPage(url)
	if err != nil {
		t.Fatal("Could not get page")
	}
	defer page.Body.Close()

	_, err = io.ReadAll(page.Body)
	if err != nil {
		t.Fatal("Could not read page body")
	}
	
	assert.Equal(t, page.StatusCode, http.StatusOK)
}

func TestGetPage_InvalidUrl(t *testing.T) {
	const url = "google"

	_, err := GetPage(url)
	
	assert.Error(t, err)
}

func TestGetPage_InvalidStatusCode(t *testing.T) {
	const url = "http://google.com/nonexistent-page"

	_, err := GetPage(url)
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "status code error:")
}

func TestGetFilms_Success(t *testing.T) {
	films := GetFilms()

	for _, film := range films {
		assert.NotEmpty(t, film.DescriptionLink)
		assert.NotEmpty(t, film.ImageLink)
		assert.NotEmpty(t, film.LocalTitle)
		assert.NotEmpty(t, film.OriginalTitle)
	}
}
