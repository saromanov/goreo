package archive

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pierrre/archivefile/zip"
	"github.com/pkg/errors"
)

// Run provides adding of the repository to the
// zip archive and packaging
func Run(path string, targetPath, fileName string) (string, error) {
	if err := checkOutpathPath(path); err != nil {
		return "", err
	}
	err := os.MkdirAll(targetPath, os.ModePerm)
	if err != nil {
		return "", errors.Wrap(err, "unable to make temp dir")
	}
	fileName = fileName + "zip"
	err = zip.ArchiveFile(targetPath, fileName, func(archivePath string) {
	})
	// update
	if err != nil {
		return "", errors.Wrap(err, "unable to archive file")
	}
	if err := removeContentFromDirectory(targetPath); err != nil {
		return "", errors.Wrap(err, "unable to remove files")
	}
	return fileName, nil
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

// checkOutpathPath provides checking or creating output
// archive paths
func checkOutpathPath(path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return nil
	}

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "unable to make directory")
	}
	return nil
}
