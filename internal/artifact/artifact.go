package artifact

import "sync"

type Type int

const (
	UploadableArchive Type = iota
	UploadableBinary
	DockerImage
)

func (t Type) String() string {
	switch t {
	case DockerImage:
		return "Docker Image"
	case UploadableArchive:
		return "Archive"
	case UploadableBinary:
		return "Binary"
	}
	return "unknown"
}

type Artifact struct {
	Name    string
	Version string
}

type Artifacts struct {
	items []*Artifact
	lock  *sync.Mutex
}

func New() Artifacts {
	return Artifacts{
		items: []*Artifact{},
		lock:  &sync.Mutex{},
	}
}

func (artifacts Artifacts) List() []*Artifact {
	return artifacts.items
}
