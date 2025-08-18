package storage

type Music struct {
	Artist string
	Title  string
}

func NewMusic() *Music {
	return &Music{
		Artist: "",
		Title:  "",
	}
}

func (mc *Music) SetArtist(artist string) *Music {
	mc.Artist = artist

	return mc
}

func (mc *Music) SetTitle(title string) *Music {
	mc.Title = title

	return mc
}

func (mc *Music) GetArtist() string {
	return mc.Artist
}

func (mc *Music) GetTitle() string {
	return mc.Title
}
