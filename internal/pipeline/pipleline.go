package pipeline

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/saromanov/goreo/internal/archive"
	"github.com/saromanov/goreo/internal/builder"
	"github.com/saromanov/goreo/internal/checksum"
	"github.com/saromanov/goreo/internal/config"
	log "github.com/sirupsen/logrus"
)

const (
	defaultPath = "./"
)

// Pipeline provides implementation of the main pipeline
type Pipeline struct {
	conf *config.Config
}

// New initialize new pipleline
func New(c *config.Config) *Pipeline {
	return &Pipeline{
		conf: c,
	}
}

// Run provides executing of the builder
func (p *Pipeline) Run() error {
	err := p.execute(p.conf.Before)
	if err != nil {
		return err
	}
	paths, err := p.getPaths()
	if err != nil {
		return errors.Wrap(err, "unable to get paths from build")
	}

	archive := p.conf.GetArchive()
	if err := p.startPipeline(archive, paths); err != nil {
		log.WithError(err).Error("unable to execute pipeline")
		return errors.Wrap(err, "unable to execute pipeline")
	}
	if err := p.execute(p.conf.After); err != nil {
		return errors.Wrap(err, "unable to apply execute after")
	}

	return nil
}

func (p *Pipeline) startPipeline(archive *config.Archive, result *builder.Response) error {
	log.Info("starting of pipeline")
	if archive.Name == "" {
		return nil
	}
	for i, name := range result.FilePaths {
		checksumConf := p.conf.GetChecksum()
		if checksumConf != nil {
			log.WithField("fileName", name).Info("Calculating of checksum")
			resultSum, err := checksum.Run(checksumConf.Algorithm, name)
			if err != nil {
				return errors.Wrap(err, "unable to calc checksum")
			}

			if err := writeChecksum(resultSum, checksumConf.Name); err != nil {
				return errors.Wrap(err, "unable to write check sum file")
			}
		}
		log.Info("making of archive")
		if err := p.makeArchive(result.ArchivePaths[i], name, checksumConf, p.conf.GetArchive()); err != nil {
			return errors.Wrap(err, "unable to archive files")
		}
	}

	return nil
}

// execute provides executing of the command
// before start of the pipeline
func (p *Pipeline) execute(commands []string) error {
	if len(commands) == 0 {
		return nil
	}

	for _, c := range commands {
		commandsSplit := strings.Split(c, " ")
		if len(commandsSplit) == 0 {
			continue
		}
		cmd := exec.Command(commandsSplit[0], commandsSplit[1:]...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return errors.Wrap(err, "unable to execute command")
		}
		fmt.Println(string(out))

	}
	return nil
}

// return list of paths
func (p *Pipeline) getPaths() (*builder.Response, error) {
	return builder.Run(p.conf.GetBuild(), p.conf.GetArchive())
}

// makeArchive provides making of archive before release
func (p *Pipeline) makeArchive(name, path string, checksum *config.Checksum, archiveConf *config.Archive) error {
	archiveConf.Files = append(archiveConf.Files, path)
	if err := os.Mkdir(name, 0755); err != nil {
		return fmt.Errorf("unable to create dir: %s %v", name, err)
	}
	fileName := filepath.Base(name)
	if checksum.Algorithm != "" {
		archiveConf.Files = append(archiveConf.Files, checksum.Name)
	}
	if len(archiveConf.Files) > 0 {
		for _, fileName := range archiveConf.Files {
			if err := copyFile(fileName, fmt.Sprintf("./%s/%s", name, fileName)); err != nil {
				log.WithError(err).Error("unable to copy file")
				return err
			}
		}
	}

	archivePath := defaultPath
	if archiveConf.Path != "" {
		archivePath = archiveConf.Path
	}

	// create archive and remove temp dir
	outArchivePath, err := archive.Run(archivePath, name, fileName)
	if err != nil {
		return errors.Wrap(err, "unable to archive files")
	}

	if archivePath != defaultPath {
		if err := copyFile(outArchivePath, fmt.Sprintf("%s/%s", archivePath, outArchivePath)); err != nil {
			log.WithError(err).Error("unable to copy file")
			return err
		}

		if err := deleteFiles([]string{fmt.Sprintf("%s/%s", defaultPath, outArchivePath)}); err != nil {
			return errors.Wrap(err, "unable to delete files")
		}
	}

	if err := deleteFiles(archiveConf.Files); err != nil {
		return errors.Wrap(err, "unable to delete files")
	}

	return nil
}

func copyFile(fileName, dest string) error {
	srcFile, err := os.Open(fileName)
	if err != nil {
		return errors.Wrap(err, "unable to open file")
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return errors.Wrap(err, "unable to create target file")
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return errors.Wrap(err, "unable to copy file")
	}

	err = destFile.Sync()
	if err != nil {
		return errors.Wrap(err, "unable to sync")
	}

	return nil
}

func writeChecksum(data, name string) error {
	return ioutil.WriteFile(name, []byte(data), 0644)
}

func deleteFiles(files []string) error {
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return err
		}
	}

	return nil
}
