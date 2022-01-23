package searcher

import (
	"log"
	"net/url"
	"strconv"

	"github.com/ArtyomArtamonov/filmparser/src/models"
	"github.com/ArtyomArtamonov/filmparser/src/scrapper"
	"github.com/PuerkitoBio/goquery"
)


const URL_THEPIRATEBAY_SEARCH = "https://thepiratebay10.org/search/"

func SearchFor(film *models.Film) []*models.TorrentFilm {
	res, err := scrapper.GetPage(URL_THEPIRATEBAY_SEARCH + url.QueryEscape(film.OriginalTitle))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	torrentFilms := make([]*models.TorrentFilm, 0)

	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			if indextr == 0 {
				return
			}
			torrentFilm := &models.TorrentFilm{}
			rowhtml.Find("td").Each(func(indextd int, tablecell *goquery.Selection) {
				if indextd == 1 {
					torrentFilm.Title = tablecell.Find(".detLink").Text()
					torrentFilm.DescriptionLink = tablecell.Find(".detLink").AttrOr("href", "ERROR")
					// TODO: add torrent size here
				}else if indextd == 2 {
					torrentFilm.Seeders, _ = strconv.Atoi(tablecell.Text())
				}
			})
			torrentFilms = append(torrentFilms, torrentFilm)
		})
	})

	return torrentFilms
}
