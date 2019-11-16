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
)

const (
	defaultPath = "./"
)

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
	result, err := p.getPaths()
	if err != nil {
		return errors.Wrap(err, "unable to get paths from build")
	}

	archive := p.conf.GetArchive()
	p.startPipeline(archive, result)
	if err := p.execute(p.conf.After); err != nil {
		return errors.Wrap(err, "unable to apply execute after")
	}

	return nil
}

func (p *Pipeline) startPipeline(archive *config.Archive, result *builder.Response) error {
	if archive.Name == "" {
		return nil
	}
	for i, name := range result.FilePaths {
		resultSum, err := checksum.Run(p.conf.Checksum.Algorithm, name)
		if err != nil {
			return errors.Wrap(err, "unable to calc checksum")
		}

		if err := writeChecksum(resultSum); err != nil {
			return errors.Wrap(err, "unable to write check sum file")
		}
		if err := p.makeArchive(result.ArchivePaths[i], name, p.conf.GetChecksum(), p.conf.GetArchive()); err != nil {
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

func (p *Pipeline) makeArchive(name, path string, checksum *config.Checksum, archiveConf *config.Archive) error {
	archiveConf.Files = append(archiveConf.Files, path)
	if err := os.Mkdir(name, 0755); err != nil {
		return err
	}
	fileName := filepath.Base(name)
	if checksum.Algorithm != "" {
		archiveConf.Files = append(archiveConf.Files, "checksum.sum")
	}
	if len(archiveConf.Files) > 0 {
		for _, fileName := range archiveConf.Files {
			if err := copyFile(fileName, fmt.Sprintf("./%s/%s", name, fileName)); err != nil {
				fmt.Println(err)
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
			fmt.Println(err)
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

func writeChecksum(data string) error {
	return ioutil.WriteFile("checksum.sum", []byte(data), 0644)
}

func deleteFiles(files []string) error {
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return err
		}
	}

	return nil
}
