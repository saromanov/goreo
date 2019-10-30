package archive

import (
	"os"
	"path/filepath"

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

	if err := removeContentFromDirectory(targetPath); err != nil {
		return errors.Wrap(err, "unable to remove files")
	}
	return nil
}

func removeContentFromDirectory(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
