package searcher

import (
	"testing"

	"github.com/ArtyomArtamonov/filmparser/src/models"
	"github.com/stretchr/testify/assert"
)

func TestSearchFor_Success(t *testing.T) {
	films := []*models.Film{
		{
			OriginalTitle: "Harry Potter",
		},
	}

	var torrentFilms []*models.TorrentFilm
	for _, film := range films {
		torrentFilms = SearchFor(film)
	}

	assert.NotEmpty(t, torrentFilms)
}	
