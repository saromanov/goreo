package archive

import (
	"fmt"
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
	err = zip.ArchiveFile(targetPath, fileName+".zip", func(archivePath string) {
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
		return errors.Wrap(err, fmt.Sprintf("unable to open dir %s", dir))
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return errors.Wrap(err, "unable to read dir names")
	}
	for _, name := range names {
		path := filepath.Join(dir, name)
		err = os.RemoveAll(path)
		if err != nil {
			return errors.Wrap(err, "unable to remove files")
		}
	}
	if err := os.Remove(dir); err != nil {
		return errors.Wrap(err, "unable to remove directory")
	}
	return nil
}
