package images

import (
	"fmt"
	"strings"

	commandhandler "github.com/pjkaufman/dotfiles/go-tools/pkg/command-handler"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
)

const imgComperssionProgramName = "imgp"

var compressionParams = []string{"-x", "800x800", "-e", "-O", "-q", "40", "-m", "-w"}
var compressableImageExts = []string{"png", "jpg", "jpeg"}

func CompressImages(destFolder, opfFolder string, images map[string]struct{}) {
	for imagePath := range images {
		if !isCompressableImage(imagePath) {
			continue
		}

		var params = append(compressionParams, filehandler.JoinPath(opfFolder, imagePath))
		commandhandler.MustRunCommand(imgComperssionProgramName, fmt.Sprintf(`failed to compress "%s"`, imagePath), params...)
	}
}

func isCompressableImage(imagePath string) bool {
	for _, ext := range compressableImageExts {
		if strings.HasSuffix(strings.ToLower(imagePath), ext) {
			return true
		}
	}

	return false
}
