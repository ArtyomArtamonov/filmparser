package models

type Film struct {
	LocalTitle string
	OriginalTitle string
	DescriptionLink string
	ImageLink string
}

type TorrentFilm struct {
	Title string
	DescriptionLink string
	Seeders int
	Size string
}
