package playlist

import playlist_entitles "github.com/nhsh1997/go-algorithms/data-structures/linked-list/playlist/entities"

func CreatePlaylist(name string) *playlist_entitles.Playlist {

	return &playlist_entitles.Playlist{
		Name: name,
	}
}
