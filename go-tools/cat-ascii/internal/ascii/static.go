package ascii

import (
	"embed"
	"fmt"
)

//go:embed *.txt
var asciiEmbeds embed.FS

func GetAllAsciiFileContent() ([]string, error) {
	asciiFiles, err := asciiEmbeds.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("failed to get embeded ascii files: %w", err)
	}

	var fileContent = make([]string, len(asciiFiles))
	for i, file := range asciiFiles {
		content, err := asciiEmbeds.ReadFile(file.Name())
		if err != nil {
			return nil, fmt.Errorf(`failed to get embeded ascii file content for "%s", %w`, content, err)
		}

		fileContent[i] = string(content)
	}

	return fileContent, nil
}
