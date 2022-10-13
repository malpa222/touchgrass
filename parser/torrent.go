package parser

type Torrent struct {
	Announce string
	Info     InfoDict
}

type InfoDict struct {
	Name        string
	PieceLength int
	Pieces      string
	Length      int
	Path        []string
}
