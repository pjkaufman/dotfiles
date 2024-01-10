package ascii

import "embed"

//go:embed *.txt
var AsciiEmbeds embed.FS

// make sure that all file names are in this list or else they will not be pulled by the program
var CatAsciiFileNames = []string{
	"scared-cat",
	"sitting-cat",
	"sleeping-cat",
	"stalking-cat",
}
