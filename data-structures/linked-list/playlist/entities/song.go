package playlist_entitles

type Song struct {
	Name string
	Artist string
	Next *Song
}

