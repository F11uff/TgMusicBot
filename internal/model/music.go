package model

type Music struct {
	artist string
	music  string
}

func NewMusic(artist string, music string) *Music {
	return &Music{
		artist: artist,
		music:  music,
	}
}

func (m *Music) SetArtist(artist string) *Music {
	m.artist = artist

	return m
}

func (m *Music) SetMusic(music string) *Music {
	m.music = music

	return m
}

func (m *Music) GetArtist() string {
	return m.artist
}

func (m *Music) GetMusic() string {
	return m.music
}

func (m *Music) ClearArtistAndMusic() *Music {
	m.artist = ""
	m.music = ""

	return m
}
