package archive

import (
	"os"

	"github.com/pierrre/archivefile/zip"
	"github.com/pkg/errors"
)

// Run provides adding of the repository to the
// zip archive
func Run(path string, targetPath, fileName string) error {
	err := os.MkdirAll(targetPath, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "unable to make temp dir")
	}

	err = zip.ArchiveFile(path, fileName, func(archivePath string) {
	})
	if err != nil {
		return errors.Wrap(err, "unable to archive file")
	}
	return nil
}
