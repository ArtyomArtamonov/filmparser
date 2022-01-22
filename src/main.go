package main

import (
	"fmt"

	"github.com/ArtyomArtamonov/filmparser/src/scrapper"
	"github.com/ArtyomArtamonov/filmparser/src/searcher"
)

// How many films do we get from baskino.me
var MAX_FILMS int = 5

// How many sources do we get from thepiratebay10.org
var MAX_TORRENTSOURCES = 1

func main() {
	// Scrapping
	films := scrapper.GetFilms()

	// Searching
	for i, film := range films {
		if i == MAX_FILMS {
			break
		}
		torrentFilms := searcher.SearchFor(film)

		fmt.Printf("Film number %d\n", i + 1)
		for j, torrentFilm := range torrentFilms {
			if j == MAX_TORRENTSOURCES {
				break
			}
			fmt.Printf("Title: %s\nSeeders: %d\nLink: %s\n\n", torrentFilm.Title, torrentFilm.Seeders, torrentFilm.DescriptionLink)
		}
		fmt.Println()
	}

	// TODO: Downloading
}
