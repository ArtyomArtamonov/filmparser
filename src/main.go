package main

import (
	"fmt"

	"github.com/ArtyomArtamonov/filmparser/src/scrapper"
	"github.com/ArtyomArtamonov/filmparser/src/searcher"
)


func main() {
	// Scrapping
	films := scrapper.GetFilms()

	// Searching
	for i, film := range films {
		torrentFilms := searcher.SearchFor(film)
		fmt.Printf("Film number %d\n", i + 1)
		for j := 0; j < 3; j++ {
			fmt.Printf("Title: %s\nSeeders: %d\nLink: %s\n\n", torrentFilms[j].Title, torrentFilms[j].Seeders, torrentFilms[j].DescriptionLink)
		}
		fmt.Println()
	}
	


	// Downloading
}
