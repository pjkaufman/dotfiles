package filehandler

import (
	"context"
	"fmt"
	"os"

	archiver "github.com/mholt/archiver/v4"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

func ConvertRarToCbz(src string) {
	format := archiver.Rar{}

	var fileList = []archiver.File{}
	handler := func(ctx context.Context, f archiver.File) error {
		fmt.Println(f)
		fileList = append(fileList, f)

		return nil
	}

	cbrFile, err := os.Open(src)
	if err != nil {
		logger.WriteError(fmt.Sprintf(`failed to read in cbr file "%s": %s`, src, err))
	}

	err = format.Extract(context.Background(), cbrFile, nil, handler)
	if err != nil {
		logger.WriteError(fmt.Sprintf(`failed to extract data from cbr file "%s": %s`, src, err))
	}

	createCbz(fileList)

	logger.WriteError("TODO: implement the rest")
}

func createCbz(fileList []archiver.File) {
	var cbzName = "example.cbz"
	out, err := os.Create(cbzName)
	if err != nil {
		logger.WriteError(fmt.Sprintf(`failed to create "%s": %s`, cbzName, err))
	}
	defer out.Close()

	format := archiver.CompressedArchive{
		Archival: archiver.Zip{},
	}

	err = format.Archive(context.Background(), out, fileList)
	if err != nil {
		logger.WriteError(fmt.Sprintf(`failed to write contents to "%s": %s`, cbzName, err))
	}
}
