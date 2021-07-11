package scrapper

import (
	"errors"
	"log"
	"net/http"

	"github.com/ArtyomArtamonov/filmparser/src/models"
	"github.com/PuerkitoBio/goquery"
)

const URL_BASKINO_NEW string = "http://baskino.me/new/"

func GetFilms() []models.Film {
	response, err := getNewPage()
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	films := make([]models.Film, 0)
	doc.Find(".shortpost").Each(func(i int, s *goquery.Selection){
		film := models.Film{
			LocalTitle: s.Find("img").AttrOr("title", "ERROR"),
			ImageLink: s.Find("img").AttrOr("src", "https://www.google.com"),
			DescriptionLink: s.Find("a").AttrOr("href", "https://www.google.com"),
		}
		films = append(films, film)
	})

	return films
}

func getNewPage() (*http.Response, error) {
	res, err := http.Get(URL_BASKINO_NEW)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("status code error: " + string(res.StatusCode) + string(res.Status))
	}
	return res, nil
}