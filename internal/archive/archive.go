package archive

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pierrre/archivefile/zip"
	"github.com/pkg/errors"
)

// addRepoToZip provides adding of the repository to the
// zip archive
func addRepoToZip(path string) error {
	tmpDir, err := ioutil.TempDir("", "test_zip")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	outFilePath := filepath.Join(tmpDir, fmt.Sprintf("%s.zip", "release"))
	err = zip.ArchiveFile(path, outFilePath, func(archivePath string) {
	})
	if err != nil {
		return errors.Wrap(err, "unable to archive file")
	}
	return nil
}
