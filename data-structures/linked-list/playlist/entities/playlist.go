package playlist_entitles

import "fmt"

type Playlist struct {
	Name       string
	Head      *Song
	NowPlaying *Song
}

type PlaylistMethodsInterface interface {
	CreatePlaylist(name string) *Playlist
	addSong(name string, artist string) error
	showAllSongs() error
	startPlaying() *Song
	nextSong() *Song
}


func (p *Playlist) addSong(name, artist string) error {
	song := &Song{
		Name:   name,
		Artist: artist,
	}

	if p.Head == nil {
		p.Head = song
	} else {
		currentNode := p.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = song
	}
	return nil
}

func (p *Playlist) showAllSongs() error {
	currentNode := p.Head
	if currentNode == nil {
		 fmt.Println("Playlist is empty")
		 return nil
	}
	fmt.Println("%+v\n", *currentNode)
	for currentNode != nil  {
		currentNode = currentNode.Next
		fmt.Println("%+v\n", *currentNode)
	}
	return nil
}

func (p *Playlist) startPlaying() error {
	p.NowPlaying = p.Head
	return nil
}