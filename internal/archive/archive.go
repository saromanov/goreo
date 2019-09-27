package archive

import (
	"os"

	"github.com/pierrre/archivefile/zip"
	"github.com/pkg/errors"
)

// Run provides adding of the repository to the
// zip archive
func Run(path string) error {
	err := os.MkdirAll("../test_zip", os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "unable to make temp dir")
	}

	//outFilePath := filepath.Join(tmpDir, fmt.Sprintf("%s.zip", "release"))
	err = zip.ArchiveFile(path, "../release.zip", func(archivePath string) {
	})
	if err != nil {
		return errors.Wrap(err, "unable to archive file")
	}
	return nil
}
